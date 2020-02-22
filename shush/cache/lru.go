package cache

import (
	"github.com/evamayerova/lrucache"
	log "github.com/sirupsen/logrus"
	"github.com/tivvit/shush/shush/backend"
	"github.com/tivvit/shush/shush/config/cache"
)

type Lru struct {
	backend backend.Backend
	cache   *lrucache.Cache
	ttlSec  int
}

func NewLru(b backend.Backend, conf *cache.Lru) *Lru {
	return &Lru{
		backend: b,
		cache: lrucache.NewCache(conf.MaxElems),
		ttlSec: conf.ExpireSec,
	}
}

func (lru Lru) Get(k string) (string, error) {
	entry := lru.cache.Read(k)
	if entry != nil {
		return entry.(string), nil
	}
	v, err := lru.backend.Get(k)
	if err != nil {
		return "", err
	}
	// todo support entry expiration
	err = lru.cache.Write(k, v, lru.ttlSec)
	if err != nil {
		log.Warn(err)
	}
	return v, nil
}
