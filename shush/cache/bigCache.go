package cache

import (
	"github.com/tivvit/shush/shush/backend"
	"github.com/tivvit/shush/shush/config/cache"
	"time"

	"github.com/allegro/bigcache"
	log "github.com/sirupsen/logrus"
)

type BigCache struct {
	backend backend.Backend
	cache   *bigcache.BigCache
}

func NewBigCache(b backend.Backend, config *cache.Big) *BigCache {
	bc := &BigCache{
		backend: b,
	}
	conf := bigcache.DefaultConfig(time.Duration(config.LifeWindowSec) * time.Second)
	if config.ShardsPow != nil {
		conf.Shards = 2 ** config.ShardsPow
	}
	if config.CleanWindowSec != nil {
		conf.CleanWindow = time.Duration(*config.CleanWindowSec) * time.Second
	}
	if config.MaxEntriesInWindow != nil {
		conf.MaxEntriesInWindow = *config.MaxEntriesInWindow
	}
	if config.MaxEntrySizeBytes != nil {
		conf.MaxEntrySize = *config.MaxEntrySizeBytes
	}
	if config.Verbose != nil {
		conf.Verbose = *config.Verbose
	}
	if config.HardMaxCacheSizeMb != nil {
		conf.HardMaxCacheSize = *config.HardMaxCacheSizeMb
	}
	cache, initErr := bigcache.NewBigCache(conf)
	if initErr != nil {
		log.Fatal(initErr)
	}
	bc.cache = cache
	return bc
}

func (bc *BigCache) Get(k string) (string, error) {
	if entry, err := bc.cache.Get(k); err == nil {
		return string(entry), nil
	}
	v, err := bc.backend.Get(k)
	if err != nil {
		return "", err
	}
	err = bc.cache.Set(k, []byte(v))
	if err != nil {
		log.Warn(err)
	}
	return v, nil
}
