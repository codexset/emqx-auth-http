package routers

import (
	"context"
	"github.com/kataras/iris/v12"
)

func (c *Routers) Super(ctx iris.Context) {
	username := ctx.FormValue("username")
	exists, err := c.redis.SIsMember(context.Background(), c.redisKey.ForSuper, username).Result()
	if err != nil || !exists {
		ctx.StatusCode(401)
		return
	}
	ctx.StatusCode(200)
}
