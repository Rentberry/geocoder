package server

import (
	"context"
	"fmt"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/Rentberry/geocoder/pkg/config"
	"github.com/Rentberry/geocoder/pkg/geocoder"
	"github.com/Rentberry/geocoder/pkg/provider"
)

type geocoderService struct {
	provider provider.GeocodingProvider
	geocoder.UnimplementedGeocodeServiceServer
}

func NewGeocoderServer(cs provider.CacheStore, registry *prometheus.Registry, specification config.Specification) (geocoder.GeocodeServiceServer, error) {
	p, err := provider.NewGeocodingProvider(cs, registry, specification)
	if err != nil {
		return nil, err
	}

	return &geocoderService{
		provider: p,
	}, nil
}

func (g geocoderService) Geocode(ctx context.Context, r *geocoder.LocationRequest) (*geocoder.LocationResponse, error) {
	q := provider.Query{
		Address:  r.Address,
		Provider: r.Provider,
		Query:    r.Query,
	}

	if r.Latlng != nil {
		q.Lat = r.Latlng.Lat
		q.Lng = r.Latlng.Lng
	}

	if !q.IsValid() {
		return nil, fmt.Errorf("invalid query %+v", q)
	}

	res, err := g.provider.Geocode(q)
	if err != nil {
		return nil, fmt.Errorf("%w; query: %+v", err, q)
	}

	grpc_ctxtags.Extract(ctx).
		Set("request.address", q.Address).
		Set("request.provider", q.Provider).
		Set("request.lat", q.Lat).
		Set("request.lng", q.Lng)

	if res == nil {
		return &geocoder.LocationResponse{
			Locations: nil,
			Exists:    false,
		}, nil
	}

	grpc_ctxtags.Extract(ctx).
		Set("response.country", res.Locations[0].Country.Code)

	err = grpc.SetTrailer(ctx, res.Metadata)
	if err != nil {
		log.Warn(err)
	}

	return &geocoder.LocationResponse{
		Locations: res.Locations,
		Exists:    true,
	}, nil
}
