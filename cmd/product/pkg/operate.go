package pkg

import (
	"context"

	"github.com/cqqqq777/go-kitex-mall/shared/errz"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/common"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/operate"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/operate/operateservice"
)

type OperateManager struct {
	OperateService operateservice.Client
}

func (o *OperateManager) GetProductOperateInfo(ctx context.Context, uid, pid int64) (*common.ProductOperateInfo, error) {
	resp, err := o.OperateService.GetProductOperateInfo(ctx, &operate.MallGetProductOperateInfoRequest{
		ProductId: pid,
		UserId:    uid,
	})
	if err != nil {
		return nil, err
	}
	if resp.CommonResp.Code != errz.Success {
		return nil, errz.NewErrZ(errz.WithErr(err))
	}
	return resp.OperateInfo, nil
}
