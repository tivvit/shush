package shortner

type Conf struct {
	GenUrlPattern     string          `mapstructure:"gen-url-pattern"`
	GenMaxRetries     int             `mapstructure:"gen-max-retries"`
	DefaultShortener  string          `mapstructure:"default-shortener"`
	DefaultHashAlgo   string          `mapstructure:"default-hash-algo"`
	DefaultLen        int             `mapstructure:"default-len"`
	AllowedShorteners map[string]bool `mapstructure:"allowed-shorteners"`
	AllowedHashAlgo   map[string]bool `mapstructure:"allowed-hash-algo"`
}
