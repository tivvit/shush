package backend

type InMem struct{}

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
	InMem    *InMem    `yaml:"in-mem,omitempty" mapstructure:"in-mem"`
	JsonFile *JsonFile `yaml:"json-file,omitempty" mapstructure:"json-file"`
	Redis    *Redis    `yaml:"redis,omitempty" mapstructure:"redis"`
	Badger   *Badger   `yaml:"badger,omitempty" mapstructure:"badger"`
}
