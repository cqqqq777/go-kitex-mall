package config

import (
	"fmt"
	"github.com/cqqqq777/go-kitex-mall/shared/consts"
)

func (c ServerConfig) GetMysqlDsn() string {
	dsn := fmt.Sprintf(consts.MysqlDns, c.MysqlInfo.User, c.MysqlInfo.Password, c.MysqlInfo.Host, c.MysqlInfo.Port, c.MysqlInfo.Name)
	return dsn
}
