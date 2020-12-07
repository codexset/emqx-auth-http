package controller

import (
	"context"
	"emqx-auth-http/application/common"
	"emqx-auth-http/config/options"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	UserNotExists  = errors.New("username does not exist")
	SuperNotExists = errors.New("super does not exist")
	AclNotExists   = errors.New("acl does not exist")
)

type controller struct {
	Key options.Key
	*common.Dependency
}

func New(dep *common.Dependency) *controller {
	c := new(controller)
	c.Dependency = dep
	c.Key = dep.Config.Key
	return c
}

type AuthBody struct {
	Username string `form:"username" binding:"required"`
	Token    string `form:"token" binding:"required"`
}

func (c *controller) Auth(ctx *gin.Context) interface{} {
	var body AuthBody
	var err error
	if err = ctx.ShouldBind(&body); err != nil {
		return err
	}
	var secret string
	if secret, err = c.Redis.HGet(context.Background(), c.Key.Auth, body.Username).Result(); err != nil {
		return UserNotExists
	}
	var token *jwt.Token
	if token, err = jwt.Parse(body.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}); err != nil {
		return err
	}
	if !token.Valid {
		return jwt.ErrInvalidKey
	}
	return true
}

type SuperBody struct {
	Username string `form:"username" binding:"required"`
}

func (c *controller) Super(ctx *gin.Context) interface{} {
	var body SuperBody
	var err error
	if err = ctx.ShouldBind(&body); err != nil {
		return err
	}
	var exists bool
	if exists, err = c.Redis.SIsMember(context.Background(), c.Key.Super, body.Username).Result(); err != nil {
		return err
	}
	if !exists {
		return SuperNotExists
	}
	return true
}

type AclBody struct {
	Username string `form:"username" binding:"required"`
	Topic    string `form:"topic" binding:"required"`
}

func (c *controller) Acl(ctx *gin.Context) interface{} {
	var body AclBody
	var err error
	if err = ctx.ShouldBind(&body); err != nil {
		return err
	}
	var exists bool
	if exists, err = c.Redis.SIsMember(context.Background(), c.Key.AclKey(body.Username), body.Topic).Result(); err != nil {
		return err
	}
	if !exists {
		return AclNotExists
	}
	return true
}
