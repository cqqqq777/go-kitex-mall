package initialize

import (
	"context"
	"fmt"

	"github.com/cqqqq777/go-kitex-mall/cmd/product/config"
	"github.com/cqqqq777/go-kitex-mall/shared/log"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

// InitMongo init mongodb
func InitMongo() *mongo.Database {
	options, name := config.GlobalServerConfig.GetMongoOptions()
	db, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		msg := fmt.Sprintf("connect mongo failed err:%s", err.Error())
		log.Zlogger.Fatal(msg)
	}
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		msg := fmt.Sprintf("ping mongo failed err:%s", err.Error())
		log.Zlogger.Fatal(msg)
	}
	return db.Database(name)
}

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
