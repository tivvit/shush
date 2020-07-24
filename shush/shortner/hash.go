package shortner

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"github.com/tivvit/shush/shush/backend"
	"hash/adler32"
	"hash/crc32"
	"hash/fnv"
	"math"
	"strconv"
)

type Hasher struct {
	//hash    hash.Hash
	backend *backend.ShushBackend
	len     int
}

func NewHasher(len int, bck *backend.ShushBackend) *Hasher {
	return &Hasher{
		backend: bck,
		len:     len,
	}
}

func (h *Hasher) Hash(hash string, url []byte) (string, error) {
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
	case "adler32":
		return strconv.FormatUint(uint64(adler32.Checksum(url)), 16), nil
	case "crc32":
		return strconv.FormatUint(uint64(crc32.Checksum(url, crc32.MakeTable(0xD5828281))), 16), nil
	default:
		return "", errors.New("unknown hasher")
	}
	return hex.EncodeToString(hv)[:int(math.Min(float64(len(hv)), float64(h.len)))], nil
}
