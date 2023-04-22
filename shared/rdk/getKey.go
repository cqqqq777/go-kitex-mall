package rdk

import "strconv"

func GetVerificationKey(email string) string {
	return Prefix + email + RKVerification
}

func GetCacheUserInfoKey(id int64) string {
	return Prefix + strconv.FormatInt(id, 10) + RKCacheUserInfo
}

func GetCacheMerchantInfoKey(id int64) string {
	return Prefix + strconv.FormatInt(id, 10) + RKCacheMerchantInfo
}

func GetCacheProductDetailKey(id int64) string {
	return Prefix + strconv.FormatInt(id, 10) + RKCacheProductDetail
}

func GetUserFavoriteProductKey(id int64) string {
	return Prefix + strconv.FormatInt(id, 10) + RKFavoriteProduct
}

func GetProductSale(id int64) string {
	return Prefix + strconv.FormatInt(id, 10) + RKProductSaleNum
}

func GetSetOrderKey(id int64) string {
	return Prefix + strconv.FormatInt(id, 10) + RKSetOrder
}

func GetCartKey(id int64) string {
	return Prefix + strconv.FormatInt(id, 10) + RKCart
}

func GetCartProductKey(id int64) string {
	return Prefix + strconv.FormatInt(id, 10) + RKCartProduct
}
