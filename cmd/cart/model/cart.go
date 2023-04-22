package model

type CartProduct struct {
	ProductId  int64 `json:"product_id"`
	MerchantId int64 `json:"merchant_id"`
	ProductNum int64 `json:"product_num"`
	Amount     int64 `json:"amount"`
	AddTime    int64 `json:"add_time"`
}
