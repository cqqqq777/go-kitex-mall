package config

import (
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/merchant/merchantservice"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/operate/operateservice"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/order/orderservice"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/pay/payservice"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/product/productservice"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/user/userservice"
)

var (
	GlobalServerConfig = &ServerConfig{}
	GlobalNacosConfig  = &NacosConfig{}

	GlobalUserClient     userservice.Client
	GlobalMerchantClient merchantservice.Client
	GlobalProductClient  productservice.Client
	GlobalOperateClient  operateservice.Client
	GlobalOrderClient    orderservice.Client
	GlobalPayClient      payservice.Client
)
