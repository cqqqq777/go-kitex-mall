package model

type User struct {
	Id       int64  `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(33);unique;not null"`
	Password string `gorm:"type:varchar(33);not null"`
	Email    string `gorm:"type:varchar(33);unique;not null"`
}

type UserM struct {
	Id         int64  `bson:"_id,omitempty"`
	Username   string `bson:"username,omitempty"`
	Avatar     string `bson:"avatar,omitempty"`
	Background string `bson:"background,omitempty"`
	Signature  string `bson:"signature,omitempty"`
}
