package dao

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	rdb *redis.Client
	mdb *mongo.Database
}

func NewProduct(rdb *redis.Client, mdb *mongo.Database) *Product {
	return &Product{
		rdb: rdb,
		mdb: mdb,
	}
}
