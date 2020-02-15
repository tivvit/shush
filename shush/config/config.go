package config

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

type LogConf struct {
	Level string `yaml:"level,omitempty"`
}

type ServerConf struct {
	Address string `yaml:"address,omitempty"`
}

type BackendInMemConf struct {
}

type BackendJsonFileConf struct {
	Path string `yaml:"path"`
}

type BackendConf struct {
	InMem    *BackendInMemConf    `yaml:"in-mem,omitempty"`
	JsonFile *BackendJsonFileConf `yaml:"json-file,omitempty"`
}

type Conf struct {
	Log     LogConf     `yaml:"log,omitempty"`
	Server  ServerConf  `yaml:"server,omitempty"`
	Backend BackendConf `yaml:"backend,omitempty"`
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
		conf.Backend.InMem = &BackendInMemConf{}
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

func (c *Conf) validate() error {
	backendCnt := c.numBackends()
	if backendCnt == 0 {
		return errors.New("no backend configured")
	}
	if backendCnt > 1 {
		return errors.New("more than one backend configured")
	}
	return nil
}
