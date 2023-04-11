package initialize

import (
	"context"
	"fmt"

	"github.com/cqqqq777/go-kitex-mall/cmd/order/config"
	"github.com/cqqqq777/go-kitex-mall/shared/log"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitMysql init mysql
func InitMysql() *gorm.DB {
	dsn := config.GlobalServerConfig.GetMysqlDsn()
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		msg := fmt.Sprintf("init database failed err:%s", err.Error())
		log.Zlogger.Fatal(msg)
	}
	return db
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
