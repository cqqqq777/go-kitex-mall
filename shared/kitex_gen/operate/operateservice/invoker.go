// Code generated by Kitex v0.5.1. DO NOT EDIT.

package operateservice

import (
	server "github.com/cloudwego/kitex/server"
	operate "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/operate"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler operate.OperateService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
