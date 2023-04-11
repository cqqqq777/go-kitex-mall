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
	CommentConfigFile  = "./cmd/comment/config.yaml"
	PayConfigFile      = "./cmd/pay/config.yaml"
	ProductConfigFile  = "./cmd/product/config.yaml"
	OperateConfigFile  = "./cmd/operate/config.yaml"
	OrderConfigFile    = "./cmd/order/config.yaml"
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
	OrderSnowflakeNode
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

	CacheExpTime = time.Second * 1800

	TokenExpiredAt = 43200
	OrderExpTime   = time.Second * 900
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

const (
	ApiGroup = "API_GROUP"

	UserGroup     = "USER_GROUP"
	MerchantGroup = "MERCHANT_GROUP"
	OperateGroup  = "OPERATE_GROUP"
	ProductGroup  = "PRODUCT_GROUP"
)

const (
	StatusSuccess int8 = iota + 1
	StatusWaitPay
	StatusCancel
)
