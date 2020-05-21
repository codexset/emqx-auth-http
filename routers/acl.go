package routers

import (
	"context"
	"github.com/kataras/iris/v12"
)

func (c *Routers) Acl(ctx iris.Context) {
	username := ctx.FormValue("username")
	topic := ctx.FormValue("topic")
	exists, err := c.redis.SIsMember(
		context.Background(),
		c.redisKey.ForAcl+":"+username,
		topic,
	).Result()
	if err != nil || !exists {
		ctx.StatusCode(401)
		return
	}
	ctx.StatusCode(200)
}
