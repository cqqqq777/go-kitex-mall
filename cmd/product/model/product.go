package model

type ProductBasic struct {
	Id          int64    `bson:"product_id"`
	Mid         int64    `bson:"merchant_id"`
	Price       int64    `bson:"price"`
	Stock       int64    `bson:"stock"`
	Name        string   `bson:"name"`
	Description string   `bson:"description"`
	Status      int8     `bson:"status"`
	Images      []*Image `bson:"images"`
}

type Image struct {
	Id   int16  `bson:"image_id"`
	Path string `bson:"path"`
}
