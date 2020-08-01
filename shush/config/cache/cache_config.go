package cache

type Big struct {
	LifeWindowSec      int   `mapstructure:"life-window-sec"`
	ShardsPow          *int  `mapstructure:"shards-pow"`
	CleanWindowSec     *int  `mapstructure:"clean-window-sec"`
	MaxEntriesInWindow *int  `mapstructure:"max-entries-in-window"`
	MaxEntrySizeBytes  *int  `mapstructure:"max-entry-size-bytes"`
	Verbose            *bool `mapstructure:"verbose"`
	HardMaxCacheSizeMb *int  `mapstructure:"hard-max-cache-size-mb"`
}

type Free struct {
	CacheSizeKb int  `mapstructure:"size-kb"`
	ExpireSec   int  `mapstructure:"expire-sec"`
	GcPercent   *int `mapstructure:"gc-percent"`
}

type Lru struct {
	MaxElems  int `mapstructure:"max-elems"`
	ExpireSec int `mapstructure:"expire-sec"`
}

type Fast struct {
	MaxBytes int `mapstructure:"size-bytes"`
}

type Ristretto struct {
	Counters int64 `mapstructure:"counters"`
	MaxCost  int64 `mapstructure:"max-cost"`
	Metrics  *bool `mapstructure:"metrics"`
}

type Conf struct {
	BigCache       *Big       `mapstructure:"big-cache"`
	FreeCache      *Free      `mapstructure:"free-cache"`
	LruCache       *Lru       `mapstructure:"lru-cache"`
	FastCache      *Fast      `mapstructure:"fast-cache"`
	RistrettoCache *Ristretto `mapstructure:"ristretto-cache"`
}
