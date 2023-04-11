package model

import "time"

type Order struct {
	Id         int64     `gorm:"primaryKey"`
	UserId     int64     `gorm:"index;not null;column:user_id"`
	ProductId  int64     `gorm:"column:product_id"`
	ProductNum int64     `gorm:"column:product_num"`
	Amount     int64     `gorm:"not null"`
	ExpTime    int64     `gorm:"column:exp_time;not null"`
	Status     int8      `gorm:"not null"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdateTime time.Time `gorm:"colum:update_time;autoCreateTime"`
}
