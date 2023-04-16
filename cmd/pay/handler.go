package main

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/cqqqq777/go-kitex-mall/cmd/pay/dao"
	"github.com/cqqqq777/go-kitex-mall/cmd/pay/model"
	"github.com/cqqqq777/go-kitex-mall/cmd/pay/pkg"
	"github.com/cqqqq777/go-kitex-mall/shared/consts"
	"github.com/cqqqq777/go-kitex-mall/shared/errz"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/common"
	pay "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/pay"
	"github.com/cqqqq777/go-kitex-mall/shared/log"
	"github.com/cqqqq777/go-kitex-mall/shared/response"
)

// PayServiceImpl implements the last service interface defined in the IDL.
type PayServiceImpl struct {
	Dao *dao.Pay
	UserManager
	OrderManager
	Producer
}

type UserManager interface {
	GetUserInfo(ctx context.Context, uid int64) error
}

type OrderManager interface {
	GetOrderInfo(ctx context.Context, orderID int64) (*common.Order, error)
}

type Producer interface {
	Produce(ctx context.Context, msg pkg.Msg) error
}

// CreatePay implements the PayServiceImpl interface.
func (s *PayServiceImpl) CreatePay(ctx context.Context, req *pay.MallCreatePayRequest) (resp *pay.MallCreatePayResponse, err error) {
	// TODO: Your code here...
	resp = new(pay.MallCreatePayResponse)

	// check user
	if err = s.UserManager.GetUserInfo(ctx, req.UserId); err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrGetUserInfo)
		log.Zlogger.Errorf("get user info failed err:%s", err.Error())
		return resp, nil
	}

	// check order
	orderInfo, err := s.OrderManager.GetOrderInfo(ctx, req.OrderId)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrGetOrder)
		log.Zlogger.Errorf("get order info failed err:%s", err)
		return resp, nil
	}
	if orderInfo.Status == consts.StatusSuccess {
		resp.CommonResp = response.NewCommonResp(errz.ErrOrderPaid)
		return resp, nil
	}
	if orderInfo.Status == consts.StatusCancel {
		resp.CommonResp = response.NewCommonResp(errz.ErrOrderCancel)
		return resp, nil
	}
	if orderInfo.Amount != req.Amount {
		resp.CommonResp = response.NewCommonResp(errz.ErrAmountWrong)
		return resp, nil
	}

	// generate pay id
	sf, err := snowflake.NewNode(consts.PaySnowflakeNode)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrPayInternal)
		log.Zlogger.Errorf("generate pay id failed err:%s", err.Error())
		return resp, nil
	}

	// create pay
	url, err := pkg.Pay(req.OrderId, req.Amount)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrGenerateAlipay)
		log.Zlogger.Errorf("generate alipay failed err:%s", err.Error())
		return resp, nil
	}

	// create in mysql
	payInfo := new(model.Pay)
	payInfo.PayID = sf.Generate().Int64()
	payInfo.OrderID = req.OrderId
	payInfo.UserID = req.UserId
	payInfo.Amount = req.Amount
	payInfo.Status = consts.StatusWaitPay
	payInfo.Url = url
	if err = s.Dao.CreatePay(payInfo); err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrPayInternal)
		log.Zlogger.Errorf("create pay failed err:%s", err.Error())
	}

	resp.CommonResp = response.NewCommonResp(nil)
	resp.PayId = payInfo.PayID
	resp.Url = url
	return resp, nil
}

// PayDetail implements the PayServiceImpl interface.
func (s *PayServiceImpl) PayDetail(ctx context.Context, req *pay.MallPayDetailRequest) (resp *pay.MallPayDetailResponse, err error) {
	// TODO: Your code here...
	resp = new(pay.MallPayDetailResponse)

	// get info
	payInfo, err := s.Dao.GetPay(req.PayId)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrPayInternal)
		log.Zlogger.Errorf("get pay info failed err:%s", err.Error())
		return
	}

	// build response
	resp.Pay.PayId = payInfo.PayID
	resp.Pay.Status = payInfo.Status
	resp.Pay.Url = payInfo.Url
	resp.Pay.OrderId = payInfo.OrderID
	resp.Pay.Amount = payInfo.Amount
	resp.Pay.UserId = payInfo.UserID
	resp.Pay.CreateTime = payInfo.CreateTime.UnixNano()
	resp.Pay.UpdateTime = payInfo.UpdateTime.UnixNano()
	resp.CommonResp = response.NewCommonResp(nil)
	return resp, nil
}

// PayReturn implements the PayServiceImpl interface.
func (s *PayServiceImpl) PayReturn(ctx context.Context, req *pay.MallPayReturnRequest) (resp *pay.MallPayReturnResponse, err error) {
	// TODO: Your code here...
	resp = new(pay.MallPayReturnResponse)

	// check order
	orderInfo, err := s.OrderManager.GetOrderInfo(ctx, req.OrderId)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrGetOrder)
		log.Zlogger.Errorf("get order failed err:%s", err.Error())
		return resp, nil
	}

	// get pay info
	payInfo, err := s.Dao.GetPay(req.PayId)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrPayInternal)
		log.Zlogger.Errorf("get pay info failed err:%s", err.Error())
		return resp, nil
	}

	// check amount
	if payInfo.Amount != orderInfo.Amount {
		resp.CommonResp = response.NewCommonResp(errz.ErrAmountWrong)
		return resp, nil
	}

	// update order and pay info
	if err = s.Producer.Produce(ctx, pkg.Msg{OrderId: orderInfo.OrderId, Status: consts.StatusSuccess}); err != nil {
		log.Zlogger.Errorf("update order status failed err:%s", err.Error())
		resp.CommonResp = response.NewCommonResp(errz.ErrPayInternal)
		return resp, nil
	}
	err = s.Dao.UpdatePayStatus(payInfo.PayID, consts.StatusSuccess)
	if err != nil {
		log.Zlogger.Errorf("update pay stauts failed err:%s", err.Error())
		resp.CommonResp = response.NewCommonResp(errz.ErrPayInternal)
		return resp, nil
	}

	resp.CommonResp = response.NewCommonResp(nil)
	return
}

// PayNotify implements the PayServiceImpl interface.
func (s *PayServiceImpl) PayNotify(ctx context.Context, req *pay.MallPayNotifyRequest) (resp *pay.MallPayNotifyResponse, err error) {
	// TODO: Your code here...
	resp = new(pay.MallPayNotifyResponse)

	// update order status
	if err = s.Producer.Produce(ctx, pkg.Msg{
		OrderId: req.OrderId,
		Status:  req.Status,
	}); err != nil {
		log.Zlogger.Errorf("update order status failed err:%s", err.Error())
		resp.CommonResp = response.NewCommonResp(errz.ErrPayInternal)
		return resp, nil
	}

	// update pay status
	err = s.Dao.UpdatePayStatus(req.PayId, consts.StatusCancel)
	if err != nil {
		log.Zlogger.Errorf("update pay stauts failed err:%s", err.Error())
		resp.CommonResp = response.NewCommonResp(errz.ErrPayInternal)
		return resp, nil
	}

	resp.CommonResp = response.NewCommonResp(nil)
	return
}
