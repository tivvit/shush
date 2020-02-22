package backend

type InMem struct {
}

type JsonFile struct {
	Path string `yaml:"path"`
}

type Redis struct {
	Address string `yaml:"address,omitempty"`
}

type Badger struct {
	Path string `yaml:"path"`
}

type Conf struct {
	InMem    *InMem    `yaml:"in-mem,omitempty"`
	JsonFile *JsonFile `yaml:"json-file,omitempty"`
	Redis    *Redis    `yaml:"redis,omitempty"`
	Badger   *Badger   `yaml:"badger,omitempty"`
}

