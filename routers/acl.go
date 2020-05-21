package routers

import "github.com/kataras/iris/v12"

func (c *Routers) Acl(ctx iris.Context) {
	ctx.StatusCode(200)
}
