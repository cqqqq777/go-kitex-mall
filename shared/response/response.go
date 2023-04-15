package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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

func SendResp(c *app.RequestContext, data interface{}) {
	c.JSON(consts.StatusOK, data)
}
