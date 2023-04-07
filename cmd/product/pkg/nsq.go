package pkg

import (
	"fmt"

	"github.com/cqqqq777/go-kitex-mall/cmd/product/config"

	"github.com/nsqio/go-nsq"
)

type Producer struct {
	Producer *nsq.Producer
}

func NewPublisher() (pro *Producer, err error) {
	pro = new(Producer)
	conf := nsq.NewConfig()
	host := fmt.Sprintf("%s:%d", config.GlobalServerConfig.NsqInfo.Host, config.GlobalServerConfig.NsqInfo.Port)
	pro.Producer, err = nsq.NewProducer(host, conf)
	if err != nil {
		pro.Producer.Stop()
		return
	}
	return
}
