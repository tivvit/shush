package main

import (
	"errors"
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/tivvit/shush/shush/backend"
	"github.com/tivvit/shush/shush/config"
	"github.com/valyala/fasthttp"
)

var b backend.Backend

func main() {
	confFile := flag.String("confFile", "conf.yml", "Configuration file path")
	flag.Parse()
	c, err := config.NewConf(*confFile)
	if err != nil {
		log.Fatal(err)
	}
	setupLogger(c.Log)
	b, err = initBackend(c.Backend)
	if err != nil {
		log.Fatal(err)
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

func setupLogger(c config.LogConf) {
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
	return nil, errors.New("unknown backend")
}
