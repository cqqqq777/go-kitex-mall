namespace go api

include "common.thrift"

struct mall_verification_request{
    1: string email(api.query="email")
}

struct mall_verification_response{
    1: i64 code
    2: string msg
    3: i64 verification
}

struct mall_user_register_request{
    1: string username(api.query="username", api.vd="len($)>0 && len($)<33")
    2: string password(api.query="password",api.vd="len($)>0 && len($)<33")
    4:i64 verification(api.query="verification")
    3: string email(api.query="email")
}

struct mall_user_register_response {
    1: i64 code
    2: string msg
    3: i64 user_id
    4: string token
}

struct mall_user_login_request{
    1: string username(api.query="username", api.vd="len($)>0 && len($)<33")
    2: string password(api.query="password",api.vd="len($)>0 && len($)<33")
}

struct mall_user_login_response{
    1: i64 code
    2: string msg
    3: string token
}

service ApiService {
    mall_verification_response GetVerification(1: mall_verification_request req)(api.post="/api/verification")
    mall_user_register_response Register(1: mall_user_register_request req)(api.post="/api/register")
    mall_user_login_response Login(1: mall_user_login_request req)(api.get="/api/user/login")
}