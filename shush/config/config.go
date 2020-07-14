package config

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/tivvit/shush/shush/config/backend"
	"github.com/tivvit/shush/shush/config/cache"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

type Log struct {
	Level string `yaml:"level,omitempty"`
}

type Server struct {
	Address string `yaml:"address,omitempty"`
}

type Conf struct {
	Log           Log          `yaml:"log,omitempty"`
	Server        Server       `yaml:"server,omitempty"`
	Api           Server       `yaml:"api,omitempty"`
	Backend       backend.Conf `yaml:"backend,omitempty"`
	Cache         *cache.Conf  `yaml:"cache,omitempty"`
	GenUrlPattern string       `yaml:"gen-url-pattern,omitempty"`
}

func NewConf(fn string) (*Conf, error) {
	conf := &Conf{}
	conf.defaults()
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Error(err)
	}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Error(err)
	}
	if conf.numBackends() == 0 {
		log.Info("No backend configured using in-mem")
		conf.Backend.InMem = &backend.InMem{}
	}
	err = conf.validate()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func (c *Conf) defaults() {
	c.Log.Level = "Info"
	c.Server.Address = "127.0.0.1:8080"
	c.GenUrlPattern = "[a-zA-Z0-9]{5}"
}

func (c *Conf) numBackends() int {
	backendCnt := 0
	val := reflect.ValueOf(c.Backend)
	for i := 0; i < val.Type().NumField(); i++ {
		log.Debug(val.Type().Field(i).Name, val.Field(i).Interface(), val.Field(i).Type())
		if !val.Field(i).IsNil() {
			backendCnt++
		}
	}
	return backendCnt
}

func (c *Conf) numCaches() int {
	cacheCnt := 0
	if c.Cache == nil {
		return cacheCnt
	}
	val := reflect.ValueOf(*c.Cache)
	for i := 0; i < val.Type().NumField(); i++ {
		log.Debug(val.Type().Field(i).Name, val.Field(i).Interface(), val.Field(i).Type())
		if !val.Field(i).IsNil() {
			cacheCnt++
		}
	}
	return cacheCnt
}

func (c *Conf) validate() error {
	backendCnt := c.numBackends()
	if backendCnt == 0 {
		return errors.New("no backend configured")
	}
	if backendCnt > 1 {
		return errors.New("more than one backend configured")
	}
	if c.numCaches() > 1 {
		return errors.New("more then one cache configured")
	}
	return nil
}
