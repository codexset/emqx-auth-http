package router

import (
	"context"
	"github.com/valyala/fasthttp"
)

func (c *Router) super(ctx *fasthttp.RequestCtx) {
	username := string(ctx.FormValue("username"))
	exists, err := c.redis.SIsMember(context.Background(), c.key.ForSuper, username).Result()
	if err != nil {
		ctx.Error(err.Error(), 401)
		return
	}
	if !exists {
		ctx.Error("super does not exist", 401)
		return
	}
	ctx.SetStatusCode(200)
}
