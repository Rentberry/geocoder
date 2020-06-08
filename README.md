# gRPC based geocoding service
Adapter for external geocoding services with caching, monitoring, metrics etc. Has built-in support for timezone queries (resolve coordinates to timezone) implemented using [amazing Brad Fitzpatrick's library](https://github.com/bradfitz/latlong). 
Supports [Google's Geocoding API](https://developers.google.com/maps/documentation/geocoding/start) and [Opencage](https://opencagedata.com) for now.

# Usage
## Docker image
```
docker run \
    -p 8080:8080 -p 9092:9092 \
    -e REDIS_HOST=redis \
    -e GOOGLE_API_KEY_SERVER=<your-google-key> \
    -e OPENCAGE_API_KEY=<your-opencage-key> \
    rentberry:geocoder
```
This will start geocoding service that listens for grpc connections 
on port `8080` and exposes prometheus metrics on `/metrics` on port `9092`

## gRPC methods
### Geocoding
Forward and reverse geocoding
```GeocodeService.Geocode (LocationRequest) returns (LocationResponse)```
#### Example
Request:
```json5
// Forward
{
  "address": "berlin", 
  "provider": "opencage", 
  "query": {
    "country": "de", 
    "language": "en"
  }
}

// Reverse
{
  "latLng": {"lat":52.51006317138672, "lng":13.40505599975586}, 
  "provider": "opencage",
  "query": {
    "language": "en"
  }
}
```
Response:
```json5
{
  "locations": [
    {
      "provider": "opencage",
      "formattedAddress": "Best Western Hotel am Spittelmarkt, Neue Grünstraße 28, 10179 Berlin, Germany",
      "country": {
        "name": "Germany",
        "code": "DE"
      },
      "streetNumber": "28",
      "streetName": "Neue Grünstraße",
      "locality": "Berlin",
      "sublocality": "Spandauer Vorstadt",
      "postalCode": "10179",
      "latLng": {
        "lat": 52.51006317138672,
        "lng": 13.40505599975586
      },
      "adminLevels": [
        {
          "level": 1,
          "name": "Berlin",
          "code": "BE"
        }
      ],
      "state": {
        "name": "Berlin",
        "code": "BE"
      },
      "bounds": {
        "northEast": {
          "lat": 52.510196685791016,
          "lng": 13.405620574951172
        },
        "southWest": {
          "lat": 52.50993347167969,
          "lng": 13.40455150604248
        }
      }
    },
    {
      "provider": "opencage",
      "formattedAddress": "Neue Grünstraße 28, 10179 Berlin, Germany",
      "country": {
        "name": "Germany",
        "code": "DE"
      },
      "streetNumber": "28",
      "streetName": "Neue Grünstraße",
      "locality": "Berlin",
      "sublocality": "Spandauer Vorstadt",
      "postalCode": "10179",
      "latLng": {
        "lat": 52.5099983215332,
        "lng": 13.40558910369873
      },
      "adminLevels": [
        {
          "level": 1,
          "name": "Berlin",
          "code": "BE"
        }
      ],
      "state": {
        "name": "Berlin",
        "code": "BE"
      },
      "bounds": {
        "northEast": {
          "lat": 52.510047912597656,
          "lng": 13.405638694763184
        },
        "southWest": {
          "lat": 52.50994873046875,
          "lng": 13.405538558959961
        }
      }
    },
    {
      "provider": "opencage",
      "formattedAddress": "10179, Germany",
      "country": {
        "name": "Germany",
        "code": "DE"
      },
      "postalCode": "10179",
      "latLng": {
        "lat": 52.51359939575195,
        "lng": 13.423199653625488
      },
      "state": {

      },
      "bounds": {
        "northEast": {

        },
        "southWest": {

        }
      }
    }
  ],
  "exists": true
}

// Response trailers received:
// cache-key: b6d1655e0fc8945a33b050c4dfdb43b2cb0ecf58
```

### Timezone lookup
Find location's timezone by provided coordinates
````TimezoneService.Lookup (TimezoneRequest) returns (Timezone)````
#### Example
Request:
```json5
{"latlng": {"lat":52.51006317138672, "lng":13.40505599975586}}
```
Response:
```json5
{
  "code": "Europe/Berlin"
}
```
