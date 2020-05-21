package routers

import "github.com/kataras/iris/v12"

func (c *Routers) Index(ctx iris.Context) {
	ctx.Writef("You know, for authorization")
}
