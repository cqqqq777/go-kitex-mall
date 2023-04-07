namespace go operate

include "common.thrift"

struct mall_favorite_product_request {
   1: i64 user_id
   2: i64 product_id
}

struct mall_favorite_product_response {
   1: common.common_response common_resp
}

struct mall_get_product_operate_info_request{
    1: i64 user_id
    2: i64 product_id
}

struct mall_get_product_operate_info_response{
    1: common.common_response common_resp
    2: common.ProductOperateInfo operate_info
}

service OperateService{
    mall_favorite_product_response FavoriteProduct(1: mall_favorite_product_request req)
    mall_get_product_operate_info_response GetProductOperateInfo(1: mall_get_product_operate_info_request req)
}