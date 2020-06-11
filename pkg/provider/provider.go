package provider

import (
	"errors"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"

	"github.com/Rentberry/geocoder/pkg/cache"
	"github.com/Rentberry/geocoder/pkg/config"
	geocoder "github.com/Rentberry/geocoder/pkg/geocoder"
)

const GoogleProviderType = "google"
const OpencageProviderType = "opencage"

type Result struct {
	Provider  string
	Locations []*geocoder.Location
	Metadata  metadata.MD
}

type AggregateProvider struct {
	cs       cache.CacheStore
	google   GeocodingProvider
	opencage GeocodingProvider
}

var (
	providerVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "geo_provider_metric",
	}, []string{"provider", "reverse"})
	cacheVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "geo_cache_hit",
	}, []string{"provider", "cache"})
	countryVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "geo_country",
	}, []string{"provider", "country"})
	providerResponseTimesHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "geo_provider_response_times",
		Buckets: []float64{50, 100, 250, 500, 1000, 2000, 5000},
	}, []string{"provider"})
)

func NewGeocodingProvider(cs cache.CacheStore, registry *prometheus.Registry, specification config.Specification) (GeocodingProvider, error) {
	registry.MustRegister(providerVec, cacheVec, countryVec, providerResponseTimesHistogram)

	google, err := NewGoogleGeocodeProvider(specification.GoogleApiKey)
	if err != nil {
		return nil, err
	}
	opencage, err := NewOpencageProvider(specification.OpencageApiKey)
	if err != nil {
		return nil, err
	}

	return &AggregateProvider{
		cs:       cs,
		google:   google,
		opencage: opencage,
	}, nil
}

type GeocodingProvider interface {
	Geocode(q Query) (*Result, error)
}

func (ap AggregateProvider) Geocode(q Query) (*Result, error) {
	if q.Provider == "" {
		q.Provider = OpencageProviderType
	}

	k := q.Hash()
	result, err := ap.checkInCache(k)
	if err != nil {
		logrus.Warning(err)
	}

	if result != nil {
		cacheVec.WithLabelValues(q.Provider, "hit").Inc()
		result.Metadata = metadata.Pairs("cache-key", string(k))

		return result, nil
	}

	cacheVec.WithLabelValues(q.Provider, "miss").Inc()

	t1 := time.Now()
	switch q.Provider {
	case GoogleProviderType:
		result, err = ap.google.Geocode(q)

	case OpencageProviderType:
		result, err = ap.opencage.Geocode(q)
	}

	providerResponseTimesHistogram.
		WithLabelValues(q.Provider).
		Observe(float64(time.Since(t1).Milliseconds()))

	providerVec.
		WithLabelValues(q.Provider, strconv.FormatBool(q.isReverse())).
		Inc()

	if err == nil {
		err = ap.storeInCache(result, k)
		if err != nil {
			logrus.Warning(err)
		}

		if result != nil && len(result.Locations) > 0 {
			countryVec.WithLabelValues(q.Provider, result.Locations[0].Country.Code).Inc()
			result.Metadata = metadata.Pairs("cache-key", string(k))
		}
	}

	return result, err
}

func (ag *AggregateProvider) checkInCache(key []byte) (*Result, error) {
	item, err := ag.cs.Get(key)
	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, nil
	}

	res, ok := item.(Result)
	if !ok {
		return nil, errors.New("corrupted cached result")
	}

	return &res, nil
}

func (ag *AggregateProvider) storeInCache(result *Result, key []byte) error {
	if result != nil {
		err := ag.cs.Set(key, *result)
		if err != nil {
			return err
		}
	} else {
		err := ag.cs.SetWithTTL(key, nil, 24*time.Hour)
		if err != nil {
			return err
		}
	}

	return nil
}
