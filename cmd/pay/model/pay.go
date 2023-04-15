package model

import "time"

type Pay struct {
	PayID      int64     `gorm:"primaKey"`
	UserID     int64     `gorm:"index;not null;column:user_id"`
	OrderID    int64     `gorm:"column:order_id"`
	Amount     int64     `gorm:"column:amount"`
	Status     int8      `gorm:"column:status"`
	Url        string    `gorm:"type:varchar(1024);column:url"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdateTime time.Time `gorm:"colum:update_time;autoCreateTime"`
}
