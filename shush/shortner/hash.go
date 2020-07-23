package shortner

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/tivvit/shush/shush/backend"
	"hash"
)

type Hasher struct {
	hash    hash.Hash
	backend *backend.ShushBackend
	len     int
}

func NewHasher(hash string, len int, bck *backend.ShushBackend) (*Hasher, error) {
	switch hash {
	case "md5":
		return &Hasher{
			hash:    md5.New(),
			backend: bck,
			len:     len,
		}, nil
	default:
		return nil, errors.New("unknown hasher")
	}

}

func (h *Hasher) Hash(url string) string {
	return hex.EncodeToString(h.hash.Sum([]byte("Sum returns bytes"))[:])[:h.len]
}
