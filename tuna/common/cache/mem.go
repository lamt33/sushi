package cache

import (
	"encoding/json"
	"time"

	//do not import this eleswhere
	r "github.com/dgraph-io/ristretto"
	"github.com/lamt3/sushi/tuna/common/logger"
)

// In Memory Cache
type inMemCache struct {
	DB *r.Cache
}

func (pc inMemCache) Get(key string, dest interface{}) error {
	v, ok := pc.DB.Get(key)
	if !ok {
		return &EntryNotFound{Key: key}
	}
	r := v.([]byte)
	return json.Unmarshal(r, dest)
}

// set ttl to 0 if no expiration
func (pc inMemCache) Set(key string, value interface{}, ttl time.Duration) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}

	ok := pc.DB.SetWithTTL(key, p, 1, ttl)
	if !ok {
		return logger.Error("could not set cache key: %s with value %+v", key, value)
	}
	return nil
}
