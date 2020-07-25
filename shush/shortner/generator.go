package shortner

import (
	"github.com/lucasjones/reggen"
)

type ShortUrlGenerator struct {
	generator *reggen.Generator
}

func NewShortUrlGenerator(pattern string) (*ShortUrlGenerator, error) {
	g, err := reggen.NewGenerator(pattern)
	if err != nil {
		return &ShortUrlGenerator{}, nil
	}
	return &ShortUrlGenerator{
		generator: g,
	}, nil
}

func (sg ShortUrlGenerator) Generate(ln int) (string, error) {
	// todo max len should be server config
	maxLen := 256
	r := sg.generator.Generate(maxLen)
	if ln > len(r) {
		return "", &RequestedLongerThanGeneratedError{}
	}
	return r[:ln], nil
}

type RequestedLongerThanGeneratedError struct {}

func (e *RequestedLongerThanGeneratedError) Error() string {
	return "requested len is bigger then generated len"
}

