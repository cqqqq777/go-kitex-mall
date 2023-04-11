namespace go product

include "common.thrift"

struct mall_publish_product_request {
    1: i64 merchant_id
    2: string name
    3: string description
    4: i64 price
    5: list<common.Image> images
    6: i64 stock
}

struct mall_publish_product_response {
    1: common.common_response common_resp
    2: i64 product_id
}

struct mall_del_product_request{
    1: i64 merchant_id
    2: i64 product_id
}

struct mall_update_product_request{
    1: i64 product_id
    2: string name
    3: string description
    4: i64 price
    5: i64 stock
}

struct mall_update_product_response{
    1: common.common_response common_resp
}

struct mall_del_product_response {
    1: common.common_response common_resp
}


struct mall_product_list_request {
    1: i8 page
    2: i8 page_size
    3: string sort
}

struct mall_product_list_response {
    1: common.common_response common_resp
    2: i64 total_num
    3: list<common.Product> products
}

struct mall_product_detail_request{
    1: i64 user_id // Optional parameter
    2: i64 product_id
}

struct mall_product_detail_response {
    1: common.common_response common_resp
    2: common.ProductDetail product
}

struct mall_search_product_request {
    1: string key
}

struct mall_search_product_response{
    1: common.common_response common_resp
    2: list<common.Product> products
}

struct mall_product_favorite_list_request{
   1: i64 user_id
}

struct mall_product_favorite_list_response{
    1:common.common_response common_resp
    2: list<common.Product> products
}

struct mall_product_published_list_request {
    1: i64 merchant_id
}

struct mall_product_published_list_response{
    1: common.common_response common_resp
    2: list<common.Product> products
}

struct mall_update_stock_request{
    1: i64 product_id
    2: i64 stock
}

struct mall_update_stock_response{
    1: common.common_response common_resp
}

service ProductService {
    mall_publish_product_response PublishProduct(1: mall_publish_product_request req)
    mall_update_product_response UpdateProduct(1: mall_update_product_request req)
    mall_del_product_response DelProduct(1: mall_del_product_request req)
    mall_product_list_response ProductList(1: mall_product_list_request req)
    mall_product_detail_response ProductDetail(1: mall_product_detail_request req)
    mall_search_product_response SearchProduct(1: mall_search_product_request req)
    mall_product_favorite_list_response ProductFavoriteList(1: mall_product_favorite_list_request req)
    mall_product_published_list_response PublishedProducts(1: mall_product_published_list_request req)
    mall_update_stock_response UpdateStock(1: mall_update_stock_request req)
}