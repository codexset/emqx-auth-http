package main

import (
	"context"
	"emqx-auth-http/bootstrap"
	"emqx-auth-http/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			bootstrap.LoadConfiguration,
			bootstrap.InitializeRedis,
		),
		fx.Invoke(
			Mock,
		),
	).Done()
}

func Mock(cfg *config.Config, redis *redis.Client) {
	key := cfg.Key
	ctx := context.Background()
	redis.HMSet(ctx, key.Auth, map[string]interface{}{
		"zZ11v3G6DrqjeMSo": "q!%QlIXvXNpZ1bPe",
		"QiqhdD3wKWJE6rgK": "XuEUnzEjCTz3*&nE",
	})
	redis.SAdd(ctx, key.Super, "zZ11v3G6DrqjeMSo")
	redis.SAdd(ctx, key.AclKey("QiqhdD3wKWJE6rgK"), "notice", "tests")
}
