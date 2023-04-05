namespace go user

include "common.thrift"

struct mall_verification_request {
    1: string email
}

struct mall_verification_response{
    1: common.common_response comon_resp
    2: i64 verification
}

struct mall_user_register_request{
    1: string username
    2: string password
    3: string email
    4: i64 verification //verification code
}

struct mall_user_register_response{
    1: common.common_response common_resp
    2: i64 user_id
    3: string token
}

struct mall_user_login_request {
    1:string username
    2:string password
}

struct mall_user_login_response{
    1:common.common_response common_resp
    2:i64 user_id
    3:string token
}

struct mall_get_user_info_request{
    1:i64 id //user_id
}

struct mall_get_user_info_response{
    1: common.common_response common_resp
    2: common.User user_info
}

struct mall_change_user_avatar_request{
    1:i64 id
    2:string avatar
}

struct mall_change_user_avatar_response{
    1:common.common_response common_resp
}

struct mall_change_user_background_request{
    1:i64 id
    2:string background
}

struct mall_change_user_background_response{
    1:common.common_response common_resp
}

struct mall_change_user_info_request{
    1: string signature
}

struct mall_change_user_info_response{
    1:common.common_response common_resp
}

service UserService {
    mall_verification_response GetVerification(1: mall_verification_request req)
    mall_user_register_response Register(1: mall_user_register_request req)
    mall_user_login_response Login(1: mall_user_login_request req)
    mall_get_user_info_response GetUserInfo(1: mall_get_user_info_request req)
    mall_change_user_avatar_response ChangeAvatar(1: mall_change_user_avatar_request req)
    mall_change_user_background_response ChangeBackground(1: mall_change_user_background_request req)
}