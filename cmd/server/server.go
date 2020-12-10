package main

import (
	"flag"
	"fmt"
	"github.com/Rentberry/geocoder/pkg/provider"
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

	"github.com/Rentberry/geocoder/pkg/config"
	"github.com/Rentberry/geocoder/pkg/geocoder"
	"github.com/Rentberry/geocoder/pkg/server"
)

var cfg config.Specification
var debug bool

var (
	reg         = prometheus.NewRegistry()
	grpcMetrics = grpc_prometheus.NewServerMetrics()
)

func init() {
	flag.BoolVar(&debug, "debug", false, "Debug mode")
}

func main() {
	flag.Parse()

	logrusLogger := logrus.New()

	logrusLogger.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
		logrusLogger.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrusLogger.SetFormatter(&logrus.TextFormatter{})
	}

	grpcMetrics.EnableHandlingTimeHistogram()
	logrusEntry := logrus.NewEntry(logrusLogger)
	logrusOpts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}

	var err error
	cfg, err = processConfig()
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

	cs, err := provider.NewCacheStore(rc)
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

	// Starts http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logrus.Fatal("unable to start a http server.")
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
		DB:           cfg.RedisDatabase,
		MinIdleConns: 2,
		MaxConnAge:   1 * time.Hour,
	})

	status := client.Ping()
	if err := status.Err(); err != nil {
		return nil, err
	}

	logrus.Infof("connected to redis: %s:%s[%d]", cfg.RedisHost, cfg.RedisPort, cfg.RedisDatabase)

	return client, nil
}

func processConfig() (config.Specification, error) {
	var s config.Specification
	err := envconfig.Process("", &s)
	return s, err
}
