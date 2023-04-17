package pkg

import (
	"github.com/cqqqq777/go-kitex-mall/cmd/api/pkg/upload/config"
	"github.com/nsqio/go-nsq"
)

func HandleMsg(msg *nsq.Message) error {
	config.GlobalChannel <- msg.Body
	return nil
}
