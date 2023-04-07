package config

import (
	"fmt"

	"github.com/cqqqq777/go-kitex-mall/shared/consts"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c ServerConfig) GetMongoOptions() (*options.ClientOptions, string) {
	url := fmt.Sprintf(consts.MongoUrl, c.MongoInfo.Host, c.MongoInfo.Port)
	clientOptions := options.Client().ApplyURI(url)
	credential := c.GetMongoAuth()
	clientOptions.SetAuth(credential)
	return clientOptions, c.MongoInfo.Name
}

func (c ServerConfig) GetMongoAuth() options.Credential {
	return options.Credential{
		Username: c.MongoInfo.User,
		Password: c.MongoInfo.Password,
	}
}

func (c ServerConfig) GetRedisOption() *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.RedisInfo.Host, c.RedisInfo.Port),
		Password: c.RedisInfo.Password,
	}
}
