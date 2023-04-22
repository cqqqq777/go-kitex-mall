namespace go cart


include "common.thrift"

struct mall_add_product_to_cart_request {
    1: i64 user_id
    2: i64 product_id
    3: i64 product_num
}

struct mall_add_product_to_cart_response{
    1: common.common_response common_resp
}

struct mall_get_cart_request{
    1: i64 user_id
}

struct mall_get_cart_response{
    1: common.common_response common_resp
    2: common.cart cart
}

struct mall_del_cart_product_request{
    1: i64 user_id
    2: i64 product_id
}

struct mall_del_cart_product_response{
    1: common.common_response common_resp
}

service CartService{
    mall_add_product_to_cart_response AddProductToCart(1: mall_add_product_to_cart_request req)
    mall_get_cart_response GetCart(1: mall_get_cart_request req)
    mall_del_cart_product_response DelCartProduct(1: mall_del_cart_product_request req)
}