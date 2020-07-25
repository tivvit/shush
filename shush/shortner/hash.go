package shortner

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"math"
)

func Hash(url []byte, hash string, ln int) (string, error) {
	var hv []byte
	switch hash {
	case "md5":
		hvs := md5.Sum(url)
		hv = hvs[:]
	case "sha1":
		hvs := sha1.Sum(url)
		hv = hvs[:]
	case "sha256":
		hvs := sha256.Sum256(url)
		hv = hvs[:]
	case "sha512":
		hvs := sha512.Sum512(url)
		hv = hvs[:]
	case "fnv32":
		hvs := fnv.New32().Sum(url)
		hv = hvs[:]
	case "fnv32a":
		hvs := fnv.New32a().Sum(url)
		hv = hvs[:]
	case "fnv64":
		hvs := fnv.New64().Sum(url)
		hv = hvs[:]
	case "fnv64a":
		hvs := fnv.New64a().Sum(url)
		hv = hvs[:]
	case "fnv128":
		hvs := fnv.New128().Sum(url)
		hv = hvs[:]
	case "fnv128a":
		hvs := fnv.New128a().Sum(url)
		hv = hvs[:]
	case "adler32":
		ad := adler32.New()
		hv = ad.Sum(url)
	case "crc32ieee":
		crc := crc32.NewIEEE()
		hv = crc.Sum(url)
	case "crc64iso":
		crc := crc64.New(crc64.MakeTable(crc64.ISO))
		hv = crc.Sum(url)
	case "crc64ecma":
		crc := crc64.New(crc64.MakeTable(crc64.ECMA))
		hv = crc.Sum(url)
	default:
		return "", errors.New("unknown hash func")
	}
	hx := hex.EncodeToString(hv)
	return hx[:int(math.Min(float64(len(hx)), float64(ln)))], nil
}
