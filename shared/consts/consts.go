package consts

import "time"

const (
	MysqlDns = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	MongoUrl = "mongodb://%s:%d"
)

const (
	UserConfigFile    = "./cmd/user/config.yaml"
	CartConfigFile    = "./cmd/cart/config.yaml"
	ChatConfigFile    = "./cmd/chat/config.yaml"
	CommentConfigFile = "./comment/user/config.yaml"
	PayConfigFile     = "./cmd/pay/config.yaml"
	ProductConfigFile = "./product/user/config.yaml"
)

const (
	NacosLogDir   = "tmp/nacos/log"
	NacosCacheDir = "tmp/nacos/cache"
	NacosLogLevel = "debug"
)

const (
	TCP             = "tcp"
	FreePortAddress = "localhost:0"
)

const (
	NacosSnowflakeNode = iota
	UserSnowflakeNode
	CartSnowflakeNode
	CommentSnowflakeNode
	ProductSnowflakeNode
	PaySnowflakeNode
	ChatSnowflakeNode
	MinioSnowflakeNode
)

const (
	VerificationExpTime = time.Second * 600
)

const (
	CacheExpTime = time.Second * 300
)

const (
	CollectionUsers = "users"
)

const (
	TokenExpiredAt = 43200
)

const (
	IPFlagName  = "ip"
	IPFlagValue = "0.0.0.0"
	IPFlagUsage = "address"

	PortFlagName  = "port"
	PortFlagUsage = "port"
)
