package backend

import (
	"github.com/sirupsen/logrus"
	"github.com/tivvit/shush/shush/model"
	"time"
)

type Backend interface {
	Get(string) (string, error)
	GetAll() (map[string]string, error)
	Set(string, string, time.Duration) error
	Remove(string) error
	Close() error
}

type ShushBackend struct {
	backend Backend
}

func NewShushBackend(b Backend) *ShushBackend {
	return &ShushBackend{
		backend: b,
	}
}

func (sb ShushBackend) Get(k string) (model.Url, error) {
	v, err := sb.backend.Get(k)
	if err != nil {
		return model.Url{}, err
	}
	return model.UrlDeserialize(v)
}

func (sb ShushBackend) GetRaw(k string) (string, error) {
	return sb.backend.Get(k)
}

func (sb ShushBackend) GetAll() (map[string]model.Url, error) {
	d, err := sb.backend.GetAll()
	if err != nil {
		return map[string]model.Url{}, err
	}
	r := make(map[string]model.Url, len(d))
	for k, v := range d {
		uv, err := model.UrlDeserialize(v)
		if err != nil {
			logrus.Error(err)
			continue
		}
		r[k] = uv
	}
	return r, nil
}

func (sb ShushBackend) Set(k string, v model.Url, d time.Duration) error {
	sv, err := model.UrlSerialize(v)
	if err != nil {
		return err
	}
	return sb.backend.Set(k, sv, d)
}


func (sb ShushBackend) Remove(key string) error {
	return sb.backend.Remove(key)
}

func (sb ShushBackend) Close() error {
	return sb.backend.Close()
}

