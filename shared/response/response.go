package response

import (
	"github.com/cqqqq777/go-kitex-mall/shared/errz"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/common"
)

func NewCommonResp(err *errz.ErrZ) *common.CommonResponse {
	if err == nil {
		return NewCommonResp(errz.NewErrZ())
	}
	return &common.CommonResponse{
		Code: err.GetCode(),
		Msg:  err.Error(),
	}
}
