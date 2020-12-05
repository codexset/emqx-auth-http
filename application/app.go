package application

import (
	"emqx-auth-http/application/common"
	"emqx-auth-http/application/controller"
	"github.com/gin-gonic/gin"
	_ "net/http/pprof"
)

func Application(router *gin.Engine, dep common.Dependency) (err error) {
	control := controller.New(&dep)
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "You know, for authorization")
	})
	router.POST("/auth", common.Handle(control.Auth))
	router.POST("/super", common.Handle(control.Super))
	router.POST("/acl", common.Handle(control.Acl))
	return
}
