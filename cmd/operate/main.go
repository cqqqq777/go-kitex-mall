package main

import (
	operate "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/operate/opereteservice"
	"log"
)

func main() {
	svr := operate.NewServer(new(OperateServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
