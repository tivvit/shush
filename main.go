package main

import (
	"errors"
	"flag"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"github.com/tivvit/shush/shush/backend"
	"github.com/tivvit/shush/shush/cache"
	"github.com/tivvit/shush/shush/config"
	"github.com/valyala/fasthttp"
)

var b cache.Cache

func main() {
	confFile := flag.String("confFile", "conf.yml", "Configuration file path")
	flag.Parse()
	c, err := config.NewConf(*confFile)
	if err != nil {
		log.Fatal(err)
	}
	setupLogger(c.Log)
	bck, err := initBackend(c.Backend)
	if err != nil {
		log.Fatal(err)
	}
	b, err = initCache(bck, c.Cache)
	if err != nil {
		log.Warn(err)
	}
	err = fasthttp.ListenAndServe(c.Server.Address, fastHTTPHandler)
	if err != nil {
		log.Error(err)
	}
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	short := string(ctx.Path()[1:])
	url, err := b.Get(short)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	} else {
		ctx.Redirect(url, fasthttp.StatusFound)
	}
}

func setupLogger(c config.Log) {
	lvl, err := log.ParseLevel(c.Level)
	if err != nil {
		log.Error(err)
	}
	log.SetLevel(lvl)
}

func initBackend(bc config.BackendConf) (backend.Backend, error) {
	if bc.InMem != nil {
		return backend.NewInMem(), nil
	}
	if bc.JsonFile != nil {
		return backend.NewJsonFile(bc.JsonFile.Path), nil
	}
	if bc.Redis != nil {
		return backend.NewRedis(&redis.Options{
			Addr: bc.Redis.Address,
		}), nil
	}
	return nil, errors.New("unknown backend")
}

func initCache(b backend.Backend, cc *config.CacheConf) (cache.Cache, error) {
	if cc == nil {
		return b, nil
	}
	if cc.BigCache != nil {
		return cache.NewBigCache(b, cc.BigCache), nil
	}
	if cc.FreeCache != nil {
		return cache.NewFreeCache(b, cc.FreeCache), nil
	}
	if cc.LruCache != nil {
		return cache.NewLru(b, cc.LruCache), nil
	}
	if cc.FastCache != nil {
		return cache.NewFastCache(b, cc.FastCache), nil
	}
	if cc.RistrettoCache != nil {
		return cache.NewRistrettoCache(b, cc.RistrettoCache), nil
	}
	return b, errors.New("unknown cache")
}
