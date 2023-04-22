package pkg

import (
	"context"
	"errors"
	"github.com/cqqqq777/go-kitex-mall/shared/errz"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/user"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/user/userservice"
)

var ErrNoSuchUser = errors.New("no such user")

type UserManager struct {
	UserService userservice.Client
}

func (u *UserManager) GetUserInfo(ctx context.Context, userId int64) error {
	resp, err := u.UserService.GetUserInfo(ctx, &user.MallGetUserInfoRequest{
		Id: userId,
	})
	if err != nil {
		return err
	}
	if resp.CommonResp.Code != errz.Success {
		return ErrNoSuchUser
	}
	return nil
}
