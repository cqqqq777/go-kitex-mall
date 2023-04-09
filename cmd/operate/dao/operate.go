package dao

import (
	"context"

	"github.com/cqqqq777/go-kitex-mall/shared/rdk"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Operate struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewOperate(db *gorm.DB, rdb *redis.Client) *Operate {
	return &Operate{
		db:  db,
		rdb: rdb,
	}
}

func (o *Operate) GetUserFavoriteStatus(ctx context.Context, uid, pid int64) (bool, error) {
	cmd := o.rdb.SIsMember(ctx, rdk.GetUserFavoriteProductKey(uid), pid)
	return cmd.Val(), cmd.Err()
}

func (o *Operate) FavoriteProduct(ctx context.Context, uid, pid int64) error {
	return o.rdb.SAdd(ctx, rdk.GetUserFavoriteProductKey(uid), pid).Err()
}

func (o *Operate) CancelFavorite(ctx context.Context, uid, pid int64) error {
	return o.rdb.SRem(ctx, rdk.GetUserFavoriteProductKey(uid), pid).Err()
}

func (o *Operate) GetCommentNum(pid int64) (int64, error) {
	var num int64
	err := o.db.Raw("select count(1) from comments where product_id = ?", pid).Scan(&num).Error
	return num, err
}

func (o *Operate) GetSaleNum(ctx context.Context, pid int64) (int64, error) {
	cmd := o.rdb.Get(ctx, rdk.GetProductSale(pid))
	return cmd.Int64()
}
