package config

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tivvit/shush/shush/config/backend"
	"github.com/tivvit/shush/shush/config/cache"
	"github.com/tivvit/shush/shush/config/shortner"
	"reflect"
	"strings"
)

type Log struct {
	Level string
}

type Server struct {
	Address string
}

type Conf struct {
	Log       Log
	Server    Server
	Api       Server
	Backend   backend.Conf
	Cache     *cache.Conf
	Shortener shortner.Conf
}

func NewConf(fn string) (*Conf, error) {
	if fn == "" {
		viper.SetConfigName("config")      // name of config file (without extension)
		viper.AddConfigPath("/etc/shush/") // path to look for the config file in
		viper.AddConfigPath(".")           // optionally look for config in the working directory
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Info("No config file loaded, using defaults")
			} else {
				log.Fatalf("Fatal error config file: %s \n", err)
			}
		} else {
			log.Infof("Used config file %s", viper.ConfigFileUsed())
		}
	} else {
		viper.SetConfigFile(fn)
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Fatal error config file: %s \n", err)
		}
		log.Infof("Used config file %s", viper.ConfigFileUsed())
	}

	viper.SetDefault("log.level", "Info")
	viper.SetDefault("server.address", "127.0.0.1:8080")
	viper.SetDefault("api.address", "127.0.0.1:8081")
	viper.SetDefault("backend.badger", &backend.Badger{Path: "badger"})
	viper.SetDefault("shortener.gen-url-pattern", "[a-zA-Z0-9]{5}")
	viper.SetDefault("shortener.default-shortener", "generator")
	viper.SetDefault("shortener.default-hash-algo", "fnv32")
	viper.SetDefault("shortener.default-len", "5")
	viper.SetDefault("shortener.allowed-shorteners", map[string]bool{
		"generator": true,
		"hash":      true,
	})
	viper.SetDefault("shortener.allowed-hash-algo", map[string]bool{
		"md5":       true,
		"sha1":      true,
		"sha256":    true,
		"sha512":    true,
		"fnv32":     true,
		"fnv32a":    true,
		"fnv64":     true,
		"fnv64a":    true,
		"fnv128":    true,
		"fnv128a":   true,
		"adler32":   true,
		"crc32ieee": true,
		"crc64iso":  true,
		"crc64ecma": true,
	})

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.SetEnvPrefix("shush")
	viper.AutomaticEnv()

	var conf Conf

	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	err = conf.validate()
	if err != nil {
		return nil, err
	}

	return &conf, nil
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
