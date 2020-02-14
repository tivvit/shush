package backend

type Backend interface {
	Get(string) (string, error)
	GetAll() (map[string]string, error)
}