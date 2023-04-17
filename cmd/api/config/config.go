package config

import "github.com/cqqqq777/go-kitex-mall/cmd/api/pkg/upload/config"

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
}

type ServerConfig struct {
	Name              string                     `mapstructure:"name" json:"name"`
	Host              string                     `mapstructure:"host" json:"host"`
	Port              int                        `mapstructure:"port" json:"port"`
	JWTInfo           JWTConfig                  `mapstructure:"jwt" json:"jwt"`
	OtelInfo          OtelConfig                 `mapstructure:"otel" json:"otel"`
	RedisInfo         RedisConfig                `mapstructure:"redis" json:"redis"`
	MerchantSrvInfo   RPCSrvConfig               `mapstructure:"merchant_srv" json:"merchant_srv"`
	UserSrvInfo       RPCSrvConfig               `mapstructure:"user_srv" json:"user_srv"`
	ProductSrvInfo    RPCSrvConfig               `mapstructure:"product_srv" json:"product_srv"`
	OperateSrvInfo    RPCSrvConfig               `mapstructure:"operate_srv" json:"operate_srv"`
	OrderSrvInfo      RPCSrvConfig               `mapstructure:"order_srv" json:"order_srv"`
	PaySrvInfo        RPCSrvConfig               `mapstructure:"pay_srv" json:"pay_srv"`
	UploadServiceInfo config.UploadServiceConfig `mapstructure:"upload_srv" json:"upload_srv"`
}

type RPCSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
