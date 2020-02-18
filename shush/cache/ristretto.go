package cache

import (
	"github.com/dgraph-io/ristretto"
	log "github.com/sirupsen/logrus"
	"github.com/tivvit/shush/shush/backend"
	"github.com/tivvit/shush/shush/config"
)

type Ristretto struct {
	backend backend.Backend
	cache   *ristretto.Cache
}

func NewRistrettoCache(b backend.Backend, conf *config.RistrettoCache) *Ristretto {
	r := &Ristretto{
		backend: b,
	}
	rc := &ristretto.Config{
		NumCounters: conf.Counters,
		MaxCost:     conf.MaxCost,
		BufferItems: 64, // Number of keys per Get buffer. - Recommended value
		Metrics:     false,
	}
	if conf.Metrics != nil {
		rc.Metrics = *conf.Metrics
	}
	cache, err := ristretto.NewCache(rc)
	if err != nil {
		log.Fatal(err)
	}
	r.cache = cache
	return r
}

func (rs Ristretto) Get(k string) (string, error) {
	entry, found := rs.cache.Get(k)
	if found {
		return entry.(string), nil
	}
	v, err := rs.backend.Get(k)
	if err != nil {
		return "", err
	}
	rs.cache.Set(k, v, 1)
	return v, nil
}
