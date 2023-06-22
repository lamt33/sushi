package cache

import (
	"fmt"
	"time"
)

type ICache interface {
	Get(key string, dest interface{}) error
	Set(key string, value interface{}, ttl time.Duration) error
	Close() error
}

type EntryNotFound struct{ Key string }

func (m *EntryNotFound) Error() string {
	return fmt.Sprintf("%s notfound", m.Key)
}
