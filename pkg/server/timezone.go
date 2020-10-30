package server

import (
	"context"

	"github.com/bradfitz/latlong"

	geocoder "github.com/Rentberry/geocoder/pkg/geocoder"
)

type TimezoneServer struct {
	geocoder.UnimplementedTimezoneServiceServer
}

func NewTimezoneServer() *TimezoneServer {
	return &TimezoneServer{}
}

func (t TimezoneServer) Lookup(ctx context.Context, r *geocoder.TimezoneRequest) (*geocoder.Timezone, error) {
	tx := latlong.LookupZoneName(r.Latlng.Lat, r.Latlng.Lng)

	return &geocoder.Timezone{Code: tx}, nil
}
