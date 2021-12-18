package controllers

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func responseError(ctx *fasthttp.RequestCtx, code int, message string) {
	ctx.SetStatusCode(code)
	fmt.Fprintf(ctx, message)
}
