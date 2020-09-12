package router

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
)

type AuthData struct {
	Username string
	Token    string
}

func (c *Router) auth(ctx *fasthttp.RequestCtx) {
	data := &AuthData{
		Username: string(ctx.FormValue("username")),
		Token:    string(ctx.FormValue("password")),
	}
	secret, err := c.redis.HGet(context.Background(), c.key.ForAuth, data.Username).Result()
	if err != nil {
		ctx.Error(err.Error(), 401)
		return
	}
	token, err := jwt.Parse(data.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			ctx.Error(err.Error(), 401)
			return nil, nil
		}
		return []byte(secret), nil
	})
	if err != nil {
		ctx.Error(err.Error(), 401)
		return
	}
	if !token.Valid {
		ctx.Error("Token verification is incorrect", 401)
		return
	}
	ctx.SetStatusCode(200)
}
