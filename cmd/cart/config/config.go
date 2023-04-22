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

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type ServerConfig struct {
	Name           string        `mapstructure:"name" json:"name"`
	Host           string        `mapstructure:"host" json:"host"`
	OtelInfo       OtelConfig    `mapstructure:"otel" json:"otel"`
	RedisInfo      RedisConfig   `mapstructure:"redis" json:"redis"`
	UserSrvInfo    UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	ProductSrvInfo ProductConfig `mapstructure:"product_srv" json:"product_srv"`
}

type UserSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type ProductConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
