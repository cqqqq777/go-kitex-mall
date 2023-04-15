package pkg

import (
	"context"
	"fmt"
	"github.com/cqqqq777/go-kitex-mall/shared/errz"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/common"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/order"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/order/orderservice"
)

type OrderManager struct {
	OrderService orderservice.Client
}

func (o OrderManager) GetOrderInfo(ctx context.Context, orderId int64) (*common.Order, error) {
	resp, err := o.OrderService.GetOrder(ctx, &order.MallGetOrderRequest{OrderId: orderId})
	if err != nil {
		return nil, err
	}
	if resp.CommonResp.Code != errz.Success {
		return nil, fmt.Errorf("%s", resp.CommonResp.Msg)
	}
	return resp.Order, nil
}
