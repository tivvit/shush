package cache

import (
	"github.com/coocood/freecache"
	log "github.com/sirupsen/logrus"
	"github.com/tivvit/shush/shush/backend"
	"github.com/tivvit/shush/shush/config"
	"runtime/debug"
)

type FreeCache struct {
	backend       backend.Backend
	cache         *freecache.Cache
	expirationSec int
}

func NewFreeCache(backend backend.Backend, conf *config.FreeCache) *FreeCache {
	fc := &FreeCache{
		backend:       backend,
		expirationSec: conf.ExpireSec,
		cache:         freecache.NewCache(conf.CacheSizeKb),
	}
	if conf.GcPercent != nil {
		debug.SetGCPercent(*conf.GcPercent)
	}
	return fc
}

func (fc FreeCache) Get(k string) (string, error) {
	if entry, err := fc.cache.Get([]byte(k)); err == nil {
		return string(entry), nil
	}
	v, err := fc.backend.Get(k)
	if err != nil {
		return "", err
	}
	// todo support entry expiration
	err = fc.cache.Set([]byte(k), []byte(v), fc.expirationSec)
	if err != nil {
		log.Warn(err)
	}
	return v, nil
}
