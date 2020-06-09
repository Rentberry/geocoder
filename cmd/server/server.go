package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-redis/redis/v7"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/Rentberry/geocoder/pkg/cache"
	"github.com/Rentberry/geocoder/pkg/config"
	"github.com/Rentberry/geocoder/pkg/geocoder"
	"github.com/Rentberry/geocoder/pkg/server"
)

var cfg config.Specification
var debug bool

var (
	logrusLogger = logrus.New()
	reg          = prometheus.NewRegistry()
	grpcMetrics  = grpc_prometheus.NewServerMetrics()
)

func init() {
	flag.BoolVar(&debug, "debug", false, "Debug mode")
}

func main() {
	flag.Parse()

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}
	logrus.SetFormatter(&logrus.TextFormatter{})

	grpcMetrics.EnableHandlingTimeHistogram()
	logrusEntry := logrus.NewEntry(logrusLogger)
	logrusOpts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}

	var err error
	cfg, err = proccessConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	lis, err := net.Listen("tcp", cfg.ListenAddr)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Listening at %s\r\n", cfg.ListenAddr)

	rc, err := setupRedis()
	if err != nil {
		logrus.Fatal(err)
	}

	cs, err := cache.NewCacheStore(rc)
	if err != nil {
		logrus.Fatal(err)
	}

	srv, err := server.NewGeocoderServer(cs, reg, cfg)
	if err != nil {
		logrus.Fatal(err)
	}

	// Create a HTTP server for prometheus.
	httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: cfg.MetricsListenAddr}

	grpcServer := grpc.NewServer(
		grpc_middleware.WithStreamServerChain(
			grpcMetrics.StreamServerInterceptor(),
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.StreamServerInterceptor(logrusEntry, logrusOpts...),
		),
		grpc_middleware.WithUnaryServerChain(
			grpcMetrics.UnaryServerInterceptor(),
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry, logrusOpts...),
		),
	)

	geocoder.RegisterGeocodeServiceServer(grpcServer, srv)
	geocoder.RegisterTimezoneServiceServer(grpcServer, server.NewTimezoneServer())

	// Initialize all metrics.
	reg.MustRegister(grpcMetrics)
	grpcMetrics.InitializeMetrics(grpcServer)

	// Start your http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logrus.Fatal("Unable to start a http server.")
		}
	}()

	err = grpcServer.Serve(lis)
	if err != nil {
		logrus.Error(err)
	}
}

func setupRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		DB:           15,
		MinIdleConns: 2,
		MaxConnAge:   1 * time.Hour,
	})

	status := client.Ping()
	if err := status.Err(); err != nil {
		return nil, err
	}

	return client, nil
}

func proccessConfig() (config.Specification, error) {
	var s config.Specification
	err := envconfig.Process("", &s)
	return s, err
}
