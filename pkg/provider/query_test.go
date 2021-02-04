package provider

import "testing"

func TestIsValidQuery(t *testing.T) {
	testCases := []struct {
		Address       string
		Lat           float64
		Lng           float64
		Query         map[string]string
		ShouldBeValid bool
	}{
		// Valid - place_id is not empty
		{"", 0, 0, map[string]string{"place_id": "foo"}, true},
		// Invalid - address, coordinates and place_id are empty
		{"", 0, 0, map[string]string{"place_id": ""}, false},
		{"", 0, 0, map[string]string{}, false},
		// Valid - coordinates present and in correct range
		{"", 42, 42, map[string]string{}, true},
		// Invalid - coordinate value out of range
		{"", 91, 0, map[string]string{}, false},
		{"", -91, 0, map[string]string{}, false},
		{"", 0, 181, map[string]string{}, false},
		{"", 0, -181, map[string]string{}, false},
		// Valid - address is not empty. Unicode is supported
		{"foo-bar (baz)", 0, 0, map[string]string{}, true},
		{"روات", 0, 0, map[string]string{}, true},
		{"北京", 0, 0, map[string]string{}, true},
		{"İstanbul", 0, 0, map[string]string{}, true},
		// Invalid - does not contain language or mark characters
		{"___", 0, 0, map[string]string{}, false},
		{"   ", 0, 0, map[string]string{}, false},
		{"()", 0, 0, map[string]string{}, false},
	}

	for i, tc := range testCases {
		q := Query{
			Address: tc.Address,
			Lat:     tc.Lat,
			Lng:     tc.Lng,
			Query:   tc.Query,
		}

		var expectedValidity string
		if tc.ShouldBeValid {
			expectedValidity = "valid"
		} else {
			expectedValidity = "invalid"
		}

		if q.IsValid() != tc.ShouldBeValid {
			t.Errorf("%d: %#v is expected to be %s but it is not", i, q, expectedValidity)
		}
	}
}
