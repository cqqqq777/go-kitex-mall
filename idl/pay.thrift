namespace go pay

include "common.thrift"

struct mall_create_pay_request{
    1: i64 order_id
    2: i64 user_id
    3: i64 amount
}

struct mall_create_pay_response{
    1: common.common_response common_resp
    2: i64 pay_id
    3: string url
}

struct mall_pay_detail_request {
    1: i64 pay_id
}

struct mall_pay_detail_response{
    1: common.common_response common_resp
    2: common.Pay pay
}

struct mall_pay_return_request{
    1: i64 pay_id
    2: i64 order_id
    3: i64 user_id
    4: i64 amount
    5: i8 status
}

struct mall_pay_return_response{
    1: common.common_response common_resp
}

service PayService {
    mall_create_pay_response CreatePay(1: mall_create_pay_request req)
    mall_pay_detail_response PayDetail(1: mall_pay_detail_request req)
    mall_pay_return_response PayReturn(1: mall_pay_return_request req)
}