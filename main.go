package main

import (
	"errors"
	"flag"
	"github.com/dgraph-io/badger"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"github.com/tivvit/shush/shush-api"
	"github.com/tivvit/shush/shush/backend"
	"github.com/tivvit/shush/shush/cache"
	"github.com/tivvit/shush/shush/config"
	backendConf "github.com/tivvit/shush/shush/config/backend"
	cacheConf "github.com/tivvit/shush/shush/config/cache"
	"github.com/tivvit/shush/shush/generator"
	"github.com/valyala/fasthttp"
	"net/http"
	"sync"
)

var b *cache.ShushCache

func main() {
	confFile := flag.String("confFile", "conf.yml", "Configuration file path")
	flag.Parse()
	cfg, err := config.NewConf(*confFile)
	if err != nil {
		log.Fatal(err)
	}
	setupLogger(cfg.Log)
	bck, err := initBackend(cfg.Backend)
	if err != nil {
		log.Fatal(err)
	}
	cach, err := initCache(bck, cfg.Cache)
	if err != nil {
		log.Warn(err)
	}
	b = cache.NewShushCache(cach)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		log.Infof("starting server at %s", cfg.Server.Address)
		err = fasthttp.ListenAndServe(cfg.Server.Address, fastHTTPHandler)
		if err != nil {
			log.Error(err)
		}
		wg.Done()
	}()
	wg.Add(1)
	sb := backend.NewShushBackend(bck)
	g, err := generator.NewShortUrlGenerator(cfg.GenUrlPattern, sb)
	if err != nil {
		log.Fatal(err)
	}
	shush_api.SetBackend(sb)
	shush_api.SetGenerator(g)
	go func() {
		log.Printf("API Server starting at %s", cfg.Api.Address)
		log.Fatal(http.ListenAndServe(cfg.Api.Address, shush_api.NewRouter()))
		wg.Done()
	}()
	wg.Wait()
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	short := string(ctx.Path()[1:])
	url, err := b.Get(short)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	} else {
		ctx.Redirect(url.Target, fasthttp.StatusFound)
	}
}

func setupLogger(c config.Log) {
	lvl, err := log.ParseLevel(c.Level)
	if err != nil {
		log.Error(err)
	}
	log.SetLevel(lvl)
}

func initBackend(bc backendConf.Conf) (backend.Backend, error) {
	if bc.InMem != nil {
		return backend.NewInMem(), nil
	}
	if bc.JsonFile != nil {
		return backend.NewJsonFile(bc.JsonFile.Path)
	}
	if bc.Redis != nil {
		return backend.NewRedis(&redis.Options{
			Addr: bc.Redis.Address,
		}), nil
	}
	if bc.Redis != nil {
		return backend.NewRedis(&redis.Options{
			Addr: bc.Redis.Address,
		}), nil
	}
	if bc.Badger != nil {
		return backend.NewBadger(badger.DefaultOptions(bc.Badger.Path))
	}
	return nil, errors.New("unknown backend")
}

func initCache(b backend.Backend, cc *cacheConf.Conf) (cache.Cache, error) {
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
