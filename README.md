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
  todo cache
```
todo shortener 

## Other
number of possible generated urls in default settings (26+26+10)**5 = 916132832
