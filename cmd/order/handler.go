package main

import (
	"context"
	"errors"
	"time"

	"github.com/cqqqq777/go-kitex-mall/cmd/order/dao"
	"github.com/cqqqq777/go-kitex-mall/cmd/order/model"
	"github.com/cqqqq777/go-kitex-mall/cmd/order/pkg"
	"github.com/cqqqq777/go-kitex-mall/shared/consts"
	"github.com/cqqqq777/go-kitex-mall/shared/errz"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/common"
	order "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/order"
	"github.com/cqqqq777/go-kitex-mall/shared/log"
	"github.com/cqqqq777/go-kitex-mall/shared/response"

	"github.com/bwmarrin/snowflake"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct {
	Dao *dao.Order
	UserManager
	ProductManager
	Producer
}

type UserManager interface {
	GetUserInfo(ctx context.Context, uid int64) error
}

type ProductManager interface {
	GetProductInfo(ctx context.Context, pid int64) (*common.Product, error)
	UpdateStock(ctx context.Context, pid, stock int64) error
}

type Producer interface {
	Produce(msg pkg.ProducerMsg) error
}

// CreateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order.MallCreateOrderRequest) (resp *order.MallCreateOrderResponse, err error) {
	// TODO: Your code here...
	resp = new(order.MallCreateOrderResponse)

	// check user
	err = s.UserManager.GetUserInfo(ctx, req.UserId)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrGetUserInfo)
		return resp, nil
	}

	// check product
	productInfo, err := s.ProductManager.GetProductInfo(ctx, req.ProductId)
	if err != nil {
		if errors.Is(err, pkg.ErrNoSuchProduct) {
			resp.CommonResp = response.NewCommonResp(errz.ErrNoProduct)
			return resp, nil
		}
		resp.CommonResp = response.NewCommonResp(errz.ErrProductInternal)
		log.Zlogger.Errorf("get product info failed err:%s", err.Error())
		return resp, nil
	}

	if productInfo.Stock < req.ProductNum {
		resp.CommonResp = response.NewCommonResp(errz.ErrShortage)
		return resp, nil
	}

	// generate id
	sf, err := snowflake.NewNode(consts.OrderSnowflakeNode)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrGenerateOrderId)
		log.Zlogger.Errorf("generate order id failed err:%s", err.Error())
		return resp, nil
	}

	// create order in mysql
	orderInfo := &model.Order{
		Id:         sf.Generate().Int64(),
		UserId:     req.UserId,
		ProductId:  req.ProductId,
		ProductNum: req.ProductNum,
		Status:     consts.StatusWaitPay,
		Amount:     req.Amount,
		ExpTime:    time.Now().UnixNano() + consts.OrderExpTime.Nanoseconds(),
	}
	if err = s.Dao.CreateOrder(orderInfo); err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrOrderInternal)
		log.Zlogger.Errorf("create order failed err:%s", err)
		return resp, nil
	}

	// update stock
	err = s.ProductManager.UpdateStock(ctx, productInfo.Id, productInfo.Stock-1)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrProductInternal)
		log.Zlogger.Errorf("update product stock failed err:%s", err.Error())
		return resp, nil
	}

	// set order in redis
	//err = s.Dao.SetOrderInRedis(ctx, orderInfo.Id, orderInfo.Amount, orderInfo.ExpTime)
	//if err != nil {
	//	log.Zlogger.Errorf("set order in redis failed err:%s", err.Error())
	//}

	// publish msg to nsq
	err = s.Producer.Produce(pkg.ProducerMsg{
		OrderID: orderInfo.Id,
		Amount:  orderInfo.Amount,
	})
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrOrderInternal)
		log.Zlogger.Errorf("publish order to nsq failed err:%s", err.Error())
		return resp, nil
	}

	resp.CommonResp = response.NewCommonResp(nil)
	resp.OrderId = orderInfo.Id
	return resp, nil
}

// UpdateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrder(ctx context.Context, req *order.MallUpdateOrderRequest) (resp *order.MallUpdateOrderResponse, err error) {
	// TODO: Your code here...
	resp = new(order.MallUpdateOrderResponse)

	// update status
	err = s.Dao.UpdateOrder(req.OrderId, req.Status)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrUpdateOrder)
		log.Zlogger.Errorf("update order status failed err:%s", err.Error())
		return resp, nil
	}

	resp.CommonResp = response.NewCommonResp(nil)
	return resp, nil
}

// OrderList implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) OrderList(ctx context.Context, req *order.MallOrderListRequset) (resp *order.MallOrderListResponse, err error) {
	// TODO: Your code here...
	resp = new(order.MallOrderListResponse)

	// get list
	list, err := s.Dao.GetOrderList(req.UserId)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrGetOrder)
		log.Zlogger.Errorf("get order list failed err:%s", err.Error())
		return resp, nil
	}

	// build response
	for _, v := range list {
		var singleOrder *common.Order
		singleOrder.OrderId = v.Id
		singleOrder.UserId = v.UserId
		singleOrder.ProductId = v.UserId
		singleOrder.Status = v.Status
		singleOrder.Amount = v.Amount
		singleOrder.ExpTime = v.ExpTime
		singleOrder.CreateTime = v.CreateTime.UnixNano()
		singleOrder.UpdateTime = v.UpdateTime.UnixNano()
		singleOrder.ProductNum = v.ProductNum
		resp.Orders = append(resp.Orders, singleOrder)
	}
	resp.CommonResp = response.NewCommonResp(nil)

	return resp, nil
}

// GetOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrder(ctx context.Context, req *order.MallGetOrderRequest) (resp *order.MallGetOrderResponse, err error) {
	// TODO: Your code here...
	resp = new(order.MallGetOrderResponse)

	// get order
	v, err := s.Dao.GetOrder(req.OrderId)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrGetOrder)
		log.Zlogger.Errorf("get order info failed err:%s", err.Error())
		return resp, nil
	}

	// build response
	var singleOrder *common.Order
	singleOrder.OrderId = v.Id
	singleOrder.UserId = v.UserId
	singleOrder.ProductId = v.UserId
	singleOrder.Status = v.Status
	singleOrder.Amount = v.Amount
	singleOrder.ExpTime = v.ExpTime
	singleOrder.CreateTime = v.CreateTime.UnixNano()
	singleOrder.UpdateTime = v.UpdateTime.UnixNano()
	singleOrder.ProductNum = v.ProductNum
	resp.Order = singleOrder
	resp.CommonResp = response.NewCommonResp(nil)

	return resp, nil
}
