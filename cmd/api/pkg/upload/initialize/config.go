package initialize

import (
	"github.com/cqqqq777/go-kitex-mall/cmd/api/config"
	config1 "github.com/cqqqq777/go-kitex-mall/cmd/api/pkg/upload/config"
)

func initConfig() {
	config1.GlobalServiceConfig = &config.GlobalServerConfig.UploadServiceInfo
	config1.GlobalChannel = make(chan []byte, 1)
}
