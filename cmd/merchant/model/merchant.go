package model

type Merchant struct {
	Id          int64  `gorm:"primarykey"`
	Alipay      int64  `gorm:"type:bigint;not null"`
	Name        string `gorm:"type:varchar(33);unique;not null"`
	Password    string `gorm:"type:varchar(33);not null"`
	Description string `gorm:"type:varchar(1024);not null"`
}
