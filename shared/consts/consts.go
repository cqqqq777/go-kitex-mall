package consts

import "time"

// database constant
const (
	MysqlDns = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	MongoUrl = "mongodb://%s:%d"
)

// config file constant
const (
	UserConfigFile     = "./cmd/user/config.yaml"
	MerchantConfigFile = "./cmd/merchant/config.yaml"
	CartConfigFile     = "./cmd/cart/config.yaml"
	ChatConfigFile     = "./cmd/chat/config.yaml"
	CommentConfigFile  = "./comment/user/config.yaml"
	PayConfigFile      = "./cmd/pay/config.yaml"
	ProductConfigFile  = "./product/user/config.yaml"
)

// nacos constant
const (
	NacosLogDir   = "tmp/nacos/log"
	NacosCacheDir = "tmp/nacos/cache"
	NacosLogLevel = "debug"
)

// net constant
const (
	TCP             = "tcp"
	FreePortAddress = "localhost:0"
)

// snowflake node constant
const (
	NacosSnowflakeNode = iota
	UserSnowflakeNode
	MerchantSnowflakeNode
	CartSnowflakeNode
	CommentSnowflakeNode
	ProductSnowflakeNode
	PaySnowflakeNode
	ChatSnowflakeNode
	MinioSnowflakeNode
)

// identity constant
const (
	UserIdentity     = "user"
	MerchantIdentity = "merchant"
)

// expire time constant
const (
	VerificationExpTime = time.Second * 600

	CacheExpTime = time.Second * 300

	TokenExpiredAt = 43200
)

// mongodb collections constant
const (
	CollectionUsers     = "users"
	CollectionMerchants = "merchants"
	CollectionProducts  = "products"
)

// ip and port constant
const (
	IPFlagName  = "ip"
	IPFlagValue = "0.0.0.0"
	IPFlagUsage = "address"

	PortFlagName  = "port"
	PortFlagUsage = "port"
)

// object type constant
const (
	ObjectAvatarType     = "avatar"
	ObjectBackgroundType = "background"
	ObjectProductType    = "product"
)

// app context key constant
const (
	AccountID = "accountID"

	AccountIdentity = "accountIdentity"
)
