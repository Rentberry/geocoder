package provider

import (
	"bytes"
	"encoding/gob"
	"github.com/dgraph-io/ristretto"
	"github.com/go-redis/redis/v7"
	"time"
)

var DefaultTTL = 180 * 24 * time.Hour

type CacheStore interface {
	Get(key []byte) (*Result, error)
	Set(key []byte, item *Result) error
	SetWithTTL(key []byte, item *Result, ttl time.Duration) error
	Del(key []byte) error
}

type Store struct {
	rdb *redis.Client
	mem *ristretto.Cache
}

type ristrettoItem struct {
	Empty    bool
	Item     Result
	ExpireAt time.Time
}

func NewCacheStore(rc *redis.Client) (*Store, error) {
	var numberOfItems int64 = 2048
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

func (s *Store) Get(key []byte) (*Result, error) {
	item, ok := s.mem.Get(key)
	if ok && item != nil {
		ritem, ok := item.(ristrettoItem)
		if ok && !time.Now().After(ritem.ExpireAt) {
			if ritem.Empty {
				return nil, nil
			}
			return &ritem.Item, nil
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

	var res Result
	dec := gob.NewDecoder(bytes.NewReader(b))
	err = dec.Decode(&res)
	if err != nil {
		s.rdb.Del(string(key))
		return nil, err
	}

	return &res, nil
}

func (s *Store) Set(key []byte, item *Result) error {
	return s.SetWithTTL(key, item, DefaultTTL)
}

func (s Store) SetWithTTL(key []byte, item *Result, ttl time.Duration) error {
	data, err := s.Serialize(item)
	if err != nil {
		return err
	}

	err = s.rdb.Set(string(key), data, ttl).Err()
	if err != nil {
		return err
	}

	if item != nil {
		s.mem.Set(key, ristrettoItem{Item: *item, ExpireAt: time.Now().Add(ttl)}, 1)
	} else {
		s.mem.Set(key, ristrettoItem{Empty: true, ExpireAt: time.Now().Add(ttl)}, 1)
	}

	return nil
}

func (s *Store) Del(key []byte) error {
	s.mem.Del(key)

	return s.rdb.Del(string(key)).Err()
}

func (s Store) Serialize(item *Result) ([]byte, error) {
	if item == nil {
		return nil, nil
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(*item)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
