package backend

type InMem struct{}

type JsonFile struct {
	Path string
}

type Redis struct {
	Address string
}

type Badger struct {
	Path string
}

type Conf struct {
	InMem    *InMem    `mapstructure:"in-mem"`
	JsonFile *JsonFile `mapstructure:"json-file"`
	Redis    *Redis    `mapstructure:"redis"`
	Badger   *Badger   `mapstructure:"badger"`
}
