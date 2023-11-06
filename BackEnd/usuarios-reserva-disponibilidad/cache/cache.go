package cache

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var cache *memcache.Client

func Init_cache() {
	cache = memcache.New("locaclhost:11211")
}

func Set(key string, value []byte) {
	cache.Set(&memcache.Item{Key: key, Value: value})
}

func Get(key string) (value string) {
	it, err := cache.Get(key)
	if err != nil {
		return ""
	}
	return string(it.Value)
}
