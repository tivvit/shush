package shortner

import (
	"regexp"
	"testing"
)

func TestShortUrlGenerator(t *testing.T) {
	testPattern(t, "[a-z]{5}", 5, 5, "^[a-z]{5}$")
	testPattern(t, "[a-zA-Z]{5}", 5, 5, "^[a-zA-Z]{5}$")
	testPattern(t, "[a-zA-Z0-9]{5}", 5, 5, "^[a-zA-Z0-9]{5}$")
	testPattern(t, "[a-zA-Z0-9]{42}", 42, 42, "^[a-zA-Z0-9]{42}$")
	testPattern(t, ".{42}", 42, 42, "^.{42}$")
	testPattern(t, "[a-zA-Z0-9]{10}", 5, 5, "^[a-zA-Z0-9]{5}$")
	testPattern(t, "[a-z]{10}a{10}", 20, 20, "^[a-zA-Z0-9]{20}$")
	g, err := NewShortUrlGenerator("[a-zA-Z0-9]{5}")
	if err != nil {
		t.Error(err)
	}
	_, err = g.Generate(10)
	if err == nil {
		t.Error("Created string longer than possible based on pattern")
	}
	if _, ok := err.(*RequestedLongerThanGeneratedError); !ok {
		t.Errorf("Uexpected err %s", err)
	}
}

func testPattern(t *testing.T, p string, ln int, expectedLen int, hasToMatch string) {
	g, err := NewShortUrlGenerator(p)
	if err != nil {
		t.Error(err)
	}
	gs, err := g.Generate(ln)
	if err != nil {
		t.Error(err)
	}
	if len(gs) != expectedLen {
		t.Errorf("incorrect generated len %d", len(gs))
	}
	m, err := regexp.MatchString(hasToMatch, gs)
	if err != nil {
		t.Error(err)
	}
	if !m {
		t.Error("generated text does not match the pattern")
	}
}
