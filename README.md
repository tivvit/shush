# shush 🤫

**S**elf **H**osted **U**rl **Sh**ortener

![Go](https://github.com/tivvit/shush/workflows/Go/badge.svg?branch=master)

the API is not stable yet. It will be stabilized soon and some real docs will follow ;)

## Config

The config file should be `config.{yml,json,ini,...}` (all formats provided by https://github.com/spf13/viper#what-is-viper) and it is expected in `/etc/shush` or in the same location as shush executable.

The config file location may also be set as param `./shush -confFile conf-test.yml`

### Config options (shown in `yaml` format)
```yaml
log:
  level: Info # default, using levels from https://github.com/sirupsen/logrus#level-logging
```
```yaml
server:
  address: 127.0.0.1:8080 # default, server interface address
```
```yaml
api:
  address: 127.0.0.1:8081 # default, API server interface address
```
Only one backend may be configured
```yaml
backend:
  in-mem: # data stored only in memory, this is not a persistent storage
  json-file:
    path: file.json # path to file which will store values encoded as json, this backend should be used for testing only. It is not suitable for production deployment
  redis:
    address: 127.0.0.1:6379 # address and port of the redis instance
  badger: # default
    path: badger # default, path to badger db 
```
Only one cache may be configured (no default)
```yaml
cache:
  big-cache: # head to https://github.com/allegro/bigcache#custom-initialization for detailed explanation of config values
    life-window-sec: 600 # entry TTL
    shards-pow: 10 # optional, 2 ** n, number of cache shards
    clean-window-sec: 0 # optional, how often should garbage collection run
    max-entries-in-window: 600000 # optional
    max-entry-size-bytes: 500 # optional
    verbose: true # optional
    hard-max-cache-size-mb: 1024 # optional
  free-cache: # https://github.com/coocood/freecache
    expire-sec: 600 # entry TTL
    size-kb: 1048576 # max cache size in kB
    gc-percent: # optional, edit go runtime gc config
  lru-cache: # github.com/evamayerova/lrucache
    expire-sec: 600 # entry TTL
    max-elems: 10000000 # max number of elements (keys)
  fast-cache: # https://github.com/VictoriaMetrics/fastcache
    size-bytes: 1073741824 # max cache size in Bytes
  ristretto-cache: # https://github.com/dgraph-io/ristretto
    counters: 100000000 # should be bigger than max-cost (keeps track about access for eviction)
    max-cost: 10000000 # max number of elements (keys)
    metrics: false # optional, statistics for debugging
```
```yaml
shortener:
  valid-url-pattern: ^[a-zA-Z0-9\-\_]{1,50}$ # optional
  gen-url-pattern: [a-zA-Z0-9]{5} # optional
  gen-max-retries: 10 # optional
  default-shortener: generator # optional ["generator", "hash"]
  default-hash-algo: fnv32 # optional ["md5", "sha1", "sha256","sha512", "fnv32", "fnv32a",  "fnv64", "fnv64a",  "fnv128",  "fnv128a",  "adler32",  "crc32ieee", "crc64iso",  "crc64ecma"]
  default-len: 5 # optional, for hash and generator
  max-len: 50 # optional, for requests
  allowed-shorteners: # optional (only listed is supported), for requests
    generator: true
    hash: true
  allowed-hash-algo: # optional (only listed is supported), for requests
    md5: true
    sha1: true
    sha256: true
    sha512: true
    fnv32: true
    fnv32a: true
    fnv64: true
    fnv64a: true
    fnv128: true
    fnv128a: true
    adler32: true
    crc32ieee: true
    crc64iso: true
    crc64ecma: true
```

## Other
number of possible generated urls in default settings (26+26+10)**5 = 916132832
