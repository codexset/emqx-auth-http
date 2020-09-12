package main

import (
	"emqx-auth-http/router"
	"github.com/valyala/fasthttp"
	"os"
)

func main() {
	//DB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	//rc := redis.NewClient(&redis.Options{
	//	Addr:     os.Getenv("REDIS_HOST"),
	//	Password: os.Getenv("REDIS_PASSWORD"),
	//	DB:       DB,
	//})
	//_, err := rc.Ping(context.Background()).Result()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//rs := routers.New(
	//	rc,
	//	&routers.RedisKey{
	//		ForAuth:  os.Getenv("REDIS_KEY_FOR_AUTH"),
	//		ForSuper: os.Getenv("REDIS_KEY_FOR_SUPER"),
	//		ForAcl:   os.Getenv("REDIS_KEY_FOR_ACL"),
	//	},
	//)
	r := router.New()
	fasthttp.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r.HandleFastHTTP)
}
