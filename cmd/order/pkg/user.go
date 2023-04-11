package pkg

import (
	"context"
	"errors"
	"github.com/cqqqq777/go-kitex-mall/shared/errz"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/user"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/user/userservice"
)

var ErrGetUserInfoFailed = errors.New("get user info failed")

type UserManager struct {
	UserService userservice.Client
}

func (u UserManager) GetUserInfo(ctx context.Context, uid int64) error {
	info, _ := u.UserService.GetUserInfo(ctx, &user.MallGetUserInfoRequest{
		Id: uid,
	})
	if info.CommonResp.Code != errz.Success {
		return ErrGetUserInfoFailed
	}
	return nil
}
