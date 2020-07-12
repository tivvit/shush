package backend

import "time"

type Backend interface {
	Get(string) (string, error)
	GetAll() (map[string]string, error)
	Set(string, string, time.Duration) error
	Close() error
}