package config

import (
	"fmt"

	"github.com/cqqqq777/go-kitex-mall/shared/consts"

	"github.com/redis/go-redis/v9"
)

func (c ServerConfig) GetMysqlDsn() string {
	dsn := fmt.Sprintf(consts.MysqlDns, c.MysqlInfo.User, c.MysqlInfo.Password, c.MysqlInfo.Host, c.MysqlInfo.Port, c.MysqlInfo.Name)
	return dsn
}

func (c ServerConfig) GetRedisOption() *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.RedisInfo.Host, c.RedisInfo.Port),
		Password: c.RedisInfo.Password,
	}
}
