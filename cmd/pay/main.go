package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/cqqqq777/go-kitex-mall/cmd/pay/config"
	"github.com/cqqqq777/go-kitex-mall/cmd/pay/dao"
	"github.com/cqqqq777/go-kitex-mall/cmd/pay/initialize"
	"github.com/cqqqq777/go-kitex-mall/cmd/pay/pkg"
	"github.com/cqqqq777/go-kitex-mall/shared/consts"
	pay "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/pay/payservice"
	"github.com/cqqqq777/go-kitex-mall/shared/log"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"net"
	"strconv"
)

func main() {
	// init
	IP, Port := initialize.InitFlag()
	r, info := initialize.InitNacos(Port)
	db := initialize.InitMysql()
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	payDao := dao.NewPayDao(db)
	userManager := initialize.InitUser()
	orderManager := initialize.InitOrder()
	producer, err := pkg.NewPublisher()
	if err != nil {
		log.Zlogger.Fatalf("new producer failed err:%s", err.Error())
	}

	impl := &PayServiceImpl{
		Dao:          payDao,
		UserManager:  &pkg.UserManager{UserService: userManager},
		OrderManager: &pkg.OrderManager{OrderService: orderManager},
		Producer:     producer,
	}

	srv := pay.NewServer(impl,
		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
	)

	err = srv.Run()

	if err != nil {
		log.Zlogger.Fatal(err)
	}
}
