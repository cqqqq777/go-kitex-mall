package dao

import (
	"context"
	"errors"
	"time"

	"github.com/cqqqq777/go-kitex-mall/cmd/product/model"
	"github.com/cqqqq777/go-kitex-mall/shared/consts"
	"github.com/cqqqq777/go-kitex-mall/shared/rdk"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrNoPermission = errors.New("no permission")
	ErrNoProduct    = errors.New("no such product")
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

func (p *Product) CreateProduct(ctx context.Context, product *model.ProductBasic) error {
	// create index
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"name":        "text",
			"description": "text",
		},
	}
	_, err := p.mdb.Collection(consts.CollectionProducts).Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}
	_, err = p.mdb.Collection(consts.CollectionProducts).InsertOne(ctx, product)
	return err
}

func (p *Product) DelProduct(ctx context.Context, mid, pid int64) error {
	product := new(model.ProductBasic)
	err := p.mdb.Collection(consts.CollectionProducts).FindOne(ctx, bson.M{"product_id": pid}).Decode(product)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrNoProduct
		}
		return err
	}
	if product.Mid != mid {
		return ErrNoPermission
	}
	_, err = p.mdb.Collection(consts.CollectionProducts).DeleteOne(ctx, bson.M{"product_id": pid})
	return err
}

func (p *Product) ProductList(ctx context.Context, key string, page, pageSize int64) (list []*model.ProductBasic, total int64, err error) {
	list = make([]*model.ProductBasic, 0, 20)
	// get list
	opts := options.Find().SetLimit(pageSize).SetSkip((page - 1) * pageSize).SetSort(bson.M{key: -1})
	cursor, err := p.mdb.Collection(consts.CollectionProducts).Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	err = cursor.All(ctx, &list)

	// get num
	total, err = p.mdb.Collection(consts.CollectionProducts).CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}
	return
}

func (p *Product) GetProductInfo(ctx context.Context, pid int64) (info *model.ProductBasic, err error) {
	info = new(model.ProductBasic)
	err = p.mdb.Collection(consts.CollectionProducts).FindOne(ctx, bson.M{"product_id": pid}).Decode(info)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, ErrNoProduct
	}
	return
}

func (p *Product) SearchProduct(ctx context.Context, key string) (list []*model.ProductBasic, err error) {
	list = make([]*model.ProductBasic, 0, 20)

	filter := bson.M{"$text": bson.M{"$search": key}}
	cursor, err := p.mdb.Collection(consts.CollectionProducts).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &list)
	return
}

func (p *Product) CacheProductInfo(ctx context.Context, id int64, data interface{}) error {
	return p.rdb.SetEx(ctx, rdk.GetCacheProductDetailKey(id), data, consts.CacheExpTime).Err()
}

func (p *Product) ClearProductCache(ctx context.Context, id int64) error {
	return p.rdb.Del(ctx, rdk.GetCacheProductDetailKey(id)).Err()
}

func (p *Product) PublishedProducts(ctx context.Context, mid int64) (list []*model.ProductBasic, err error) {
	list = make([]*model.ProductBasic, 0, 20)
	cursor, err := p.mdb.Collection(consts.CollectionProducts).Find(ctx, bson.M{"merchant_id": mid})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &list)
	return
}

func (p *Product) CacheNullProductInfo(ctx context.Context, pid int64) {
	p.rdb.SetEx(ctx, rdk.GetCacheProductDetailKey(pid), "", time.Second*120)
}

func (p *Product) GetFavoriteId(ctx context.Context, uid int64) ([]string, error) {
	var cursor uint64
	result, _, err := p.rdb.SScan(ctx, rdk.GetUserFavoriteProductKey(uid), cursor, "", 0).Result()
	return result, err
}

func (p *Product) GetFavorite(ctx context.Context, ids []string) (list []*model.ProductBasic, err error) {
	list = make([]*model.ProductBasic, 0, 20)
	filter := bson.M{"product_id": bson.M{"$in": ids}}
	cur, err := p.mdb.Collection(consts.CollectionProducts).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	err = cur.All(ctx, &list)
	return
}

func (p *Product) UpdateProduct(ctx context.Context, updateInfo *model.UpdateInfo) error {
	filter := bson.M{"product_id": updateInfo.Id}
	update := bson.M{"$set": bson.M{"name": updateInfo.Name, "price": updateInfo.Price, "stock": updateInfo.Stock, "description": updateInfo.Description}}
	_, err := p.mdb.Collection(consts.CollectionProducts).UpdateOne(ctx, filter, update)
	return err
}

func (p *Product) UpdateStock(ctx context.Context, id, stock int64) error {
	filter := bson.M{"product_id": id}
	update := bson.M{"$set": bson.M{"stock": stock}}
	_, err := p.mdb.Collection(consts.CollectionProducts).UpdateOne(ctx, filter, update)
	return err
}
