package routers

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

func (c *Routers) Super(ctx iris.Context) {
	username := ctx.FormValue("username")
	exists, err := c.redis.SIsMember(context.Background(), c.redisKey.ForSuper, username).Result()
	if err != nil {
		ctx.StatusCode(401)
		logrus.Error(err.Error())
		return
	}
	if !exists {
		ctx.StatusCode(401)
		logrus.Error("super does not exist")
	}
	ctx.StatusCode(200)
}
