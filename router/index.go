package router

import "github.com/valyala/fasthttp"

func (c *Router) index(ctx *fasthttp.RequestCtx) {
	ctx.SetBody([]byte(`You know, for authorization`))
}
