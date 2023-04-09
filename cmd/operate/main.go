package main

import (
	"context"
	"net"
	"strconv"

	"github.com/cqqqq777/go-kitex-mall/cmd/operate/config"
	"github.com/cqqqq777/go-kitex-mall/cmd/operate/dao"
	"github.com/cqqqq777/go-kitex-mall/cmd/operate/initialize"
	"github.com/cqqqq777/go-kitex-mall/shared/consts"
	operate "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/operate/operateservice"
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

	operateDao := dao.NewOperate(db, rdb)
	impl := &OperateServiceImpl{
		operateDao,
	}

	// Create new server.
	svr := operate.NewServer(impl,
		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
	)

	err := svr.Run()

	if err != nil {
		log.Zlogger.Fatal(err)
	}
}
