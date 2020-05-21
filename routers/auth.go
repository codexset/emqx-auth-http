package routers

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

type Auth struct {
	Username string
	Token    string
}

func (c *Routers) Auth(ctx iris.Context) {
	auth := &Auth{
		Username: ctx.FormValue("username"),
		Token:    ctx.FormValue("password"),
	}
	secret, err := c.redis.HGet(context.Background(), c.redisKey.ForAuth, auth.Username).Result()
	if err != nil {
		logrus.Info(err.Error())
		ctx.StatusCode(401)
		return
	}
	token, err := jwt.Parse(auth.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			ctx.StatusCode(401)
			return nil, nil
		}
		return []byte(secret), nil
	})
	if err != nil {
		ctx.StatusCode(401)
		return
	}
	if !token.Valid {
		ctx.StatusCode(401)
		return
	}
	ctx.StatusCode(200)
}
