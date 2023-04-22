package initialize

import (
	"context"
	"fmt"
	"github.com/cqqqq777/go-kitex-mall/cmd/cart/config"
	"github.com/cqqqq777/go-kitex-mall/shared/log"
	"github.com/redis/go-redis/v9"
)

// InitRedis init redis client
func InitRedis() *redis.Client {
	option := config.GlobalServerConfig.GetRedisOption()
	rdb := redis.NewClient(option)
	_, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		msg := fmt.Sprintf("init redis client failed err:%v", err)
		log.Zlogger.Fatal(msg)
	}
	return rdb
}
