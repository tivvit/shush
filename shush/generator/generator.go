package generator

import (
	"github.com/lucasjones/reggen"
	"github.com/tivvit/shush/shush/backend"
	"github.com/tivvit/shush/shush/model"
	"time"
)


type ShortUrlGenerator struct {
	generator *reggen.Generator
	backend *backend.ShushBackend
}

func NewShortUrlGenerator(pattern string, b *backend.ShushBackend) (*ShortUrlGenerator, error) {
	g, err := reggen.NewGenerator(pattern)
	if err != nil {
		return &ShortUrlGenerator{}, nil
	}
	return &ShortUrlGenerator{
		generator: g,
		backend: b,
	}, nil
}

func (sg ShortUrlGenerator) Generate(u model.Url) (model.Url, error) {
	for {
		// todo return when shortUrl is already present?
		sUrl := sg.generator.Generate(256)
		_, err := sg.backend.Get(sUrl)
		if err == nil {
			// record already exists
			continue
		}
		// todo only non-existent error should be valid here
		u.ShortUrl = sUrl
		// lock the record
		err = sg.backend.Set(sUrl, u, time.Duration(0))
		if err != nil {
			return u, err
		}
		return u, nil
	}
}
