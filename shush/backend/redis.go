package backend

import (
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"time"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(ro *redis.Options) *Redis {
	r := &Redis{}
	r.client = redis.NewClient(ro)
	return r
}

func (r Redis) Get(key string) (string, error) {
	val, err := r.client.Get(key).Result()
	if err != nil && err != redis.Nil {
		log.Warn(err)
	}
	return val, err
}

func (r Redis) GetAll() (map[string]string, error) {
	val, err := r.client.Keys("*").Result()
	if err != nil {
		return nil, err
	}
	m := map[string]string{}
	for _, k := range val {
		v, err := r.Get(k)
		if err != nil {
			log.Warn(err)
		}
		m[k] = v
	}
	return m, nil
}

func (r Redis) Set(key string, value string, ttl int) {
	err := r.client.Set(key, value, time.Duration(ttl)).Err()
	if err != nil {
		panic(err)
	}
}
