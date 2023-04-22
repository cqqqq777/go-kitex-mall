package dao

import (
	"context"
	"github.com/cqqqq777/go-kitex-mall/cmd/cart/model"
	"github.com/cqqqq777/go-kitex-mall/cmd/cart/pkg"
	"github.com/cqqqq777/go-kitex-mall/shared/rdk"
	"github.com/redis/go-redis/v9"
)

type Cart struct {
	rdb *redis.Client
}

func NewCart(rdb *redis.Client) *Cart {
	return &Cart{
		rdb: rdb,
	}
}

func (c *Cart) AddProductToCart(ctx context.Context, userId, productId int64, v []byte) error {
	return c.rdb.HSet(ctx, rdk.GetCartKey(userId), rdk.GetCartProductKey(productId), v).Err()
}

func (c *Cart) GetCart(ctx context.Context, userId int64) ([]*model.CartProduct, error) {
	val, err := c.rdb.HGetAll(ctx, rdk.GetCartKey(userId)).Result()
	if err != nil {
		return nil, err
	}
	return pkg.UnMarshalProduct(val)
}

func (c *Cart) DelProduct(ctx context.Context, userId, productId int64) error {
	count, err := c.rdb.HDel(ctx, rdk.GetCartKey(userId), rdk.GetCartProductKey(productId)).Result()
	if err != nil {
		return err
	}
	if count != 1 {
		return pkg.ErrNoSuchProduct
	}
	return nil
}
