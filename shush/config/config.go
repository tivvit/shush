package config

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tivvit/shush/shush/config/backend"
	"github.com/tivvit/shush/shush/config/cache"
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
	Log           Log
	Server        Server
	Api           Server
	Backend       backend.Conf
	Cache         *cache.Conf
	GenUrlPattern string `mapstructure:"gen-url-pattern"`
}

func NewConf(fn string) (*Conf, error) {
	if fn == "" {
		viper.SetConfigName("config")      // name of config file (without extension)
		viper.AddConfigPath("/etc/shush/") // path to look for the config file in
		viper.AddConfigPath(".")           // optionally look for config in the working directory
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// todo reasonable defaults
				// Config file not found; ignore error if desired
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
	}

	viper.SetDefault("log.level", "Info")
	viper.SetDefault("server.address", "127.0.0.1:8080")
	viper.SetDefault("api.address", "127.0.0.1:8081")
	viper.SetDefault("gen-url-pattern", "[a-zA-Z0-9]{5}")
	// todo
	// default backend (badger)
	viper.SetDefault("backend.in-mem", backend.InMem{})
	// default cache (big-cache? - based on tests)

	// todo
	//viper.WatchConfig()

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.SetEnvPrefix("shush")
	viper.AutomaticEnv()

	// todo support flags?
	//var cmd = &cobra.Command{
	//	Use:   "viper-test",
	//	Short: "testing viper",
	//	Run: func(command *cobra.Command, args []string) {
	//		fmt.Printf("thing1: %q\n", viper.GetString("thing1"))
	//		fmt.Printf("thing2: %q\n", viper.GetString("thing2"))
	//	},
	//}
	//
	//viper.AutomaticEnv()
	//flags := cmd.Flags()
	//flags.String("thing1", "", "The first thing")
	//viper.BindPFlag("thing1", flags.Lookup("thing1"))
	//flags.String("thing2", "", "The second thing")
	//viper.BindPFlag("thing2", flags.Lookup("thing2"))
	//
	//cmd.Execute()

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
