package pkg

import (
	"context"
	"fmt"
	"github.com/cqqqq777/go-kitex-mall/shared/errz"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/user"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/user/userservice"
)

type UserManager struct {
	UserService userservice.Client
}

func (u UserManager) GetUserInfo(ctx context.Context, uid int64) error {
	resp, err := u.UserService.GetUserInfo(ctx, &user.MallGetUserInfoRequest{
		Id: uid,
	})
	if err != nil {
		return err
	}
	if resp.CommonResp.Code != errz.Success {
		return fmt.Errorf("%s", resp.CommonResp.Msg)
	}
	return nil
}
