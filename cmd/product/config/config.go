package config

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
}

type MongoConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Name     string `mapstructure:"db" json:"db"`
}

type NsqConfig struct {
	Host  string `mapstructure:"host" json:"host"`
	Port  int    `mapstructure:"port" json:"port"`
	Topic string `mapstructure:"topic" json:"topic"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type ServerConfig struct {
	Name            string            `mapstructure:"name" json:"name"`
	Host            string            `mapstructure:"host" json:"host"`
	OtelInfo        OtelConfig        `mapstructure:"otel" json:"otel"`
	MongoInfo       MongoConfig       `mapstructure:"mongo" json:"mongo"`
	NsqInfo         NsqConfig         `mapstructure:"nsq" json:"nsq"`
	RedisInfo       RedisConfig       `mapstructure:"redis" json:"redis"`
	MerchantSrvInfo MerchantSrvConfig `mapstructure:"merchant_srv" json:"merchant_srv"`
	OperateSrvInfo  OperateSrvConfig  `mapstructure:"operate_srv" json:"operate_srv"`
}

type MerchantSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type OperateSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
