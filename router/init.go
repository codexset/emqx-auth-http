package router

import "github.com/valyala/fasthttp"

type Router struct {
}

func New() *Router {
	c := new(Router)
	return c
}

func (c *Router) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		c.index(ctx)
	case "/auth":
		c.auth(ctx)
	case "/super":
		c.super(ctx)
	case "/acl":
		c.acl(ctx)
	default:
		ctx.Error("Request does not exist", fasthttp.StatusNotFound)
	}
}
