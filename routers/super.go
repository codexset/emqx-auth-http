package routers

import "github.com/kataras/iris/v12"

func (c *Routers) Super(ctx iris.Context) {
	ctx.StatusCode(200)
}
