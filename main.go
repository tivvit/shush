package main

import (
	"github.com/tivvit/shush/shush/backend"
	"github.com/valyala/fasthttp"
)

var b backend.Backend

func main() {
	b = backend.NewJsonFile("urls.json")

	fasthttp.ListenAndServe(":80", fastHTTPHandler)
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	short := string(ctx.Path()[1:])
	url, err := b.Get(short)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	} else {
		ctx.Redirect(url, fasthttp.StatusFound)
	}
}
