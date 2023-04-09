// Code generated by Kitex v0.5.1. DO NOT EDIT.

package productservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	product "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/product"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	PublishProduct(ctx context.Context, req *product.MallPublishProductRequest, callOptions ...callopt.Option) (r *product.MallPublishProductResponse, err error)
	DelProduct(ctx context.Context, req *product.MallDelProductRequest, callOptions ...callopt.Option) (r *product.MallDelProductResponse, err error)
	ProductList(ctx context.Context, req *product.MallProductListRequest, callOptions ...callopt.Option) (r *product.MallProductListResponse, err error)
	ProductDetail(ctx context.Context, req *product.MallProductDetailRequest, callOptions ...callopt.Option) (r *product.MallProductDetailResponse, err error)
	SearchProduct(ctx context.Context, req *product.MallSearchProductRequest, callOptions ...callopt.Option) (r *product.MallSearchProductResponse, err error)
	ProductFavoriteList(ctx context.Context, req *product.MallProductFavoriteListRequest, callOptions ...callopt.Option) (r *product.MallProductFavoriteListResponse, err error)
	PublishedProducts(ctx context.Context, req *product.MallProductPublishedListRequest, callOptions ...callopt.Option) (r *product.MallProductPublishedListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kProductServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kProductServiceClient struct {
	*kClient
}

func (p *kProductServiceClient) PublishProduct(ctx context.Context, req *product.MallPublishProductRequest, callOptions ...callopt.Option) (r *product.MallPublishProductResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishProduct(ctx, req)
}

func (p *kProductServiceClient) DelProduct(ctx context.Context, req *product.MallDelProductRequest, callOptions ...callopt.Option) (r *product.MallDelProductResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DelProduct(ctx, req)
}

func (p *kProductServiceClient) ProductList(ctx context.Context, req *product.MallProductListRequest, callOptions ...callopt.Option) (r *product.MallProductListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ProductList(ctx, req)
}

func (p *kProductServiceClient) ProductDetail(ctx context.Context, req *product.MallProductDetailRequest, callOptions ...callopt.Option) (r *product.MallProductDetailResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ProductDetail(ctx, req)
}

func (p *kProductServiceClient) SearchProduct(ctx context.Context, req *product.MallSearchProductRequest, callOptions ...callopt.Option) (r *product.MallSearchProductResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchProduct(ctx, req)
}

func (p *kProductServiceClient) ProductFavoriteList(ctx context.Context, req *product.MallProductFavoriteListRequest, callOptions ...callopt.Option) (r *product.MallProductFavoriteListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ProductFavoriteList(ctx, req)
}

func (p *kProductServiceClient) PublishedProducts(ctx context.Context, req *product.MallProductPublishedListRequest, callOptions ...callopt.Option) (r *product.MallProductPublishedListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishedProducts(ctx, req)
}
