package main

import (
	"context"
	"errors"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/common"
	"time"

	"github.com/cqqqq777/go-kitex-mall/cmd/merchant/dao"
	"github.com/cqqqq777/go-kitex-mall/cmd/merchant/model"
	"github.com/cqqqq777/go-kitex-mall/cmd/merchant/pkg"
	"github.com/cqqqq777/go-kitex-mall/shared/consts"
	"github.com/cqqqq777/go-kitex-mall/shared/errz"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/merchant"
	"github.com/cqqqq777/go-kitex-mall/shared/log"
	"github.com/cqqqq777/go-kitex-mall/shared/middleware"
	"github.com/cqqqq777/go-kitex-mall/shared/response"

	"github.com/bwmarrin/snowflake"
	"github.com/golang-jwt/jwt"
)

// MerchantServiceImpl implements the last service interface defined in the IDL.
type MerchantServiceImpl struct {
	Jwt *middleware.JWT
	Dao *dao.Merchant
}

// Register implements the MerchantServiceImpl interface.
func (s *MerchantServiceImpl) Register(ctx context.Context, req *merchant.MallMerchantRegisterRequest) (resp *merchant.MallMerchantRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(merchant.MallMerchantRegisterResponse)

	// generate id
	sf, err := snowflake.NewNode(consts.MerchantSnowflakeNode)
	if err != nil {
		log.Zlogger.Errorf("generate merchant id failed err:%s", err.Error())
		resp.CommonResp = response.NewCommonResp(errz.ErrGenerateMerchantId)
		return resp, nil
	}
	id := sf.Generate().Int64()

	// create merchant object
	merchant1 := &model.Merchant{
		Id:          id,
		Name:        req.Name,
		Password:    pkg.Md5(req.Password),
		Alipay:      req.Alipay,
		Description: req.Description,
	}

	// create merchant in database
	if err = s.Dao.CreateMerchant(merchant1); err != nil {
		if errors.Is(err, dao.ErrMerchantExist) {
			resp.CommonResp = response.NewCommonResp(errz.ErrMerchantExist)
		} else {
			log.Zlogger.Errorf("create merchant failed err:%s", err.Error())
			resp.CommonResp = response.NewCommonResp(errz.ErrMerchantInternal)
		}
		return resp, nil
	}

	// build response
	resp.Id = id
	resp.Token, err = s.Jwt.CreateToken(middleware.CustomClaims{
		ID:       id,
		Identity: consts.MerchantIdentity,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + consts.TokenExpiredAt,
		},
	})
	if err != nil {
		log.Zlogger.Errorf("generate merchan token failed err:%s", err.Error())
		resp.CommonResp = response.NewCommonResp(errz.ErrGenerateMToken)
		return resp, nil
	}
	resp.CommonResp = response.NewCommonResp(nil)

	return
}

// Login implements the MerchantServiceImpl interface.
func (s *MerchantServiceImpl) Login(ctx context.Context, req *merchant.MallMerchantLoginRequest) (resp *merchant.MallMerchantLoginResponse, err error) {
	// TODO: Your code here...
	resp = new(merchant.MallMerchantLoginResponse)

	// check pwd
	m, err := s.Dao.GetMerchantByName(req.Name)
	req.Password = pkg.Md5(req.Password)
	if m.Password != req.Password {
		resp.CommonResp = response.NewCommonResp(errz.ErrWrongPwd)
		return resp, nil
	}

	// build response
	resp.Id = m.Id
	resp.Token, err = s.Jwt.CreateToken(middleware.CustomClaims{
		ID:       m.Id,
		Identity: consts.MerchantIdentity,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + consts.TokenExpiredAt,
		},
	})
	if err != nil {
		log.Zlogger.Errorf("generate merchan token failed err:%s", err.Error())
		resp.CommonResp = response.NewCommonResp(errz.ErrGenerateMToken)
		return resp, nil
	}
	resp.CommonResp = response.NewCommonResp(nil)

	return
}

// GetInfo implements the MerchantServiceImpl interface.
func (s *MerchantServiceImpl) GetInfo(ctx context.Context, req *merchant.MallMerchantGetInfoRequest) (resp *merchant.MallMerchantGetInfoResponse, err error) {
	// TODO: Your code here...
	resp = new(merchant.MallMerchantGetInfoResponse)

	// get merchant info by id
	m, err := s.Dao.GetMerchantById(req.Id)
	if err != nil {
		log.Zlogger.Errorf("get merchant info failed err:%s", err.Error())
		resp.CommonResp = response.NewCommonResp(errz.ErrGetMerchantInfo)
		return resp, nil
	}

	// build response
	resp.MerchantInfo = new(common.Merchant)
	resp.MerchantInfo.Id = m.Id
	resp.MerchantInfo.Name = m.Name
	resp.MerchantInfo.Alipay = m.Alipay
	resp.MerchantInfo.Description = m.Description
	resp.CommonResp = response.NewCommonResp(nil)

	//cache merchant info
	if err = s.Dao.CacheMerchantInfo(ctx, m); err != nil {
		log.Zlogger.Errorf("cache merchant info failed err:%s", err.Error())
	}

	return resp, nil
}
