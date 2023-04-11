package dao

import (
	"context"
	"fmt"
	"github.com/cqqqq777/go-kitex-mall/cmd/order/model"
	"github.com/cqqqq777/go-kitex-mall/shared/rdk"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Order struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewOrder(db *gorm.DB, rdb *redis.Client) *Order {
	m := db.Migrator()
	if !m.HasTable(&model.Order{}) {
		err := m.CreateTable(&model.Order{})
		if err != nil {
			panic(err)
		}
	}
	return &Order{
		db:  db,
		rdb: rdb,
	}
}

func (o *Order) CreateOrder(order *model.Order) error {
	return o.db.Model(order).Create(order).Error
}

func (o *Order) SetOrderInRedis(ctx context.Context, id, amount, expTime int64) error {
	val := fmt.Sprintf("%d+%d+%d", id, amount, expTime)
	return o.rdb.Set(ctx, rdk.GetSetOrderKey(id), val, 0).Err()
}

func (o *Order) UpdateOrder(id int64, status int8) error {
	return o.db.Model(&model.Order{}).Where("id = ?", id).Update("status", status).Error
}

func (o *Order) GetOrderList(uid int64) (list []*model.Order, err error) {
	list = make([]*model.Order, 0, 20)
	err = o.db.Raw("select id,user_id,product_id,product_num,amount,exp_time,status,create_time,update_time from orders where user_id = ?", uid).Scan(&list).Error
	return
}

func (o *Order) GetOrder(id int64) (order *model.Order, err error) {
	order = new(model.Order)
	err = o.db.Raw("select id,user_id,product_id,product_num,amount,exp_time,status,create_time,update_time from orders where id = ?", id).Scan(order).Error
	return
}
