package pkg

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cqqqq777/go-kitex-mall/cmd/order/config"
	"github.com/cqqqq777/go-kitex-mall/cmd/order/dao"
	"github.com/nsqio/go-nsq"
)

type Producer struct {
	Producer *nsq.Producer
}

type ProducerMsg struct {
	OrderID int64 `json:"order_id"`
	Amount  int64 `json:"amount"`
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

func (p *Producer) Produce(msg ProducerMsg) error {
	body, err := sonic.Marshal(msg)
	if err != nil {
		return fmt.Errorf("cannot marshal: %v", err)
	}
	return p.Producer.Publish(config.GlobalServerConfig.NsqInfo.ProducerTopic, body)
}

type Consumer struct {
	Consumer *nsq.Consumer
}

type ConsumerMsg struct {
	OrderID int64 `json:"order_id"`
	Status  int8  `json:"status"`
}

func NewConsumer() (consumer *Consumer, err error) {
	consumer = new(Consumer)
	conf := nsq.NewConfig()
	consumer.Consumer, err = nsq.NewConsumer(config.GlobalServerConfig.NsqInfo.ConsumerTopic, config.GlobalServerConfig.NsqInfo.Channel, conf)
	return
}

func handleMsg(dao *dao.Order) func(*nsq.Message) error {
	return func(msg *nsq.Message) error {
		consumerMsg := new(ConsumerMsg)
		err := sonic.Unmarshal(msg.Body, consumerMsg)
		if err != nil {
			return err
		}
		err = dao.UpdateOrder(consumerMsg.OrderID, consumerMsg.Status)
		if err != nil {
			return err
		}
		return nil
	}
}

func (c *Consumer) Consume(dao *dao.Order) error {
	host := fmt.Sprintf("%s:%d", config.GlobalServerConfig.NsqInfo.Host, config.GlobalServerConfig.NsqInfo.Port)
	c.Consumer.AddHandler(nsq.HandlerFunc(handleMsg(dao)))
	err := c.Consumer.ConnectToNSQD(host)
	if err != nil {
		return err
	}
	select {}
}
