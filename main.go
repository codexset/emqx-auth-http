package main

import (
	"emqx-auth-http/router"
	"github.com/go-redis/redis/v8"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"strconv"
)

func main() {
	DB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalln(err)
	}
	r := router.New(
		redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       DB,
		}),
		&router.Key{
			ForAuth:  os.Getenv("REDIS_KEY_FOR_AUTH"),
			ForSuper: os.Getenv("REDIS_KEY_FOR_SUPER"),
			ForAcl:   os.Getenv("REDIS_KEY_FOR_ACL"),
		},
	)
	fasthttp.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r.HandleFastHTTP)
}
