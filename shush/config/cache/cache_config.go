package cache

type Big struct {
	LifeWindowSec      int   `mapstructure:"life-window-sec"`
	ShardsPow          *int  `mapstructure:"shards-pow,omitempty"`
	CleanWindowSec     *int  `mapstructure:"clean-window-sec,omitempty"`
	MaxEntriesInWindow *int  `mapstructure:"max-entries-in-window,omitempty"`
	MaxEntrySizeBytes  *int  `mapstructure:"max-entry-size-bytes,omitempty"`
	Verbose            *bool `mapstructure:"verbose,omitempty"`
	HardMaxCacheSizeMb *int  `mapstructure:"hard-max-cache-size-mb,omitempty"`
}

type Free struct {
	CacheSizeKb int  `mapstructure:"size-kb"`
	ExpireSec   int  `mapstructure:"expire-sec"`
	GcPercent   *int `mapstructure:"gc-percent,omitempty"`
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
	Metrics  *bool `mapstructure:"metrics,omitempty"`
}

type Conf struct {
	BigCache       *Big       `mapstructure:"big-cache"`
	FreeCache      *Free      `mapstructure:"free-cache"`
	LruCache       *Lru       `mapstructure:"lru-cache"`
	FastCache      *Fast      `mapstructure:"fast-cache"`
	RistrettoCache *Ristretto `mapstructure:"ristretto-cache"`
}
