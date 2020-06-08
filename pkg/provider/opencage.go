package provider

import (
	"fmt"
	"strings"

	geocoder "github.com/Rentberry/geocoder/pkg/geocoder"
	opencagedata "github.com/Rentberry/geocoder/pkg/opencage"
)

type OpencageProvider struct {
	client *opencagedata.Geocoder
}

func NewOpencageProvider(key string) (GeocodingProvider, error) {
	client := opencagedata.NewGeocoder(key)
	return &OpencageProvider{client: client}, nil
}

func (o OpencageProvider) Geocode(q Query) (*Result, error) {
	var qs string
	if q.isReverse() {
		qs = fmt.Sprintf("%f,%f", q.Lat, q.Lng)
	} else {
		qs = q.Address
	}

	language, ok := q.Query["language"]
	if !ok {
		language = "en"
	}

	results, err := o.client.Geocode(qs, &opencagedata.GeocodeParams{CountryCode: q.Query["country"], Language: language})
	if err != nil {
		return nil, err
	}

	if len(results.Results) == 0 {
		return nil, nil
	}

	res := &Result{
		Provider:  OpencageProviderType,
		Locations: make([]*geocoder.Location, len(results.Results)),
	}

	for i, result := range results.Results {
		loc := &geocoder.Location{
			Provider:         OpencageProviderType,
			FormattedAddress: result.Formatted,
			LatLng:           &geocoder.LatLng{Lat: float64(result.Geometry.Latitude), Lng: float64(result.Geometry.Longitude)},
			Bounds: &geocoder.Bounds{
				SouthWest: &geocoder.LatLng{Lat: float64(result.Bounds.SouthWest.Latitude), Lng: float64(result.Bounds.SouthWest.Longitude)},
				NorthEast: &geocoder.LatLng{Lat: float64(result.Bounds.NorthEast.Latitude), Lng: float64(result.Bounds.NorthEast.Longitude)},
			},
			Country: &geocoder.Country{
				Code: strings.ToUpper(stringify(result.Components["country_code"])),
				Name: stringify(result.Components["country"]),
			},
			State: &geocoder.State{
				Code: stringify(result.Components["state_code"]),
				Name: stringify(result.Components["state"]),
			},
			StreetName: stringify(matchOpencageComponent(
				result,
				[]string{"road", "footway", "street", "street_name", "residential", "path", "pedestrian", "road_reference", "road_reference_intl"},
			)),
			StreetNumber: stringify(result.Components["house_number"]),
			Locality: stringify(matchOpencageComponent(
				result,
				[]string{"city", "town", "municipality", "village", "hamlet", "locality", "croft"},
			)),
			Sublocality: stringify(matchOpencageComponent(
				result, []string{"neighbourhood", "suburb", "city_district", "district", "quarter", "houses", "subdivision"},
			)),
			PostalCode: stringify(result.Components["postcode"]),
			Timezone:   stringify(result.Components["timezone"]),
		}

		if state, ok := result.Components["state"]; ok {
			loc.AdminLevels = append(loc.AdminLevels, &geocoder.AdminLevel{Name: stringify(state), Code: stringify(result.Components["state_code"]), Level: 1})
		}

		if county, ok := result.Components["county"]; ok {
			loc.AdminLevels = append(loc.AdminLevels, &geocoder.AdminLevel{Name: stringify(county), Level: 2})
		}

		res.Locations[i] = loc
	}

	return res, nil
}

func matchOpencageComponent(data opencagedata.GeocodeResultItem, keys []string) interface{} {
	var match interface{}
	for _, v := range keys {
		var ok bool
		match, ok = data.Components[v]
		if ok {
			break
		}
	}

	return match
}

func stringify(i interface{}) string {
	if i == nil {
		return ""
	}
	return fmt.Sprintf("%s", i)
}
