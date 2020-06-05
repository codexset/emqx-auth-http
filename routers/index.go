package routers

import (
	"github.com/kataras/iris/v12"
)

func (c *Routers) Index(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"tip":     "You know, for authorization",
		"version": "1.0.0",
	})
}
