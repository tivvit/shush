package cache

type Cache interface {
	Get(key string) (string, error)
}
