package router

import (
	"github.com/go-redis/redis/v8"
	"github.com/valyala/fasthttp"
)

type Router struct {
	redis *redis.Client
	key   *Key
}

type Key struct {
	ForAuth  string
	ForSuper string
	ForAcl   string
}

func New(redis *redis.Client, key *Key) *Router {
	c := new(Router)
	c.redis = redis
	c.key = key
	return c
}

func (c *Router) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	if path == "/" {
		ctx.SetBody([]byte(`You know, for authorization`))
		return
	}
	if string(ctx.Method()) == "POST" {
		switch path {
		case "/auth":
			c.auth(ctx)
		case "/super":
			c.super(ctx)
		case "/acl":
			c.acl(ctx)
		}
		return
	}
	ctx.Error("Request does not exist", fasthttp.StatusNotFound)
}
