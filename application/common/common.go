package common

import (
	"emqx-auth-http/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
)

type Dependency struct {
	fx.In

	Config *config.Config
	Redis  *redis.Client
}

func Handle(handlersFn interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if method, ok := handlersFn.(func(ctx *gin.Context) interface{}); ok {
			result := method(ctx)
			switch val := result.(type) {
			case bool:
				if val {
					ctx.Status(200)
				} else {
					ctx.Status(500)
				}
				break
			case error:
				ctx.String(401, val.Error())
				break
			}
		} else {
			ctx.Status(404)
		}
	}
}
