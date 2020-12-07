package config

import (
	"emqx-auth-http/config/options"
	"github.com/go-redis/redis/v8"
)

type Config struct {
	Debug  string        `yaml:"debug"`
	Listen string        `yaml:"listen"`
	Redis  redis.Options `yaml:"redis"`
	Key    options.Key   `yaml:"key"`
}
