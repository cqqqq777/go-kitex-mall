package main

import (
	"context"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/cqqqq777/go-kitex-mall/cmd/cart/dao"
	"github.com/cqqqq777/go-kitex-mall/cmd/cart/model"
	"github.com/cqqqq777/go-kitex-mall/cmd/cart/pkg"
	"github.com/cqqqq777/go-kitex-mall/shared/errz"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/cart"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/common"
	"github.com/cqqqq777/go-kitex-mall/shared/log"
	"github.com/cqqqq777/go-kitex-mall/shared/response"
	"time"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct {
	Dao *dao.Cart
	ProductManager
	UserManager
}

type ProductManager interface {
	GetProductInfo(ctx context.Context, productId int64) (productInfo *pkg.Product, err error)
}

type UserManager interface {
	GetUserInfo(ctx context.Context, userId int64) error
}

// AddProductToCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddProductToCart(ctx context.Context, req *cart.MallAddProductToCartRequest) (resp *cart.MallAddProductToCartResponse, err error) {
	// TODO: Your code here...
	resp = new(cart.MallAddProductToCartResponse)

	// get user info
	if err != s.UserManager.GetUserInfo(ctx, req.UserId) {
		if errors.Is(err, pkg.ErrNoSuchUser) {
			resp.CommonResp = response.NewCommonResp(errz.NewErrZ(errz.WithErr(err), errz.WithCode(errz.CodeGetUserInfo)))
			return resp, nil
		}
		resp.CommonResp = response.NewCommonResp(errz.ErrGetUserInfo)
		log.Zlogger.Errorf("get user info failed err:%s", err.Error())
		return resp, nil
	}

	// get product info
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

	// check stock
	if req.ProductNum > productInfo.Stock {
		resp.CommonResp = response.NewCommonResp(errz.ErrShortage)
		return resp, nil
	}

	// add product to cart
	body, err := sonic.Marshal(&model.CartProduct{
		ProductId:  req.ProductId,
		ProductNum: req.ProductNum,
		MerchantId: productInfo.MerchantId,
		Amount:     productInfo.Price * req.ProductNum,
		AddTime:    time.Now().UnixNano(),
	})
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrCartInternal)
		log.Zlogger.Errorf("marshal product failed err:%s", err.Error())
		return resp, nil
	}
	err = s.Dao.AddProductToCart(ctx, req.UserId, req.ProductId, body)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrAddProductToCart)
		log.Zlogger.Errorf("%s err:%s", errz.ErrAddProductToCart.Msg, err.Error())
		return resp, nil
	}

	resp.CommonResp = response.NewCommonResp(nil)
	return resp, nil
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.MallGetCartRequest) (resp *cart.MallGetCartResponse, err error) {
	// TODO: Your code here...
	resp = new(cart.MallGetCartResponse)

	// get user info
	if err != s.UserManager.GetUserInfo(ctx, req.UserId) {
		if errors.Is(err, pkg.ErrNoSuchUser) {
			resp.CommonResp = response.NewCommonResp(errz.NewErrZ(errz.WithErr(err), errz.WithCode(errz.CodeGetUserInfo)))
			return resp, nil
		}
		resp.CommonResp = response.NewCommonResp(errz.ErrGetUserInfo)
		log.Zlogger.Errorf("get user info failed err:%s", err.Error())
		return resp, nil
	}

	// get cart
	products, err := s.Dao.GetCart(ctx, req.UserId)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrGetCart)
		log.Zlogger.Errorf("get cart failed err:%s", err.Error())
		return resp, nil
	}

	// build response
	for _, product := range products {
		cartProduct := new(common.CartProduct)
		// check whether valid
		productInfo, err := s.ProductManager.GetProductInfo(ctx, product.ProductId)
		if err != nil {
			log.Zlogger.Errorf("get product info failed err:%s", err.Error())
			continue
		}
		if productInfo.Stock < product.ProductNum {
			cartProduct.IsInvalid = false
			continue
		}
		cartProduct.ProductId = product.ProductId
		cartProduct.ProductNum = product.ProductNum
		cartProduct.Amount = product.Amount
		cartProduct.AddTime = product.AddTime
		cartProduct.MerchantId = product.MerchantId
		cartProduct.IsInvalid = true
		resp.Cart.CartProducts = append(resp.Cart.CartProducts, cartProduct)
	}
	resp.CommonResp = response.NewCommonResp(nil)
	return resp, nil
}

// DelCartProduct implements the CartServiceImpl interface.
func (s *CartServiceImpl) DelCartProduct(ctx context.Context, req *cart.MallDelCartProductRequest) (resp *cart.MallDelCartProductResponse, err error) {
	// TODO: Your code here...
	resp = new(cart.MallDelCartProductResponse)

	// get user info
	if err != s.UserManager.GetUserInfo(ctx, req.UserId) {
		if errors.Is(err, pkg.ErrNoSuchUser) {
			resp.CommonResp = response.NewCommonResp(errz.NewErrZ(errz.WithErr(err), errz.WithCode(errz.CodeGetUserInfo)))
			return resp, nil
		}
		resp.CommonResp = response.NewCommonResp(errz.ErrGetUserInfo)
		log.Zlogger.Errorf("get user info failed err:%s", err.Error())
		return resp, nil
	}

	// del product
	err = s.Dao.DelProduct(ctx, req.UserId, req.ProductId)
	if err != nil {
		if errors.Is(err, pkg.ErrNoSuchProduct) {
			resp.CommonResp = response.NewCommonResp(errz.ErrNoProduct)
			return resp, nil
		}
		resp.CommonResp = response.NewCommonResp(errz.ErrCartInternal)
		log.Zlogger.Errorf("del product in cart failed err:%s", err.Error())
		return resp, nil
	}

	resp.CommonResp = response.NewCommonResp(nil)
	return resp, nil
}
