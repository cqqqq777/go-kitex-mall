package main

import (
	"context"

	"github.com/cqqqq777/go-kitex-mall/cmd/operate/dao"
	"github.com/cqqqq777/go-kitex-mall/shared/errz"
	"github.com/cqqqq777/go-kitex-mall/shared/log"
	"github.com/cqqqq777/go-kitex-mall/shared/response"

	operate "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/operate"
)

// OperateServiceImpl implements the last service interface defined in the IDL.
type OperateServiceImpl struct {
	Dao *dao.Operate
}

// FavoriteProduct implements the OperateServiceImpl interface.
func (s *OperateServiceImpl) FavoriteProduct(ctx context.Context, req *operate.MallFavoriteProductRequest) (resp *operate.MallFavoriteProductResponse, err error) {
	// TODO: Your code here...
	resp = new(operate.MallFavoriteProductResponse)

	// get status
	status, err := s.Dao.GetUserFavoriteStatus(ctx, req.UserId, req.ProductId)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrOperateInternal)
		log.Zlogger.Errorf("get user favorite status failed err:%s", err.Error())
		return resp, nil
	}

	// favorite or cancel favorite
	if status {
		err = s.Dao.FavoriteProduct(ctx, req.UserId, req.ProductId)
		if err != nil {
			resp.CommonResp = response.NewCommonResp(errz.ErrFavoriteProduct)
			log.Zlogger.Errorf("favorite product failed err:%s", err.Error())
			return resp, nil
		}
	} else {
		err = s.Dao.CancelFavorite(ctx, req.UserId, req.ProductId)
		if err != nil {
			resp.CommonResp = response.NewCommonResp(errz.ErrFavoriteProduct)
			log.Zlogger.Errorf("cancel favorite product failed err:%s", err.Error())
			return resp, nil
		}
	}

	resp.CommonResp = response.NewCommonResp(nil)
	return resp, nil
}

// GetProductOperateInfo implements the OperateServiceImpl interface.
func (s *OperateServiceImpl) GetProductOperateInfo(ctx context.Context, req *operate.MallGetProductOperateInfoRequest) (resp *operate.MallGetProductOperateInfoResponse, err error) {
	// TODO: Your code here...
	resp = new(operate.MallGetProductOperateInfoResponse)

	// get favorite status
	status, err := s.Dao.GetUserFavoriteStatus(ctx, req.UserId, req.ProductId)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrGetFavoriteStatus)
		log.Zlogger.Errorf("get user favorite status failed err:%s", err.Error())
		return resp, nil
	}

	// get comment num
	comNum, err := s.Dao.GetCommentNum(req.ProductId)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrGetCommentNum)
		log.Zlogger.Errorf("get comment num faile err:%s", err.Error())
		return resp, nil
	}

	// get sale num
	saleNum, err := s.Dao.GetSaleNum(ctx, req.ProductId)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrGetSaleNum)
		log.Zlogger.Errorf("get sale num failed err:%s", err.Error())
		return resp, nil
	}

	// build response
	resp.CommonResp = response.NewCommonResp(nil)
	resp.OperateInfo.CommentCount = comNum
	resp.OperateInfo.SaleCount = saleNum
	resp.OperateInfo.IsFavorite = status

	return resp, nil
}
