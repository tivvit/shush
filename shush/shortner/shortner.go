// package shortner generates short url based on target url and given params
package shortner

import (
	"errors"
	"fmt"
	"github.com/tivvit/shush/shush/backend"
	"github.com/tivvit/shush/shush/model"
)

type Shortner struct {
	gen  *ShortUrlGenerator
	b    *backend.ShushBackend
}

// todo will accept conf object
func NewShortner(b *backend.ShushBackend, pattern string) (*Shortner, error) {
	shortUrlGen, err := NewShortUrlGenerator(pattern)
	if err != nil {
		return &Shortner{}, err
	}
	return &Shortner{
		gen: shortUrlGen,
		b: b,
	}, nil
}

func checkShorturlEmpty(u *model.Url) error {
	if u.ShortUrl != "" {
		return errors.New("short_url already present")
	}
	return nil
}

// Hash updates short_url param in-place and store the result in the backend
func (s Shortner) Hash(u *model.Url, fn string, ln int) error {
	err := checkShorturlEmpty(u)
	if err != nil {
		return err
	}
	h, err := Hash([]byte(u.Target), fn, ln)
	if err != nil {
		return err
	}
	if _, err := s.b.Get(h); err == nil {
		return fmt.Errorf("already present %s", h)
	}
	// todo only non-existent error should be valid here
	u.ShortUrl = h
	return s.b.SetUnique(h, *u, u.Expires())
}

// Random updates short_url param in-place based on random pattern and store the result in the backend
func (s Shortner) Random(u *model.Url, ln int) error {
	err := checkShorturlEmpty(u)
	if err != nil {
		return err
	}
	// todo configurable retry
	maxRetry := 10
	for i := 0; i < maxRetry; i++ {
		// todo from conf / request
		sUrl, err := s.gen.Generate(ln)
		if err != nil {
			// this error is not recoverable
			return err
		}
		if _, err := s.b.Get(sUrl); err == nil {
			// record already exists
			continue
		}
		// todo only non-existent error should be valid here
		u.ShortUrl = sUrl
		err = s.b.SetUnique(sUrl, *u, u.Expires())
		if err != nil {
			continue
		}
		return nil
	}
	return fmt.Errorf("tried %d times and did not find unique short_url (pattern is probably almost exhausted)", maxRetry)
}