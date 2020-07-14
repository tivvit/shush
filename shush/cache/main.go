package cache

import (
	"github.com/tivvit/shush/shush/model"
)

type Cache interface {
	Get(key string) (string, error)
}

type ShushCache struct {
	cache Cache
}

func NewShushCache(c Cache) *ShushCache {
	return &ShushCache{
		cache: c,
	}
}

func (sc ShushCache) Get(k string) (model.Url, error) {
	v, err := sc.cache.Get(k)
	if err != nil {
		return model.Url{}, err
	}
	// todo there is probably no need for storing the not serialized value (maybe only for expiration)
	return model.UrlDeserialize(v)
}
