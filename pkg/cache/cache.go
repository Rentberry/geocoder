package cache

import (
	"bytes"
	"encoding/gob"
	"github.com/dgraph-io/ristretto"
	"github.com/go-redis/redis/v7"
	"time"
)

var DefaultTTL = 180 * 24 * time.Hour

type CacheStore interface {
	Get(key []byte) (interface{}, error)
	Set(key []byte, item interface{}) error
	SetWithTTL(key []byte, item interface{}, ttl time.Duration) error
}

type Store struct {
	rdb *redis.Client
	mem *ristretto.Cache
}

type ristrettoItem struct {
	item     interface{}
	expireAt time.Time
}

func NewCacheStore(rc *redis.Client) (*Store, error) {
	var numberOfItems int64 = 1024
	rs, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: numberOfItems * 10,
		MaxCost:     numberOfItems,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}

	return &Store{rdb: rc, mem: rs}, nil
}

func (s *Store) Get(key []byte) (interface{}, error) {
	item, ok := s.mem.Get(key)
	if ok && item != nil {
		ritem, ok := item.(ristrettoItem)
		if ok && !time.Now().After(ritem.expireAt) {
			return ritem.item, nil
		}
		s.mem.Del(key)
	}

	b, err := s.rdb.Get(string(key)).Bytes()
	if err == redis.Nil || b == nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var res interface{}
	dec := gob.NewDecoder(bytes.NewReader(b))
	err = dec.Decode(&res)
	if err != nil {
		s.rdb.Del(string(key))
		return nil, err
	}

	return res, nil
}

func (s *Store) Set(key []byte, item interface{}) error {
	return s.SetWithTTL(key, item, DefaultTTL)
}

func (s Store) SetWithTTL(key []byte, item interface{}, ttl time.Duration) error {
	data, err := s.Serialize(item)
	if err != nil {
		return err
	}

	err = s.rdb.Set(string(key), data, ttl).Err()
	if err != nil {
		return err
	}

	s.mem.Set(key, ristrettoItem{item: item, expireAt: time.Now().Add(ttl)}, 1)

	return nil
}

func (s Store) Serialize(item interface{}) ([]byte, error) {
	if item == nil {
		return nil, nil
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(item)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
