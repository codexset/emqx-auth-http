package bootstrap

import (
	"context"
	"emqx-auth-http/config"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	LoadConfigurationNotExists = errors.New("the configuration file does not exist")
)

// Load application configuration
// reference config.example.yml
func LoadConfiguration() (cfg *config.Config, err error) {
	if _, err = os.Stat("./config/config.yml"); os.IsNotExist(err) {
		err = LoadConfigurationNotExists
		return
	}
	var bs []byte
	bs, err = ioutil.ReadFile("./config/config.yml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(bs, &cfg)
	if err != nil {
		return
	}
	return
}

//Initialize redis configuration
// reference https://github.com/go-redis/redis
func InitializeRedis(cfg *config.Config) *redis.Client {
	return redis.NewClient(&cfg.Redis)
}

// Start http service
// https://gin-gonic.com/docs/examples/custom-http-config/
func HttpServer(lc fx.Lifecycle, cfg *config.Config) (serve *gin.Engine) {
	if cfg.Debug != "" {
		go http.ListenAndServe(cfg.Debug, nil)
	}
	serve = gin.New()
	serve.Use(gin.Recovery())
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go serve.Run(cfg.Listen)
			return nil
		},
	})
	return
}
