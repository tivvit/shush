package backend

import (
	"errors"
)

type InMem struct {
	data map[string]string
}

func NewInMem() *InMem {
	return &InMem{}
}

func (im InMem) Get(key string) (string, error) {
	if v, ok := im.data[key]; ok {
		return v, nil
	} else {
		return "", errors.New("key not found")
	}
}

func (im InMem) GetAll() (map[string]string, error) {
	return im.data, nil
}
