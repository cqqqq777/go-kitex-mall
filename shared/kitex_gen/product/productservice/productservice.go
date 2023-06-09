// Code generated by Kitex v0.5.1. DO NOT EDIT.

package productservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	product "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/product"
)

func serviceInfo() *kitex.ServiceInfo {
	return productServiceServiceInfo
}

var productServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ProductService"
	handlerType := (*product.ProductService)(nil)
	methods := map[string]kitex.MethodInfo{
		"PublishProduct":      kitex.NewMethodInfo(publishProductHandler, newProductServicePublishProductArgs, newProductServicePublishProductResult, false),
		"UpdateProduct":       kitex.NewMethodInfo(updateProductHandler, newProductServiceUpdateProductArgs, newProductServiceUpdateProductResult, false),
		"DelProduct":          kitex.NewMethodInfo(delProductHandler, newProductServiceDelProductArgs, newProductServiceDelProductResult, false),
		"ProductList":         kitex.NewMethodInfo(productListHandler, newProductServiceProductListArgs, newProductServiceProductListResult, false),
		"ProductDetail":       kitex.NewMethodInfo(productDetailHandler, newProductServiceProductDetailArgs, newProductServiceProductDetailResult, false),
		"SearchProduct":       kitex.NewMethodInfo(searchProductHandler, newProductServiceSearchProductArgs, newProductServiceSearchProductResult, false),
		"ProductFavoriteList": kitex.NewMethodInfo(productFavoriteListHandler, newProductServiceProductFavoriteListArgs, newProductServiceProductFavoriteListResult, false),
		"PublishedProducts":   kitex.NewMethodInfo(publishedProductsHandler, newProductServicePublishedProductsArgs, newProductServicePublishedProductsResult, false),
		"UpdateStock":         kitex.NewMethodInfo(updateStockHandler, newProductServiceUpdateStockArgs, newProductServiceUpdateStockResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "product",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.1",
		Extra:           extra,
	}
	return svcInfo
}

func publishProductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.ProductServicePublishProductArgs)
	realResult := result.(*product.ProductServicePublishProductResult)
	success, err := handler.(product.ProductService).PublishProduct(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newProductServicePublishProductArgs() interface{} {
	return product.NewProductServicePublishProductArgs()
}

func newProductServicePublishProductResult() interface{} {
	return product.NewProductServicePublishProductResult()
}

func updateProductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.ProductServiceUpdateProductArgs)
	realResult := result.(*product.ProductServiceUpdateProductResult)
	success, err := handler.(product.ProductService).UpdateProduct(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newProductServiceUpdateProductArgs() interface{} {
	return product.NewProductServiceUpdateProductArgs()
}

func newProductServiceUpdateProductResult() interface{} {
	return product.NewProductServiceUpdateProductResult()
}

func delProductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.ProductServiceDelProductArgs)
	realResult := result.(*product.ProductServiceDelProductResult)
	success, err := handler.(product.ProductService).DelProduct(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newProductServiceDelProductArgs() interface{} {
	return product.NewProductServiceDelProductArgs()
}

func newProductServiceDelProductResult() interface{} {
	return product.NewProductServiceDelProductResult()
}

func productListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.ProductServiceProductListArgs)
	realResult := result.(*product.ProductServiceProductListResult)
	success, err := handler.(product.ProductService).ProductList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newProductServiceProductListArgs() interface{} {
	return product.NewProductServiceProductListArgs()
}

func newProductServiceProductListResult() interface{} {
	return product.NewProductServiceProductListResult()
}

func productDetailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.ProductServiceProductDetailArgs)
	realResult := result.(*product.ProductServiceProductDetailResult)
	success, err := handler.(product.ProductService).ProductDetail(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newProductServiceProductDetailArgs() interface{} {
	return product.NewProductServiceProductDetailArgs()
}

func newProductServiceProductDetailResult() interface{} {
	return product.NewProductServiceProductDetailResult()
}

func searchProductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.ProductServiceSearchProductArgs)
	realResult := result.(*product.ProductServiceSearchProductResult)
	success, err := handler.(product.ProductService).SearchProduct(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newProductServiceSearchProductArgs() interface{} {
	return product.NewProductServiceSearchProductArgs()
}

func newProductServiceSearchProductResult() interface{} {
	return product.NewProductServiceSearchProductResult()
}

func productFavoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.ProductServiceProductFavoriteListArgs)
	realResult := result.(*product.ProductServiceProductFavoriteListResult)
	success, err := handler.(product.ProductService).ProductFavoriteList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newProductServiceProductFavoriteListArgs() interface{} {
	return product.NewProductServiceProductFavoriteListArgs()
}

