package provider

import (
	"context"

	"googlemaps.github.io/maps"

	geocoder "github.com/Rentberry/geocoder/pkg/geocoder"
)

type GoogleProvider struct {
	client *maps.Client
}

func NewGoogleGeocodeProvider(key string) (GeocodingProvider, error) {
	c, err := maps.NewClient(maps.WithAPIKey(key))
	if err != nil {
		return nil, err
	}
	return &GoogleProvider{client: c}, nil
}

func (g GoogleProvider) Geocode(q Query) (*Result, error) {
	var results []maps.GeocodingResult
	var err error

	language, ok := q.Query["language"]
	if !ok {
		language = "en"
	}

	r := &maps.GeocodingRequest{
		Address:    q.Address,
		LatLng:     &maps.LatLng{Lat: q.Lat, Lng: q.Lng},
		Language:   language,
		Components: make(map[maps.Component]string),
	}

	if _, ok := q.Query["place_id"]; ok {
		r.Address = ""
		r.LatLng = nil
		r.PlaceID = q.Query["place_id"]
	}

	if _, ok := q.Query["country"]; ok {
		r.Components[maps.ComponentCountry] = q.Query["country"]
	}

	if _, ok := q.Query["locality"]; ok {
		r.Components[maps.ComponentLocality] = q.Query["locality"]
	}

	if r.PlaceID != "" || r.LatLng != nil {
		results, err = g.client.ReverseGeocode(context.Background(), r)
	} else {
		results, err = g.client.Geocode(context.Background(), r)
	}

	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	res := &Result{
		Provider:  GoogleProviderType,
		Locations: make([]*geocoder.Location, len(results))}

	for i, result := range results {
		var loc = &geocoder.Location{
			Provider:         GoogleProviderType,
			Id:               result.PlaceID,
			FormattedAddress: results[0].FormattedAddress,
			Country:          &geocoder.Country{},
			LatLng: &geocoder.LatLng{
				Lat: results[0].Geometry.Location.Lat,
				Lng: results[0].Geometry.Location.Lng,
			},
			Components: make(map[string]string),
		}
		loc.Country.Code, loc.Country.Name = matchGoogleComponent(result, "country")

		if _, neighborhood := matchGoogleComponent(result, "neighborhood"); neighborhood != "" {
			loc.Components["neighborhood"] = neighborhood
		}

		if _, point_of_interest := matchGoogleComponent(result, "point_of_interest"); point_of_interest != "" {
			loc.Components["point_of_interest"] = point_of_interest
		}

		updateFromComponents(loc, result.AddressComponents)
		updateBounds(loc, result)

		res.Locations[i] = loc
	}

	return res, nil
}

func updateBounds(location *geocoder.Location, result maps.GeocodingResult) {
	if result.Geometry.Bounds.SouthWest.Lat != 0 && result.Geometry.Bounds.SouthWest.Lng != 0 {
		location.Bounds = &geocoder.Bounds{
			SouthWest: &geocoder.LatLng{Lat: result.Geometry.Bounds.SouthWest.Lat, Lng: result.Geometry.Bounds.SouthWest.Lng},
			NorthEast: &geocoder.LatLng{Lat: result.Geometry.Bounds.NorthEast.Lat, Lng: result.Geometry.Bounds.NorthEast.Lng},
		}

		return
	}

	if result.Geometry.Viewport.SouthWest.Lat != 0 && result.Geometry.Viewport.SouthWest.Lng != 0 {
		location.Bounds = &geocoder.Bounds{
			SouthWest: &geocoder.LatLng{Lat: result.Geometry.Viewport.SouthWest.Lat, Lng: result.Geometry.Viewport.SouthWest.Lng},
			NorthEast: &geocoder.LatLng{Lat: result.Geometry.Viewport.NorthEast.Lat, Lng: result.Geometry.Viewport.NorthEast.Lng},
		}

		return
	}

	// Fake bounds
	if result.Geometry.LocationType == "ROOFTOP" {
		location.Bounds = &geocoder.Bounds{
			SouthWest: &geocoder.LatLng{Lat: result.Geometry.Location.Lat, Lng: result.Geometry.Location.Lng},
			NorthEast: &geocoder.LatLng{Lat: result.Geometry.Location.Lat, Lng: result.Geometry.Location.Lng},
		}

		return
	}
}

func updateFromComponents(location *geocoder.Location, components []maps.AddressComponent) {
	for _, v := range components {
		for _, vv := range v.Types {
			switch vv {
			case "postal_code":
				location.PostalCode = v.LongName
			case "locality", "postal_town":
				location.Locality = v.LongName
			case "street_number":
				location.StreetNumber = v.LongName
			case "route":
				location.StreetName = v.LongName
			case "sublocality":
				location.Sublocality = v.LongName
			case "neighborhood":
				if location.Sublocality == "" {
					location.Sublocality = v.LongName
				}
			case "administrative_area_level_1":
				location.AdminLevels = append(location.AdminLevels, &geocoder.AdminLevel{Name: v.LongName, Code: v.ShortName, Level: 1})
				location.State = &geocoder.State{Name: v.LongName, Code: v.ShortName}
			case "administrative_area_level_2":
				location.AdminLevels = append(location.AdminLevels, &geocoder.AdminLevel{Name: v.LongName, Code: v.ShortName, Level: 2})
			case "administrative_area_level_3":
				location.AdminLevels = append(location.AdminLevels, &geocoder.AdminLevel{Name: v.LongName, Code: v.ShortName, Level: 3})
			}
		}
	}
}

func matchGoogleComponent(data maps.GeocodingResult, component string) (string, string) {
	for _, v := range data.AddressComponents {
		for _, vv := range v.Types {
			if vv == component {
				return v.ShortName, v.LongName
			}
		}
	}

	return "", ""
}
