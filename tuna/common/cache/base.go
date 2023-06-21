package cache

import (
	"fmt"
	"time"

	r "github.com/dgraph-io/ristretto"
)

type ICache interface {
	Get(key string, dest interface{}) error
	Set(key string, value interface{}, ttl time.Duration) error
}

func CreateCache() ICache {
	db, _ := r.NewCache(&r.Config{
		NumCounters: 10000,   // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      //
	})

	return inMemCache{
		DB: db,
	}
}

type EntryNotFound struct{ Key string }

func (m *EntryNotFound) Error() string {
	return fmt.Sprintf("%s notfound", m.Key)
}
