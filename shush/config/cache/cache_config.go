package cache

type Big struct {
	LifeWindowSec      int   `yaml:"life-window-sec"`
	ShardsPow          *int  `yaml:"shards-pow,omitempty"`
	CleanWindowSec     *int  `yaml:"clean-window-sec,omitempty"`
	MaxEntriesInWindow *int  `yaml:"max-entries-in-window,omitempty"`
	MaxEntrySizeBytes  *int  `yaml:"max-entry-size-bytes,omitempty"`
	Verbose            *bool `yaml:"verbose,omitempty"`
	HardMaxCacheSizeMb *int  `yaml:"hard-max-cache-size-mb,omitempty"`
}

type Free struct {
	CacheSizeKb int  `yaml:"size-kb"`
	ExpireSec   int  `yaml:"expire-sec"`
	GcPercent   *int `yaml:"gc-percent,omitempty"`
}

type Lru struct {
	MaxElems  int `yaml:"max-elems"`
	ExpireSec int `yaml:"expire-sec"`
}

type Fast struct {
	MaxBytes int `yaml:"size-bytes"`
}

type Ristretto struct {
	Counters int64 `yaml:"counters"`
	MaxCost  int64 `yaml:"max-cost"`
	Metrics  *bool `yaml:"metrics,omitempty"`
}

type Conf struct {
	BigCache       *Big       `yaml:"big-cache,omitempty"`
	FreeCache      *Free      `yaml:"free-cache,omitempty"`
	LruCache       *Lru       `yaml:"lru-cache,omitempty"`
	FastCache      *Fast      `yaml:"fast-cache,omitempty"`
	RistrettoCache *Ristretto `yaml:"ristretto-cache,omitempty"`
}
