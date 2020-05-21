package routers

import "github.com/go-redis/redis/v8"

type Routers struct {
	redis    *redis.Client
	redisKey *RedisKey
}

type RedisKey struct {
	ForAuth  string
	ForSuper string
	ForAcl   string
}

func New(redis *redis.Client, redisKey *RedisKey) *Routers {
	routers := new(Routers)
	routers.redis = redis
	routers.redisKey = redisKey
	return routers
}
