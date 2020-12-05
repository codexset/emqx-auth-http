package main

import (
	"emqx-auth-http/application"
	"emqx-auth-http/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.NopLogger,
		fx.Provide(
			bootstrap.LoadConfiguration,
			bootstrap.InitializeRedis,
			bootstrap.HttpServer,
		),
		fx.Invoke(application.Application),
	).Run()
}
