package pkg

import (
	hc "github.com/cqqqq777/go-kitex-mall/cmd/api/biz/model/common"
	kc "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/common"
)

func Merchant(m *kc.Merchant) *hc.Merchant {
	if m == nil {
		return nil
	}
	return &hc.Merchant{
		ID:          m.Id,
		Name:        m.Name,
		Alipay:      m.Alipay,
		Description: m.Description,
	}
}

func Image(i *kc.Image) *hc.Image {
	if i == nil {
		return nil
	}
	return &hc.Image{
		ID:   i.Id,
		Path: i.Path,
	}
}

func ProductOperateInfo(p *kc.ProductOperateInfo) *hc.ProductOperateInfo {
	if p == nil {
		return nil
	}
	return &hc.ProductOperateInfo{
		IsFavorite:   p.IsFavorite,
		CommentCount: p.CommentCount,
		SaleCount:    p.SaleCount,
	}
}

func ProductDetail(p *kc.ProductDetail) *hc.ProductDetail {
	if p == nil {
		return nil
	}
	return &hc.ProductDetail{
		BasicInfo:    Product(p.BasicInfo),
		OperateInfo:  ProductOperateInfo(p.OperateInfo),
		MerchantInfo: Merchant(p.MerchantInfo),
		CreateTime:   p.CreateTime,
		UpdateTime:   p.UpdateTime,
	}
}

func Product(p *kc.Product) *hc.Product {
	if p == nil {
		return nil
	}
	images := make([]*hc.Image, 0)
	for _, i := range p.Iamges {
		if v := Image(i); v != nil {
			images = append(images, v)
		}
	}
	return &hc.Product{
		Stock:       p.Stock,
		Price:       p.Price,
		ID:          p.Id,
		MID:         p.MId,
		Name:        p.Name,
		Description: p.Description,
		Status:      p.Status,
		Iamges:      images,
	}
}

func Products(products []*kc.Product) []*hc.Product {
	ps := make([]*hc.Product, 0)
	for _, product := range products {
		if v := Product(product); v != nil {
			ps = append(ps, v)
		}
	}
	return ps
}

func Order(o *kc.Order) *hc.Order {
	if o == nil {
		return nil
	}
	return &hc.Order{
		OrderID:    o.OrderId,
		UserID:     o.UserId,
		ProductID:  o.ProductId,
		ProductNum: o.ProductNum,
		Amount:     o.Amount,
		Status:     o.Status,
		CreateTime: o.CreateTime,
		UpdateTime: o.UpdateTime,
		ExpTime:    o.ExpTime,
	}
}

func Orders(orders []*kc.Order) []*hc.Order {
	os := make([]*hc.Order, 0)
	for _, o := range orders {
		if v := Order(o); v != nil {
			os = append(os, v)
		}
	}
	return os
}

func Pay(p *kc.Pay) *hc.Pay {
	if p == nil {
		return nil
	}
	return &hc.Pay{
		PayID:      p.PayId,
		OrderID:    p.OrderId,
		URL:        p.Url,
		UserID:     p.UserId,
		Status:     p.Status,
		Amount:     p.Amount,
		CreateTime: p.CreateTime,
		UpdateTime: p.UpdateTime,
	}
}
