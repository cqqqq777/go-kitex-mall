package pkg

import (
	"github.com/bytedance/sonic"
	"github.com/cqqqq777/go-kitex-mall/cmd/cart/model"
	"github.com/cqqqq777/go-kitex-mall/shared/log"
)

func UnMarshalProduct(val map[string]string) (products []*model.CartProduct, err error) {
	if len(val) == 0 {
		return nil, nil
	}
	products = make([]*model.CartProduct, 0, len(val))
	for _, v := range val {
		product := new(model.CartProduct)
		err = sonic.UnmarshalString(v, product)
		if err != nil {
			log.Zlogger.Errorf("unmarshal product failed err:%s", err.Error())
			continue
		}
		products = append(products, product)
	}
	return products, nil
}
