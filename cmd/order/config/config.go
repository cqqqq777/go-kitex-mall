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

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type NsqConfig struct {
	Host          string `mapstructure:"host" json:"host"`
	Port          int    `mapstructure:"port" json:"port"`
	ProducerTopic string `mapstructure:"producer_topic" json:"producer_topic"`
	ConsumerTopic string `mapstructure:"consumer_topic" json:"consumer_topic"`
	Channel       string `mapstructure:"channel" json:"channel"`
}

type ServerConfig struct {
	Name           string        `mapstructure:"name" json:"name"`
	Host           string        `mapstructure:"host" json:"host"`
	OtelInfo       OtelConfig    `mapstructure:"otel" json:"otel"`
	MysqlInfo      MysqlConfig   `mapstructure:"mysql" json:"mysql"`
	RedisInfo      RedisConfig   `mapstructure:"redis" json:"redis"`
	NsqInfo        NsqConfig     `mapstructure:"nsq" json:"nsq"`
	UserSrvInfo    UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	ProductSrvInfo ProductConfig `mapstructure:"product_srv" json:"product_srv"`
}

type UserSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type ProductConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
