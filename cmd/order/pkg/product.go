package pkg

import (
	"context"
	"errors"
	"github.com/cqqqq777/go-kitex-mall/shared/errz"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/common"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/product"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/product/productservice"
)

var ErrNoSuchProduct = errors.New("no such product")

type ProductManage struct {
	ProductService productservice.Client
}

func (p *ProductManage) GetProductInfo(ctx context.Context, pid int64) (*common.Product, error) {
	detail, err := p.ProductService.ProductDetail(ctx, &product.MallProductDetailRequest{ProductId: pid})
	if err != nil {
		return nil, err
	}
	if detail.CommonResp.Code != errz.Success {
		if detail.CommonResp.Code == errz.CodeNoProduct {
			return nil, ErrNoSuchProduct
		}
		return nil, errz.NewErrZ(errz.WithErr(err))
	}
	return detail.Product.BasicInfo, err
}

func (p *ProductManage) UpdateStock(ctx context.Context, pid, stock int64) error {
	_, err := p.ProductService.UpdateStock(ctx, &product.MallUpdateStockRequest{
		ProductId: pid,
		Stock:     stock,
	})
	return err
}
