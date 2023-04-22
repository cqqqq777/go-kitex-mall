package config

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

func (c ServerConfig) GetRedisOption() *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.RedisInfo.Host, c.RedisInfo.Port),
		Password: c.RedisInfo.Password,
	}
}
