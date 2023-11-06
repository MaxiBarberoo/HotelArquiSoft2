package cache

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

var (
	cacheClient *memcache.Client
)

func InitCache() {
	cacheClient = memcache.New("localhost:11211")
}

func Get(key string) []byte {
	item, err := cacheClient.Get(key)
	if err != nil {
		return nil
	}
	return item.Value
}

func Set(key string, value []byte) {
	if err := cacheClient.Set(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: 10, // 10 segundos de TTL
	}); err != nil {
		fmt.Println("Error setting item in cache", err)
	}
}
