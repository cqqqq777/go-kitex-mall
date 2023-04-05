package main

import (
	"context"
	"net"
	"strconv"

	"github.com/cqqqq777/go-kitex-mall/cmd/user/config"
	"github.com/cqqqq777/go-kitex-mall/cmd/user/dao"
	"github.com/cqqqq777/go-kitex-mall/cmd/user/initialize"
	"github.com/cqqqq777/go-kitex-mall/cmd/user/pkg"
	"github.com/cqqqq777/go-kitex-mall/shared/consts"
	user "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/user/userservice"
	"github.com/cqqqq777/go-kitex-mall/shared/log"
	"github.com/cqqqq777/go-kitex-mall/shared/middleware"

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
	mdb := initialize.InitMongo()
	rdb := initialize.InitRedis()
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	producer, err := pkg.NewPublisher()
	if err != nil {
		log.Zlogger.Fatal("new nsq producer failed")
	}
	userDao := dao.NewUser(db, mdb, rdb)

	impl := &UserServiceImpl{
		Jwt:      middleware.NewJWT(config.GlobalServerConfig.JWTInfo.SigningKey),
		Producer: producer,
		Dao:      userDao,
	}

	// Create new server.
	srv := user.NewServer(impl,
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
