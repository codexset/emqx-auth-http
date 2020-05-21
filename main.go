package main

import (
	"emqx-auth-http/routers"
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"os"
	"strconv"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("fatal")
	app.Use(recover.New())
	app.Use(logger.New())
	DB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	rs := routers.New(
		redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       DB,
		}),
		&routers.RedisKey{
			ForAuth:  os.Getenv("REDIS_KEY_FOR_AUTH"),
			ForSuper: os.Getenv("REDIS_KEY_FOR_SUPER"),
			ForAcl:   os.Getenv("REDIS_KEY_FOR_ACL"),
		},
	)
	app.Get("/", rs.Index)
	app.Post("/auth", rs.Auth)
	app.Post("/super", rs.Super)
	app.Post("/acl", rs.Acl)
	app.Run(iris.Addr(":"+os.Getenv("SERVER_PORT")), iris.WithoutServerError(iris.ErrServerClosed))
}
