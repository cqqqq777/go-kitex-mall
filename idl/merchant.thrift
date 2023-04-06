namespace go merchant

include "common.thrift"

struct mall_merchant_register_request {
    1: string name
    2: string password
    3: i64 alipay
    4: string description
    5: i64 invitation_code
}

struct mall_merchant_register_response {
    1: common.common_response common_resp
    2: i64 id
    3: string token
}

struct mall_merchant_login_request{
    1: string name
    2: string password
}

struct mall_merchant_login_response {
    1: common.common_response common_resp
    2: i64 id
    3: string token
}

struct mall_merchant_get_info_request {
    1: i64 id
}

struct mall_merchant_get_info_response {
    1: common.common_response common_resp
    2: common.Merchant merchant_info
}

service MerchantService{
    mall_merchant_register_response Register(1: mall_merchant_register_request req)
    mall_merchant_login_response Login(1: mall_merchant_login_request req)
    mall_merchant_get_info_response GetInfo(1: mall_merchant_get_info_request req)
}