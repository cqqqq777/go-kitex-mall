package main

import (
	"context"
	"github.com/cqqqq777/go-kitex-mall/cmd/product/dao"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/common"
	product "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/product"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct {
	Dao *dao.Product
	MerchantManager
	Producer
	OperateManager
}

type MerchantManager interface {
	GetMerchantInfo(ctx context.Context, MerchantId int64) (*common.Merchant, error)
}

type Producer interface {
	Produce()
}

type OperateManager interface {
	GetProductOperateInfo(ctx context.Context, ProductId int64) (*common.ProductOperateInfo, error)
}

// PublishProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) PublishProduct(ctx context.Context, req *product.MallPublishProductRequest) (resp *product.MallPublishProductResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, req *product.MallUpdateProductRequest) (resp *product.MallUpdateProductResponse, err error) {
	// TODO: Your code here...
	return
}

// DelProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DelProduct(ctx context.Context, req *product.MallDelProductRequest) (resp *product.MallDelProductResponse, err error) {
	// TODO: Your code here...
	return
}

// ProductList implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ProductList(ctx context.Context, req *product.MallProductListRequest) (resp *product.MallProductListResponse, err error) {
	// TODO: Your code here...
	return
}

// ProductDetail implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ProductDetail(ctx context.Context, req *product.MallProductDetailRequest) (resp *product.MallProductDetailResponse, err error) {
	// TODO: Your code here...
	return
}

// SearchProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) SearchProduct(ctx context.Context, req *product.MallSearchProductRequest) (resp *product.MallSearchProductResponse, err error) {
	// TODO: Your code here...
	return
}

// ProductFavoriteList implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ProductFavoriteList(ctx context.Context, req *product.MallProductFavoriteListRequest) (resp *product.MallProductFavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishedProducts implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) PublishedProducts(ctx context.Context, req *product.MallProductPublishedListRequest) (resp *product.MallProductPublishedListResponse, err error) {
	// TODO: Your code here...
	return
}
