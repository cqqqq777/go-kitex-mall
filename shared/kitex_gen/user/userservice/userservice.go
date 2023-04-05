// Code generated by Kitex v0.5.1. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	user "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/user"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetVerification":  kitex.NewMethodInfo(getVerificationHandler, newUserServiceGetVerificationArgs, newUserServiceGetVerificationResult, false),
		"Register":         kitex.NewMethodInfo(registerHandler, newUserServiceRegisterArgs, newUserServiceRegisterResult, false),
		"Login":            kitex.NewMethodInfo(loginHandler, newUserServiceLoginArgs, newUserServiceLoginResult, false),
		"GetUserInfo":      kitex.NewMethodInfo(getUserInfoHandler, newUserServiceGetUserInfoArgs, newUserServiceGetUserInfoResult, false),
		"ChangeAvatar":     kitex.NewMethodInfo(changeAvatarHandler, newUserServiceChangeAvatarArgs, newUserServiceChangeAvatarResult, false),
		"ChangeBackground": kitex.NewMethodInfo(changeBackgroundHandler, newUserServiceChangeBackgroundArgs, newUserServiceChangeBackgroundResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.1",
		Extra:           extra,
	}
	return svcInfo
}

func getVerificationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetVerificationArgs)
	realResult := result.(*user.UserServiceGetVerificationResult)
	success, err := handler.(user.UserService).GetVerification(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetVerificationArgs() interface{} {
	return user.NewUserServiceGetVerificationArgs()
}

func newUserServiceGetVerificationResult() interface{} {
	return user.NewUserServiceGetVerificationResult()
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceRegisterArgs)
	realResult := result.(*user.UserServiceRegisterResult)
	success, err := handler.(user.UserService).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRegisterArgs() interface{} {
	return user.NewUserServiceRegisterArgs()
}

func newUserServiceRegisterResult() interface{} {
	return user.NewUserServiceRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceLoginArgs)
	realResult := result.(*user.UserServiceLoginResult)
	success, err := handler.(user.UserService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceLoginArgs() interface{} {
	return user.NewUserServiceLoginArgs()
}

func newUserServiceLoginResult() interface{} {
	return user.NewUserServiceLoginResult()
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetUserInfoArgs)
	realResult := result.(*user.UserServiceGetUserInfoResult)
	success, err := handler.(user.UserService).GetUserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserInfoArgs() interface{} {
	return user.NewUserServiceGetUserInfoArgs()
}

func newUserServiceGetUserInfoResult() interface{} {
	return user.NewUserServiceGetUserInfoResult()
}

func changeAvatarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceChangeAvatarArgs)
	realResult := result.(*user.UserServiceChangeAvatarResult)
	success, err := handler.(user.UserService).ChangeAvatar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceChangeAvatarArgs() interface{} {
	return user.NewUserServiceChangeAvatarArgs()
}

func newUserServiceChangeAvatarResult() interface{} {
	return user.NewUserServiceChangeAvatarResult()
}

func changeBackgroundHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceChangeBackgroundArgs)
	realResult := result.(*user.UserServiceChangeBackgroundResult)
	success, err := handler.(user.UserService).ChangeBackground(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceChangeBackgroundArgs() interface{} {
	return user.NewUserServiceChangeBackgroundArgs()
}

func newUserServiceChangeBackgroundResult() interface{} {
	return user.NewUserServiceChangeBackgroundResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetVerification(ctx context.Context, req *user.MallVerificationRequest) (r *user.MallVerificationResponse, err error) {
	var _args user.UserServiceGetVerificationArgs
	_args.Req = req
	var _result user.UserServiceGetVerificationResult
	if err = p.c.Call(ctx, "GetVerification", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Register(ctx context.Context, req *user.MallUserRegisterRequest) (r *user.MallUserRegisterResponse, err error) {
	var _args user.UserServiceRegisterArgs
	_args.Req = req
	var _result user.UserServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, req *user.MallUserLoginRequest) (r *user.MallUserLoginResponse, err error) {
	var _args user.UserServiceLoginArgs
	_args.Req = req
	var _result user.UserServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, req *user.MallGetUserInfoRequest) (r *user.MallGetUserInfoResponse, err error) {
	var _args user.UserServiceGetUserInfoArgs
	_args.Req = req
	var _result user.UserServiceGetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChangeAvatar(ctx context.Context, req *user.MallChangeUserAvatarRequest) (r *user.MallChangeUserAvatarResponse, err error) {
	var _args user.UserServiceChangeAvatarArgs
	_args.Req = req
	var _result user.UserServiceChangeAvatarResult
	if err = p.c.Call(ctx, "ChangeAvatar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChangeBackground(ctx context.Context, req *user.MallChangeUserBackgroundRequest) (r *user.MallChangeUserBackgroundResponse, err error) {
	var _args user.UserServiceChangeBackgroundArgs
	_args.Req = req
	var _result user.UserServiceChangeBackgroundResult
	if err = p.c.Call(ctx, "ChangeBackground", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}