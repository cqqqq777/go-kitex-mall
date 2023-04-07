package pkg

import (
	"context"

	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/common"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/merchant/merchantservice"
)

type MerchantManager struct {
	MerchantService *merchantservice.Client
}

func (m *MerchantManager) GetMerchantInfo(ctx context.Context, MerchantId int64) (*common.Merchant, error) {
	return nil, nil
}
