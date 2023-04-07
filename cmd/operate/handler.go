package main

import (
	"context"
	operate "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/operate"
)

// OperateServiceImpl implements the last service interface defined in the IDL.
type OperateServiceImpl struct{}

// FavoriteProduct implements the OperateServiceImpl interface.
func (s *OperateServiceImpl) FavoriteProduct(ctx context.Context, req *operate.MallFavoriteProductRequest) (resp *operate.MallFavoriteProductResponse, err error) {
	// TODO: Your code here...
	return
}

// GetProductOperateInfo implements the OperateServiceImpl interface.
func (s *OperateServiceImpl) GetProductOperateInfo(ctx context.Context, req *operate.MallGetProductOperateInfoRequest) (resp *operate.MallGetProductOperateInfoResponse, err error) {
	// TODO: Your code here...
	return
}
