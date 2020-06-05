package routers

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

func (c *Routers) Acl(ctx iris.Context) {
	username := ctx.FormValue("username")
	topic := ctx.FormValue("topic")
	exists, err := c.redis.SIsMember(
		context.Background(),
		c.redisKey.ForAcl+":"+username,
		topic,
	).Result()
	if err != nil {
		ctx.StatusCode(401)
		logrus.Error(err.Error())
		return
	}
	if !exists {
		ctx.StatusCode(401)
		logrus.Error("acl does not exist")
	}
	ctx.StatusCode(200)
}
