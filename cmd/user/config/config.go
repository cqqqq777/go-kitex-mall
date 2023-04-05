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

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
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

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type EmailConfig struct {
	Host     string `mapstructure:"host"`
	Email    string `mapstructure:"email"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
}

type ServerConfig struct {
	Name      string      `mapstructure:"name" json:"name"`
	Host      string      `mapstructure:"host" json:"host"`
	JWTInfo   JWTConfig   `mapstructure:"jwt" json:"jwt"`
	OtelInfo  OtelConfig  `mapstructure:"otel" json:"otel"`
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
	MongoInfo MongoConfig `mapstructure:"mongo" json:"mongo"`
	NsqInfo   NsqConfig   `mapstructure:"nsq" json:"nsq"`
	RedisInfo RedisConfig `mapstructure:"redis" json:"redis"`
	EmailInfo EmailConfig `mapstructure:"email" json:"email"`
}

type RPCSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
