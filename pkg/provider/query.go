package provider

import (
	"crypto/sha1"
	"encoding/gob"
	"encoding/hex"

	"github.com/sirupsen/logrus"
)

type Query struct {
	Address  string
	Provider string
	Lat      float64
	Lng      float64
	Reverse  bool
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

func (q *Query) isReverse() bool {
	return q.Reverse
}
