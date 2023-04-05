package dao

import (
	"context"
	"errors"
	"github.com/cqqqq777/go-kitex-mall/cmd/user/model"
	"github.com/cqqqq777/go-kitex-mall/shared/consts"
	"github.com/cqqqq777/go-kitex-mall/shared/rdk"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("no such user")
	ErrUserExist    = errors.New("user already exist")
)

type User struct {
	db  *gorm.DB
	mdb *mongo.Database
	rdb *redis.Client
}

func NewUser(db *gorm.DB, mdb *mongo.Database, rdb *redis.Client) *User {
	m := db.Migrator()
	if !m.HasTable(&model.User{}) {
		err := m.CreateTable(&model.User{})
		if err != nil {
			panic(err)
		}
	}
	return &User{
		db:  db,
		mdb: mdb,
		rdb: rdb,
	}
}

func (u *User) SetVerification(ctx context.Context, email string, vCode int32) error {
	return u.rdb.SetEx(ctx, rdk.GetVerificationKey(email), vCode, consts.VerificationExpTime).Err()
}

func (u *User) GetVerification(ctx context.Context, email string) (int64, error) {
	cmd := u.rdb.Get(ctx, email)
	return cmd.Int64()
}

func (u *User) CreateUserInMysql(user *model.User) error {
	err := u.db.Model(&model.User{}).
		Where(&model.User{Username: user.Username}).First(&model.User{}).Error
	if err == nil {
		return ErrUserExist
	} else if err != gorm.ErrRecordNotFound {
		return err
	}
	return u.db.Model(&model.User{}).Create(user).Error
}

func (u *User) CreateUserInMongo(ctx context.Context, user *model.UserM) error {
	_, err := u.mdb.Collection(consts.CollectionUsers).InsertOne(ctx, user)
	return err
}

func (u *User) GetUserByUsername(username string) (user *model.User, err error) {
	user = new(model.User)
	err = u.db.Model(user).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	return
}

func (u *User) GetUserInfo(ctx context.Context, id int64) (userInfo *model.UserM, err error) {
	userInfo = new(model.UserM)
	filter := bson.M{"user_id": id}
	err = u.mdb.Collection(consts.CollectionUsers).FindOne(ctx, filter).Decode(userInfo)
	return
}

func (u *User) ChangeAvatar(ctx context.Context, id int64, avatar string) error {
	filter := bson.M{"user_id": id}
	update := bson.M{"$set": bson.M{"avatar": avatar}}
	_, err := u.mdb.Collection(consts.CollectionUsers).UpdateOne(ctx, filter, update)
	return err
}

func (u *User) ChangeBackground(ctx context.Context, id int64, background string) error {
	filter := bson.M{"user_id": id}
	update := bson.M{"$set": bson.M{"background": background}}
	_, err := u.mdb.Collection(consts.CollectionUsers).UpdateOne(ctx, filter, update)
	return err
}

func (u *User) CacheUserInfo(ctx context.Context, userInfo *model.UserM) error {
	return u.rdb.SetEx(ctx, rdk.GetCacheUserInfoKey(userInfo.Id), userInfo, consts.CacheExpTime).Err()
}

func (u *User) ClearUserInfoCache(ctx context.Context, id int64) error {
	return u.rdb.Del(ctx, rdk.GetCacheUserInfoKey(id)).Err()
}
