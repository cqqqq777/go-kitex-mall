package pkg

import (
	"context"

	"github.com/cqqqq777/go-kitex-mall/shared/errz"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/common"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/merchant"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/merchant/merchantservice"
)

type MerchantManager struct {
	MerchantService merchantservice.Client
}

func (m *MerchantManager) GetInfo(ctx context.Context, MerchantId int64) (*common.Merchant, error) {
	resp, err := m.MerchantService.GetInfo(ctx, &merchant.MallMerchantGetInfoRequest{
		Id: MerchantId,
	})
	if err != nil {
		return nil, err
	}
	if resp.CommonResp.Code != errz.Success {
		return nil, errz.NewErrZ(errz.WithErr(err))
	}
	return resp.MerchantInfo, nil
}
