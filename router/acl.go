package router

import (
	"context"
	"github.com/valyala/fasthttp"
)

func (c *Router) acl(ctx *fasthttp.RequestCtx) {
	username := string(ctx.FormValue("username"))
	topic := ctx.FormValue("topic")
	exists, err := c.redis.SIsMember(
		context.Background(),
		c.key.ForAcl+":"+username,
		topic,
	).Result()
	if err != nil {
		ctx.Error(err.Error(), 401)
		return
	}
	if !exists {
		ctx.Error("acl does not exist", 401)
		return
	}
	ctx.SetStatusCode(200)
}
