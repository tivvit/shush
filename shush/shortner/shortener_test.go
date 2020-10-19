package shortner

import (
	"github.com/tivvit/shush/shush/backend"
	"github.com/tivvit/shush/shush/config/shortner"
	"testing"
)

func TestShortnerValidator(t *testing.T) {
	s, _ := NewShortner(backend.NewShushBackend(backend.NewInMem()), shortner.Conf{ValidUrlPattern: `^[a-zA-Z0-9\-\_]{1,50}$`})
	if s.IsValidShort("") {
		t.Error("allowed empty url")
	}
	if s.IsValidShort("!") {
		t.Error("allowed invalid url `!`")
	}
	if s.IsValidShort("a@b") {
		t.Error("allowed invalid url `a@b`")
	}
	if s.IsValidShort("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa") {
		t.Error("allowed invalid url `aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa` (too long)")
	}
	if !s.IsValidShort("a") {
		t.Error("not-allowed url `a`")
	}
	if !s.IsValidShort("aaa") {
		t.Error("not-allowed url `aaa`")
	}
	if !s.IsValidShort("abcdefABCDEF") {
		t.Error("not-allowed url `abcdefABCDEF`")
	}
}