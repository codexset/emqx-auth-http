package routers

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

type Auth struct {
	Username string
	Password string
}

func (c *Routers) Auth(ctx iris.Context) {
	auth := &Auth{
		Username: ctx.FormValue("username"),
		Password: ctx.FormValue("password"),
	}
	result, err := c.redis.HGet(context.Background(), c.redisKey.ForAuth, auth.Username).Result()
	if err != nil {
		ctx.StatusCode(401)
		return
	}
	var data map[string]interface{}
	err = jsoniter.Unmarshal([]byte(result), &data)
	if err != nil {
		ctx.StatusCode(401)
		return
	}
	logrus.Info(data["secret"])
	ctx.StatusCode(200)
}
