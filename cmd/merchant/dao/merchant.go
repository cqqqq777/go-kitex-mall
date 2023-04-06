package dao

import (
	"context"
	"errors"

	"github.com/cqqqq777/go-kitex-mall/cmd/merchant/model"
	"github.com/cqqqq777/go-kitex-mall/shared/rdk"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	ErrMerchantNotFound = errors.New("no such user")
	ErrMerchantExist    = errors.New("user already exist")
)

type Merchant struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewMerchant(db *gorm.DB, rdb *redis.Client) *Merchant {
	m := db.Migrator()
	if !m.HasTable(&model.Merchant{}) {
		err := m.CreateTable(&model.Merchant{})
		if err != nil {
			panic(err)
		}
	}
	return &Merchant{
		db:  db,
		rdb: rdb,
	}
}

func (m *Merchant) CreateMerchant(merchant *model.Merchant) error {
	err := m.db.Model(&model.Merchant{}).
		Where(&model.Merchant{Name: merchant.Name}).First(&model.Merchant{}).Error
	if err == nil {
		return ErrMerchantExist
	} else if err != gorm.ErrRecordNotFound {
		return err
	}
	return m.db.Model(&model.Merchant{}).Create(merchant).Error
}

func (m *Merchant) GetMerchantByName(name string) (merchant *model.Merchant, err error) {
	merchant = new(model.Merchant)
	err = m.db.Model(merchant).
		Where(&model.Merchant{Name: name}).First(merchant).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrMerchantNotFound
	}
	return
}

func (m *Merchant) GetMerchantById(id int64) (merchant *model.Merchant, err error) {
	merchant = new(model.Merchant)
	err = m.db.Model(merchant).
		Where(&model.Merchant{Id: id}).First(merchant).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrMerchantNotFound
	}
	return
}

func (m *Merchant) CacheMerchantInfo(ctx context.Context, merchant *model.Merchant) error {
	return m.rdb.Set(ctx, rdk.GetCacheMerchantInfoKey(merchant.Id), merchant, 0).Err()
}
