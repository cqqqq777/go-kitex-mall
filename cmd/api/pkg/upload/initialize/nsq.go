package initialize

import (
	"github.com/cqqqq777/go-kitex-mall/cmd/api/pkg/upload/config"
	"github.com/cqqqq777/go-kitex-mall/shared/log"
	"github.com/nsqio/go-nsq"
)

func newConsumer() *nsq.Consumer {
	conf := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(config.GlobalServiceConfig.NsqInfo.Topic, config.GlobalServiceConfig.NsqInfo.Channel, conf)
	if err != nil {
		log.Zlogger.Fatal(err)
	}
	consumer.SetLogger(nil, 0)
	return consumer
}
