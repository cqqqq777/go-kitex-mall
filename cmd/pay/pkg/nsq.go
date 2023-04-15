package pkg

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cqqqq777/go-kitex-mall/cmd/pay/config"
	"github.com/nsqio/go-nsq"
)

type Msg struct {
	OrderId int64 `json:"order_id"`
	Status  int8  `json:"status"`
}

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

func (p *Producer) Produce(ctx context.Context, msg Msg) error {
	body, err := sonic.Marshal(msg)
	if err != nil {
		return fmt.Errorf("cannot marshal: %v", err)
	}
	return p.Producer.Publish(config.GlobalServerConfig.NsqInfo.Topic, body)
}
