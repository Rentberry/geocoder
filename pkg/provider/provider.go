package provider

import (
	"encoding/gob"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"

	"github.com/Rentberry/geocoder/pkg/config"
	geocoder "github.com/Rentberry/geocoder/pkg/geocoder"
)

const GoogleProviderType = "google"
const OpencageProviderType = "opencage"
const CacheTTL = 365 * 24 * time.Hour

type Result struct {
	Provider  string
	Locations []*geocoder.Location
	Metadata  metadata.MD
}

type AggregateProvider struct {
	cs       CacheStore
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

func init() {
	gob.Register(Result{})
}

func NewGeocodingProvider(cs CacheStore, registry *prometheus.Registry, specification config.Specification) (GeocodingProvider, error) {
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

	result, err := ap.checkInCache(q)
	if err != nil {
		logrus.Warning(err)
	}

	if result != nil {
		cacheVec.WithLabelValues(q.Provider, "hit").Inc()

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
		err = ap.storeInCache(result, q)
		if err != nil {
			logrus.Warning(err)
		}

		if result != nil && len(result.Locations) > 0 {
			countryVec.WithLabelValues(q.Provider, result.Locations[0].Country.Code).Inc()

		}
	}

	return result, err
}

func (ag *AggregateProvider) checkInCache(q Query) (*Result, error) {
	key := q.Key()
	item, err := ag.cs.Get(key)
	if err != nil {
		return nil, err
	}

	if item != nil {
		item.Metadata = metadata.Pairs("cache-key", string(key))
		return item, nil
	}

	// To gracefully migrate old keys to new ones
	oldKey := q.Hash()
	item, err = ag.cs.Get(oldKey)
	if err != nil {
		return nil, err
	}

	if item != nil {
		err = ag.cs.Set(key, item)
		if err != nil {
			return item, err
		}

		err = ag.cs.Del(oldKey)
		if err != nil {
			return item, err
		}
	}

	return item, nil
}

func (ag *AggregateProvider) storeInCache(result *Result, q Query) error {
	key := q.Key()
	if result != nil {
		err := ag.cs.SetWithTTL(key, result, CacheTTL)
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
