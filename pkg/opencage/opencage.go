package opencage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const endpoint = "https://api.opencagedata.com/geocode/v1/"

type Geocoder struct {
	Key        string
	HttpClient *http.Client
}

type GeocodeParams struct {
	// Country hint
	CountryCode string
	Language    string
}

type GeocodeResult struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`

	Rate struct {
		Limit     int   `json:"limit"`
		Remaining int   `json:"remaining"`
		Reset     int64 `json:"reset"`
	} `json:"rate"`

	Results []GeocodeResultItem `json:"results"`
}

type GeocodeResultItem struct {
	Confidence int                    `json:"confidence"`
	Formatted  string                 `json:"formatted"`
	Geometry   Geometry               `json:"geometry"`
	Components map[string]interface{} `json:"components"`

	Bounds struct {
		NorthEast Geometry `json:"northeast"`
		SouthWest Geometry `json:"southwest"`
	} `json:"bounds"`
}

type Geometry struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lng"`
}

type GeocodeError struct {
	Result *GeocodeResult
}

func (err *GeocodeError) Error() string {
	return fmt.Sprintf("%d: %s", err.Result.Status.Code, err.Result.Status.Message)
}

func NewGeocoder(key string) *Geocoder {
	return &Geocoder{
		Key: key,
		HttpClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:        8,
				MaxIdleConnsPerHost: 0,
				MaxConnsPerHost:     32,
				IdleConnTimeout:     1 * time.Minute,
				ReadBufferSize:      2 << 10,
			},
		},
	}
}

func (g *Geocoder) Geocode(query string, params *GeocodeParams) (*GeocodeResult, error) {
	u := g.geocodeUrl(query, params)
	resp, err := g.HttpClient.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result := &GeocodeResult{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, err
	}
	if result.Status.Code != 200 {
		return nil, &GeocodeError{Result: result}
	}

	return result, nil
}

func (g *Geocoder) geocodeUrl(query string, params *GeocodeParams) string {
	u, _ := url.Parse(endpoint)
	u.Path += "json"

	q := u.Query()
	q.Set("q", query)
	q.Set("key", g.Key)
	if params != nil {
		if params.CountryCode != "" {
			q.Set("countrycode", strings.ToLower(params.CountryCode))
		}

		if params.Language != "" {
			q.Set("language", strings.ToLower(params.Language))
		}
	}

	u.RawQuery = q.Encode()
	log.Debug(u.String())

	return u.String()
}
