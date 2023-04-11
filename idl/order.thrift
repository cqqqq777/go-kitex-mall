namespace go order

include "common.thrift"

struct mall_create_order_request {
    1: i64 user_id
    2: i64 product_id
    3: i64 product_num
    4: i64 amount
}

struct mall_create_order_response{
    1: common.common_response common_resp
    2: i64 order_id
}

struct mall_update_order_request {
    1: i64 order_id
    2: i8 status
}

struct mall_update_order_response {
    1: common.common_response common_resp
}

struct mall_order_list_requset{
    1: i64 user_id
}

struct mall_order_list_response{
    1: common.common_response common_resp
    2: list<common.Order> orders
}

struct mall_get_order_request{
    1: i64 order_id
    2: i64 user_id
}

struct mall_get_order_response{
    1: common.common_response common_resp
    2: common.Order order
}

service OrderService {
    mall_create_order_response CreateOrder(1:mall_create_order_request req)
    mall_update_order_response UpdateOrder(1: mall_update_order_request req)
    mall_order_list_response OrderList(1: mall_order_list_requset req)
    mall_get_order_response GetOrder(1: mall_get_order_request req)
}