package initialize

import (
	"fmt"
	"github.com/cqqqq777/go-kitex-mall/cmd/pay/config"
	"github.com/cqqqq777/go-kitex-mall/shared/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() *gorm.DB {
	dsn := config.GlobalServerConfig.GetMysqlDsn()
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		msg := fmt.Sprintf("init database failed err:%s", err.Error())
		log.Zlogger.Fatal(msg)
	}
	return db
}
