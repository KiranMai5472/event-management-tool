package cacheprovider

import (
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
)

//ALL cache function is dfined here

var myCache CacheItf
var Cacheset *cache.Cache

type CacheItf interface {
	Set(key string, DeviceName interface{}, expiration time.Duration) bool
	Get(key string) (interface{}, bool)
	Delete(key string) bool
	Close()
}

type AppCache struct {
	client *cache.Cache
}

// SetProvider sets specific in memory provider e.g, redis, mem cache etc...
func SetProvider(provider CacheItf) error {
	if provider == nil {
		return errors.New("provider can not be nil")
	}
	if myCache != nil {
		myCache.Close()
	}
	myCache = provider
	return nil
}

// GetMemoryStore returns the memory store with default as Cache provider
func GetMemoryStore() CacheItf {
	if myCache != nil {
		return myCache
	}
	myCache = NewCacheProvider()
	return myCache
}

func NewCacheProvider() *AppCache {
	return &AppCache{
		client: cache.New(24*time.Hour, 24*time.Hour),
	}
}
func (r *AppCache) Set(key string, DeviceName interface{}, expiration time.Duration) bool {

	r.client.Set(key, DeviceName, expiration)
	return true
}

func (r *AppCache) Get(key string) (interface{}, bool) {
	var catch interface{}
	var found bool
	data, found := r.client.Get(key)
	if found {
		catch = data
	}
	return catch, found
}

func (r *AppCache) Delete(key string) bool {
	r.client.Delete(key)
	r.client.DeleteExpired()
	return true
}

func (r *AppCache) Close() {
	r.Close()
}

func CacheGet() *cache.Cache {

	return Cacheset
}
