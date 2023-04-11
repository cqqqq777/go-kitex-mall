package main

import (
	"context"
	"net"
	"strconv"

	"github.com/cqqqq777/go-kitex-mall/cmd/order/config"
	"github.com/cqqqq777/go-kitex-mall/cmd/order/dao"
	"github.com/cqqqq777/go-kitex-mall/cmd/order/initialize"
	"github.com/cqqqq777/go-kitex-mall/cmd/order/pkg"
	"github.com/cqqqq777/go-kitex-mall/shared/consts"
	order "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/order/orderservice"
	"github.com/cqqqq777/go-kitex-mall/shared/log"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func main() {
	// init
	IP, Port := initialize.InitFlag()
	r, info := initialize.InitNacos(Port)
	db := initialize.InitMysql()
	rdb := initialize.InitRedis()
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	orderDao := dao.NewOrder(db, rdb)
	userManager := initialize.InitUser()
	productManager := initialize.InitProduct()

	impl := &OrderServiceImpl{
		Dao:            orderDao,
		UserManager:    &pkg.UserManager{UserService: userManager},
		ProductManager: &pkg.ProductManage{ProductService: productManager},
	}

	srv := order.NewServer(impl,
		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
	)

	err := srv.Run()

	if err != nil {
		log.Zlogger.Fatal(err)
	}
}
