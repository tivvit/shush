package shortner

type Conf struct {
	ValidUrlPattern   string          `mapstructure:"valid-url-pattern"`
	GenUrlPattern     string          `mapstructure:"gen-url-pattern"`
	GenMaxRetries     int             `mapstructure:"gen-max-retries"`
	DefaultShortener  string          `mapstructure:"default-shortener"`
	DefaultHashAlgo   string          `mapstructure:"default-hash-algo"`
	DefaultLen        int             `mapstructure:"default-len"`
	Maxlen            int             `mapstructure:"max-len"`
	AllowedShorteners map[string]bool `mapstructure:"allowed-shorteners"`
	AllowedHashAlgo   map[string]bool `mapstructure:"allowed-hash-algo"`
}
