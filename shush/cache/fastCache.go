package cache

import (
	"github.com/VictoriaMetrics/fastcache"
	"github.com/tivvit/shush/shush/backend"
	"github.com/tivvit/shush/shush/config"
)

type FastCache struct {
	backend backend.Backend
	cache   *fastcache.Cache
}

func NewFastCache(backend backend.Backend, conf *config.FastCache) *FastCache {
	return &FastCache{
		backend: backend,
		cache:   fastcache.New(conf.MaxBytes),
	}
}

func (fc FastCache) Get(k string) (string, error) {
	e := fc.cache.Get(nil, []byte(k))
	if e != nil {
		return string(e), nil
	}
	v, err := fc.backend.Get(k)
	if err != nil {
		return "", err
	}
	fc.cache.Set([]byte(k), []byte(v))
	return v, nil
}
