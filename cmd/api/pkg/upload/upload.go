package upload

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cqqqq777/go-kitex-mall/cmd/api/pkg/upload/config"
	"github.com/cqqqq777/go-kitex-mall/cmd/api/pkg/upload/pkg"
	"github.com/cqqqq777/go-kitex-mall/shared/log"
	"github.com/minio/minio-go/v7"
	"github.com/nsqio/go-nsq"
	"os"
)

type Server struct {
	config      *config.UploadServiceConfig
	minioClient *minio.Client
	consumer    *nsq.Consumer
	ch          chan []byte
}

type Msg struct {
	Id      int64  `json:"id"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func NewUploadServer(config *config.UploadServiceConfig, minioClient *minio.Client, consumer *nsq.Consumer, ch chan []byte) *Server {
	return &Server{
		config:      config,
		minioClient: minioClient,
		consumer:    consumer,
		ch:          ch,
	}
}

func (s *Server) UploadFile() {
	host := fmt.Sprintf("%s:%d", config.GlobalServiceConfig.NsqInfo.Host, config.GlobalServiceConfig.NsqInfo.Port)
	err := s.consumer.ConnectToNSQD(host)
	if err != nil {
		log.Zlogger.Fatal("connect to nsqd failed err:", err)
	}
	s.consumer.AddHandler(nsq.HandlerFunc(pkg.HandleMsg))
	for {
		select {
		case body := <-s.ch:
			err = s.upload(body)
			if err != nil {
				break
			}
		}
	}
}

func (s *Server) upload(body []byte) error {
	msg := new(Msg)
	err := sonic.Unmarshal(body, msg)
	if err != nil {
		log.Zlogger.Error("unmarshal nsq msg failed err:%s", err.Error())
		return err
	}
	defer func() {
		_ = os.Remove(msg.Message)
	}()
	if _, err = s.minioClient.FPutObject(context.Background(), s.config.MinioInfo.Bucket, msg.Message, msg.Message, minio.PutObjectOptions{
		ContentType: "image/png",
	}); err != nil {
		log.Zlogger.Error("upload file failed err:%s", err.Error())
		return err
	}
	return nil
}
