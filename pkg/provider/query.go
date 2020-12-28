package provider

import (
	"crypto/sha1"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

var queryRegexp = regexp.MustCompile("\\w+")

type Query struct {
	Address  string
	Provider string
	Lat      float64
	Lng      float64
	Query    map[string]string
}

func (q Query) Hash() []byte {
	h := sha1.New()
	enc := gob.NewEncoder(h)
	err := enc.Encode(q)

	if err != nil {
		logrus.Fatal(err)
	}

	src := h.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(src)))

	hex.Encode(dst, src)

	return dst
}

func (q Query) Key() []byte {
	h := sha1.New()
	data := []string{
		url.QueryEscape(strings.ToLower(q.Address)),
		url.QueryEscape(strings.ToLower(q.Provider)),
		fmt.Sprintf("%.6f", q.Lat),
		fmt.Sprintf("%.6f", q.Lng),
	}
	for k, v := range q.Query {
		data = append(data, url.QueryEscape(k), url.QueryEscape(v))
	}
	base := strings.Join(data, "&")

	h.Write([]byte(base))

	src := h.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(src)))

	hex.Encode(dst, src)

	return dst
}

func (q *Query) isReverse() bool {
	return q.Address == ""
}

func (q Query) IsValid() bool {
	if q.Address == "" && q.Lat == 0 && q.Lng == 0 {
		if placeId, ok := q.Query["place_id"]; !ok || placeId == "" {
			return false
		}
	}

	if q.Address != "" && !queryRegexp.MatchString(q.Address) {
		return false
	}

	if q.Lat < -90 || q.Lat > 90 {
		return false
	}

	if q.Lng < -180 || q.Lng > 180 {
		return false
	}

	return true
}
