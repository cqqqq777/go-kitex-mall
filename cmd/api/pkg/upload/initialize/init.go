package initialize

import (
	"github.com/cqqqq777/go-kitex-mall/cmd/api/pkg/upload"
	"github.com/cqqqq777/go-kitex-mall/cmd/api/pkg/upload/config"
)

func Init() *upload.Server {
	consumer := newConsumer()
	minioClient := newMinio()
	config.GlobalChannel = make(chan []byte, 1)
	return upload.NewUploadServer(config.GlobalServiceConfig, minioClient, consumer, config.GlobalChannel)
}
