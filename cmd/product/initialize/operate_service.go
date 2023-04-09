package initialize

import (
	"github.com/cqqqq777/go-kitex-mall/cmd/product/config"
	"github.com/cqqqq777/go-kitex-mall/shared/consts"
	operate "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/operate/operateservice"
	"github.com/cqqqq777/go-kitex-mall/shared/log"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	nacos "github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func InitOperate() operate.Client {
	// init resolver
	// Read configuration information from nacos
	sc := []constant.ServerConfig{
		{
			IpAddr: config.GlobalNacosConfig.Host,
			Port:   config.GlobalNacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         config.GlobalNacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              consts.NacosLogDir,
		CacheDir:            consts.NacosCacheDir,
		LogLevel:            consts.NacosLogLevel,
	}

	nacosCli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		})
	r := nacos.NewNacosResolver(nacosCli, nacos.WithGroup(consts.OperateGroup))
	if err != nil {
		log.Zlogger.Fatalf("new nacos client failed: %s", err.Error())
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)

	// create a new client
	c, err := operate.NewClient(
		config.GlobalServerConfig.MerchantSrvInfo.Name,
		client.WithResolver(r),                                     // service discovery
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()), // load balance
		client.WithMuxConnection(1),                                // multiplexing
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.OperateSrvInfo.Name}),
	)
	if err != nil {
		log.Zlogger.Fatalf("cannot init operate client err:%s", err.Error())
	}
	return c
}
