package shortner

import (
	"fmt"
	"github.com/tivvit/shush/shush/backend"
	"testing"
)

func TestShortner(t *testing.T) {
	b := backend.NewShushBackend(backend.NewInMem())
	//h, err := NewHasher("md5", 20, b)
	//if err != nil {
	//	t.Error(err)
	//}
	//fmt.Println(h.Hash([]byte("aaaaaa")))
	h := NewHasher( 30, b)
	//hv, err := h.Hash("md5", []byte("http://ahoj"))
	hv, err := h.Hash("crc32", []byte("http://ahoj"))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hv)
}