func newProductServiceProductFavoriteListResult() interface{} {
	return product.NewProductServiceProductFavoriteListResult()
}

func publishedProductsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.ProductServicePublishedProductsArgs)
	realResult := result.(*product.ProductServicePublishedProductsResult)
	success, err := handler.(product.ProductService).PublishedProducts(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newProductServicePublishedProductsArgs() interface{} {
	return product.NewProductServicePublishedProductsArgs()
}

func newProductServicePublishedProductsResult() interface{} {
	return product.NewProductServicePublishedProductsResult()
}

func updateStockHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.ProductServiceUpdateStockArgs)
	realResult := result.(*product.ProductServiceUpdateStockResult)
	success, err := handler.(product.ProductService).UpdateStock(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newProductServiceUpdateStockArgs() interface{} {
	return product.NewProductServiceUpdateStockArgs()
}

func newProductServiceUpdateStockResult() interface{} {
	return product.NewProductServiceUpdateStockResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PublishProduct(ctx context.Context, req *product.MallPublishProductRequest) (r *product.MallPublishProductResponse, err error) {
	var _args product.ProductServicePublishProductArgs
	_args.Req = req
	var _result product.ProductServicePublishProductResult
	if err = p.c.Call(ctx, "PublishProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateProduct(ctx context.Context, req *product.MallUpdateProductRequest) (r *product.MallUpdateProductResponse, err error) {
	var _args product.ProductServiceUpdateProductArgs
	_args.Req = req
	var _result product.ProductServiceUpdateProductResult
	if err = p.c.Call(ctx, "UpdateProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DelProduct(ctx context.Context, req *product.MallDelProductRequest) (r *product.MallDelProductResponse, err error) {
	var _args product.ProductServiceDelProductArgs
	_args.Req = req
	var _result product.ProductServiceDelProductResult
	if err = p.c.Call(ctx, "DelProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ProductList(ctx context.Context, req *product.MallProductListRequest) (r *product.MallProductListResponse, err error) {
	var _args product.ProductServiceProductListArgs
	_args.Req = req
	var _result product.ProductServiceProductListResult
	if err = p.c.Call(ctx, "ProductList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ProductDetail(ctx context.Context, req *product.MallProductDetailRequest) (r *product.MallProductDetailResponse, err error) {
	var _args product.ProductServiceProductDetailArgs
	_args.Req = req
	var _result product.ProductServiceProductDetailResult
	if err = p.c.Call(ctx, "ProductDetail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SearchProduct(ctx context.Context, req *product.MallSearchProductRequest) (r *product.MallSearchProductResponse, err error) {
	var _args product.ProductServiceSearchProductArgs
	_args.Req = req
	var _result product.ProductServiceSearchProductResult
	if err = p.c.Call(ctx, "SearchProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ProductFavoriteList(ctx context.Context, req *product.MallProductFavoriteListRequest) (r *product.MallProductFavoriteListResponse, err error) {
	var _args product.ProductServiceProductFavoriteListArgs
	_args.Req = req
	var _result product.ProductServiceProductFavoriteListResult
	if err = p.c.Call(ctx, "ProductFavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishedProducts(ctx context.Context, req *product.MallProductPublishedListRequest) (r *product.MallProductPublishedListResponse, err error) {
	var _args product.ProductServicePublishedProductsArgs
	_args.Req = req
	var _result product.ProductServicePublishedProductsResult
	if err = p.c.Call(ctx, "PublishedProducts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateStock(ctx context.Context, req *product.MallUpdateStockRequest) (r *product.MallUpdateStockResponse, err error) {
	var _args product.ProductServiceUpdateStockArgs
	_args.Req = req
	var _result product.ProductServiceUpdateStockResult
	if err = p.c.Call(ctx, "UpdateStock", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
