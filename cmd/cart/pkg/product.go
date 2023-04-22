package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/cqqqq777/go-kitex-mall/shared/errz"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/product"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/product/productservice"
)

var ErrNoSuchProduct = errors.New("no such product")

type Product struct {
	ProductId  int64 `json:"product_id"`
	MerchantId int64 `json:"merchant_id"`
	Price      int64 `json:"price"`
	Stock      int64 `json:"stock"`
}

type ProductManager struct {
	ProductService productservice.Client
}

func (p *ProductManager) GetProductInfo(ctx context.Context, productId int64) (productInfo *Product, err error) {
	resp, err := p.ProductService.ProductDetail(ctx, &product.MallProductDetailRequest{ProductId: productId})
	if err != nil {
		return nil, err
	}
	if resp.CommonResp.Code != errz.Success {
		if resp.CommonResp.Code == errz.CodeNoProduct {
			return nil, ErrNoSuchProduct
		}
		return nil, fmt.Errorf(resp.CommonResp.Msg)
	}
	return &Product{
		Price:      resp.Product.BasicInfo.Price,
		ProductId:  productId,
		MerchantId: resp.Product.BasicInfo.MId,
		Stock:      resp.Product.BasicInfo.Stock,
	}, nil
}
