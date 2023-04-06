package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/cqqqq777/go-kitex-mall/cmd/user/config"
	"github.com/cqqqq777/go-kitex-mall/cmd/user/dao"
	"github.com/cqqqq777/go-kitex-mall/cmd/user/model"
	"github.com/cqqqq777/go-kitex-mall/cmd/user/pkg"
	"github.com/cqqqq777/go-kitex-mall/shared/consts"
	"github.com/cqqqq777/go-kitex-mall/shared/errz"
	"github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/user"
	"github.com/cqqqq777/go-kitex-mall/shared/log"
	"github.com/cqqqq777/go-kitex-mall/shared/middleware"
	"github.com/cqqqq777/go-kitex-mall/shared/response"

	"github.com/bwmarrin/snowflake"
	"github.com/golang-jwt/jwt"
	"gopkg.in/gomail.v2"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	Jwt *middleware.JWT
	Dao *dao.User
	Producer
}

type Producer interface {
	Produce(user pkg.Msg) error
}

// GetVerification implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetVerification(ctx context.Context, req *user.MallVerificationRequest) (resp *user.MallVerificationResponse, err error) {
	// TODO: Your code here...
	resp = new(user.MallVerificationResponse)

	vCode := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)
	message := fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>你本次的验证码为%06d,为了保证账号安全，验证码有效期为10分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>`, vCode)

	m := gomail.NewMessage()
	m.SetHeader("From", config.GlobalServerConfig.EmailInfo.Email) // 发件人
	// m.SetHeader("From", "alias"+"<"+userName+">") // 增加发件人别名

	m.SetHeader("To", req.Email) // 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
	//m.SetHeader("Cc", "******@qq.com")                  // 抄送，可以多个
	//m.SetHeader("Bcc", "******@qq.com")                 // 暗送，可以多个
	m.SetHeader("Subject", "Hello!") // 邮件主题

	// text/html 的意思是将文件的 content-type 设置为 text/html 的形式，浏览器在获取到这种文件时会自动调用html的解析器对文件进行相应的处理。
	// 可以通过 text/html 处理文本格式进行特殊处理，如换行、缩进、加粗等等
	m.SetBody("text/html", message)

	// text/plain的意思是将文件设置为纯文本的形式，浏览器在获取到这种文件时并不会对其进行处理
	// m.SetBody("text/plain", "纯文本")
	// m.Attach("test.sh")   // 附件文件，可以是文件，照片，视频等等
	// m.Attach("lolcatVideo.mp4") // 视频
	// m.Attach("lolcat.jpg") // 照片

	d := gomail.NewDialer(
		config.GlobalServerConfig.EmailInfo.Host,
		config.GlobalServerConfig.EmailInfo.Port,
		config.GlobalServerConfig.EmailInfo.Email,
		config.GlobalServerConfig.EmailInfo.Password,
	)
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err = d.DialAndSend(m)
	if err != nil {
		log.Zlogger.Errorf("send email failed err:%s", err.Error())
		resp.ComonResp = response.NewCommonResp(errz.ErrSentVerification)
		return resp, nil
	}
	if err = s.Dao.SetVerification(ctx, req.Email, vCode); err != nil {
		log.Zlogger.Errorf("set verification in redis failed err:%s", err.Error())
		resp.ComonResp = response.NewCommonResp(errz.ErrUserInternal)
		return resp, nil
	}
	resp.ComonResp = response.NewCommonResp(nil)
	return resp, nil
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.MallUserRegisterRequest) (resp *user.MallUserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(user.MallUserRegisterResponse)

	// check verification
	verification, err := s.Dao.GetVerification(ctx, req.Email)
	if verification == 0 {
		resp.CommonResp = response.NewCommonResp(errz.ErrGetVerification)
		return resp, nil
	}
	if verification != req.Verification {
		resp.CommonResp = response.NewCommonResp(errz.ErrWrongVerification)
		return resp, nil
	}

	// create user in mysql and mongodb
	sf, err := snowflake.NewNode(consts.UserSnowflakeNode)
	if err != nil {
		log.Zlogger.Errorf("generate user id faield err:%s" + err.Error())
		resp.CommonResp = response.NewCommonResp(errz.ErrGenerateUid)
		return resp, nil
	}
	id := sf.Generate().Int64()
	User := &model.User{
		Id:       id,
		Username: req.Username,
		Password: pkg.Md5(req.Password),
		Email:    req.Email,
	}
	if err = s.Dao.CreateUserInMysql(User); err != nil {
		if errors.Is(err, dao.ErrUserExist) {
			resp.CommonResp = response.NewCommonResp(errz.ErrUserExist)
			return resp, nil
		} else {
			log.Zlogger.Errorf("create user in mysql failed err:%s" + err.Error())
			resp.CommonResp = response.NewCommonResp(errz.ErrUserInternal)
			return resp, nil
		}
	}
	userM := &model.UserM{
		Id:       id,
		Username: req.Username,
	}
	if err = s.Dao.CreateUserInMongo(ctx, userM); err != nil {
		log.Zlogger.Errorf("create user in mongodb failed err:%s" + err.Error())
		resp.CommonResp = response.NewCommonResp(errz.ErrUserInternal)
		return resp, nil
	}

	resp.Token, err = s.Jwt.CreateToken(middleware.CustomClaims{
		ID:       id,
		Identity: consts.UserIdentity,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + consts.TokenExpiredAt,
		},
	})
	if err != nil {
		log.Zlogger.Errorf("generate token failed failed err:%s" + err.Error())
		resp.CommonResp = response.NewCommonResp(errz.ErrGenerateToken)
		return resp, nil
	}
	resp.UserId = id
	resp.CommonResp = response.NewCommonResp(nil)

	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.MallUserLoginRequest) (resp *user.MallUserLoginResponse, err error) {
	// TODO: Your code here...
	// check password
	usr, err := s.Dao.GetUserByUsername(req.Username)
	req.Password = pkg.Md5(req.Password)
	if usr.Password != req.Password {
		resp.CommonResp = response.NewCommonResp(errz.ErrWrongPassword)
		return resp, nil
	}

	// generate token
	resp.Token, err = s.Jwt.CreateToken(middleware.CustomClaims{
		ID: usr.Id,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + consts.TokenExpiredAt,
		},
	})
	if err != nil {
		log.Zlogger.Errorf("generate token failed failed err:%s" + err.Error())
		resp.CommonResp = response.NewCommonResp(errz.ErrGenerateToken)
		return resp, nil
	}
	resp.UserId = usr.Id

	resp.CommonResp = response.NewCommonResp(nil)
	return resp, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.MallGetUserInfoRequest) (resp *user.MallGetUserInfoResponse, err error) {
	// TODO: Your code here...
	resp = new(user.MallGetUserInfoResponse)

	//get user info in mongodb
	userInfo, err := s.Dao.GetUserInfo(ctx, req.Id)
	if err != nil {
		log.Zlogger.Errorf("get user info failed failed err:%s" + err.Error())
		resp.CommonResp = response.NewCommonResp(errz.ErrGetUserInfo)
		return resp, nil
	}
	resp.UserInfo.Id = userInfo.Id
	resp.UserInfo.Name = userInfo.Username
	resp.UserInfo.Avatar = userInfo.Avatar
	resp.UserInfo.Background = userInfo.Background
	resp.UserInfo.Signature = userInfo.Signature
	resp.CommonResp = response.NewCommonResp(nil)

	// cache user info
	err = s.Dao.CacheUserInfo(ctx, userInfo)
	if err != nil {
		log.Zlogger.Errorf("cache user info failed err:%s", err.Error())
	}

	return resp, nil
}

// ChangeAvatar implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangeAvatar(ctx context.Context, req *user.MallChangeUserAvatarRequest) (resp *user.MallChangeUserAvatarResponse, err error) {
	// TODO: Your code here...
	resp = new(user.MallChangeUserAvatarResponse)

	// change avatar in mongodb
	if err = s.Dao.ChangeAvatar(ctx, req.Id, req.Avatar); err != nil {
		log.Zlogger.Errorf("change user:%d avatar failed err:%s", req.Id, err.Error())
		resp.CommonResp = response.NewCommonResp(errz.ErrChangeAvatar)
		return resp, nil
	}

	// publish message in nsq
	var msg pkg.Msg
	msg.Id = req.Id
	msg.Message = req.Avatar
	msg.Type = consts.ObjectAvatarType
	err = s.Producer.Produce(msg)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrPublishMsgInNsq)
		return resp, nil
	}

	// clear user info cache
	err = s.Dao.ClearUserInfoCache(ctx, req.Id)
	if err != nil {
		log.Zlogger.Errorf("clear user info cache failed err:%s", err.Error())
	}

	resp.CommonResp = response.NewCommonResp(nil)
	return resp, nil
}

// ChangeBackground implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangeBackground(ctx context.Context, req *user.MallChangeUserBackgroundRequest) (resp *user.MallChangeUserBackgroundResponse, err error) {
	// TODO: Your code here...
	resp = new(user.MallChangeUserBackgroundResponse)

	// change background in mongodb
	if err = s.Dao.ChangeBackground(ctx, req.Id, req.Background); err != nil {
		log.Zlogger.Errorf("change user:%d background failed err:%s", req.Id, err.Error())
		resp.CommonResp = response.NewCommonResp(errz.ErrChangeBackground)
		return resp, nil
	}

	// publish message in nsq
	var msg pkg.Msg
	msg.Id = req.Id
	msg.Message = req.Background
	msg.Type = consts.ObjectBackgroundType
	err = s.Producer.Produce(msg)
	if err != nil {
		resp.CommonResp = response.NewCommonResp(errz.ErrPublishMsgInNsq)
		return resp, nil
	}

	// clear user info cache
	err = s.Dao.ClearUserInfoCache(ctx, req.Id)
	if err != nil {
		log.Zlogger.Errorf("clear user info cache failed err:%s", err.Error())
	}

	resp.CommonResp = response.NewCommonResp(nil)
	return resp, nil
}
