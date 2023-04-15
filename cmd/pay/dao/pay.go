package dao

import (
	"github.com/cqqqq777/go-kitex-mall/cmd/pay/model"
	"gorm.io/gorm"
)

type Pay struct {
	db *gorm.DB
}

func NewPayDao(db *gorm.DB) *Pay {
	m := db.Migrator()
	if !m.HasTable(&model.Pay{}) {
		err := m.CreateTable(&model.Pay{})
		if err != nil {
			panic(err)
		}
	}
	return &Pay{
		db: db,
	}
}

func (p *Pay) CreatePay(payInfo *model.Pay) error {
	return p.db.Create(payInfo).Error
}

func (p *Pay) GetPay(id int64) (payInfo *model.Pay, err error) {
	payInfo = new(model.Pay)
	err = p.db.Model(payInfo).Where("pay_id = ?", id).Find(payInfo).Error
	return
}

func (p *Pay) UpdatePayStatus(id int64, status int8) error {
	return p.db.Model(&model.Pay{}).Where("pay_id=?", id).Update("status", status).Error
}
