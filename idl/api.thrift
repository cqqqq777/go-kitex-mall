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
    2: string password(api.query="password", api.vd="len($)>0 && len($)<33")
}

struct mall_user_login_response{
    1: i64 code
    2: string msg
    3: string token
}

struct mall_get_user_info_request{
    1: i64 id(api.query="id")
}

struct mall_get_user_info_response{
    1: i64 code
    2: string msg
    3: common.User user_info
}

struct mall_change_user_avatar_request{
    1:string token(api.query="token")
}

struct mall_change_user_avatar_response{
    1: i64 code
    2: string msg
}

struct mall_change_user_background_request{
    1:string token(api.query="token")
}

struct mall_change_user_background_response{
    1: i64 code
    2: string msg
}

struct mall_merchant_register_request{
    1: string name(api.query="name", api.vd="len($)>0 && len($)<33")
    2: string password(api.query="password", api.vd="len($)>0 && len($)<33")
    3: i64 alipay(api.query="alipay")
    4: string description(api.query="description", api.vd="len($)>0 && len($)<100")
    5: i64 invitation_code(api.query="invitation_code")
}

struct mall_merchant_register_response {
    1: i64 code
    2: string msg
    3: string token
}

struct mall_merchant_login_request{
    1: string name(api.query="name", api.vd="len($)>0 && len($)<33")
    3: string password(api.query="password", api.vd="len($)>0 && len($)<33")
}

struct mall_merchant_login_response {
    1: i64 code
    2: string msg
    3: string token
}

struct mall_merchant_get_info_request{
    1: i64 id(api.query="id")
}

struct mall_merchant_get_info_response{
    1: i64 code
    2: string msg
    3: common.Merchant merchant
}

struct mall_publish_product_request{
    1: i64 merchant_id(api.query="merchant_id")
    2: string name(api.query="name", api.vd="len($)>0 && len($)<33")
    3: string description(api.query="description", api.vd="len($)>0 && len($)<100")
    4: i64 price(api.query="price")
    6: i64 stock(api.query="stock")
    7: string token(api.query="token")
}

struct mall_publish_product_response{
    1: i64 code
    2: string msg
    3: i64 product_id
}

struct mall_del_product_request {
    1: i64 merchant_id(api.query="merchant_id")
    2: i64 product_id(api.query="product_id")
    3: string token(api.query="token")
}

struct mall_del_product_response{
    1: i64 code
    2: string msg
}

struct mall_update_product_request{
    1: i64 product_id(api.query="product_id")
    2: string name(api.query="name", api.vd="len($)>0 && len($)<33")
    3: string description(api.query="description", api.vd="len($)>0 && len($)<100")
    4: i64 price(api.query="price")
    5: i64 stock(api.query="stock")
    6: string token(api.query="token")
}

struct mall_update_product_response{
    1: i64 code
    2: string msg
}

struct mall_product_list_request {
    1: i8 page(api.query="page")
    2: i8 page_size(api.query="page_size")
    3: string sort(api.query="sort")
}

struct mall_product_list_response{
    1: i64 code
    2: string msg
    3: list<common.Product> products
}

struct mall_product_detail_request{
    1: i64 product_id(api.query="product_id")
    2: string token(api.query="token")// optional
}

struct mall_product_detail_response{
    1: i64 code
    2: string msg
    3: common.ProductDetail product
}

struct mall_search_product_request{
    1: string key(api.query="key")
}

struct mall_search_product_response{
    1: i64 code
    2: string msg
    3: list<common.Product> products
}

struct mall_product_favorite_list_request{
    1: string token(api.query="token")
}

struct mall_product_favorite_list_response{
    1: i64 code
    2: string msg
    3: list<common.Product> products
}


struct mall_product_published_list_request{
    1: string token(api.query="token")
}

struct mall_product_published_list_response{
    1: i64 code
    2: string msg
    3: list<common.Product> products
}

struct mall_favorite_product_request {
   1: string token(api.query="token")
   2: i64 product_id(api.query="product_id")
}

struct mall_favorite_product_response {
   1: i64 code
   2: string msg
}

struct mall_create_order_request {
    1: string token(api.query="token")
    2: i64 product_id(api.query="product_id")
    3: i64 product_num(api.query="product_num")
    4: i64 amount(api.query="amount")
}

struct mall_create_order_response{
    1: i64 code
    2: string msg
    3: i64 order_id
    4: i64 amount
}

struct mall_order_list_requset{
    1: string token(api.query="token")
}

struct mall_order_list_response{
    1: i64 code
    2: string msg
    3: list<common.Order> orders
}

struct mall_get_order_request{
    1: i64 order_id(api.query="order_id")
    2: string token(api.query="token")
}

struct mall_get_order_response{
    1: i64 code
    2: string msg
    3: common.Order order
}

struct mall_create_pay_request{
    1: i64 order_id(api.query="order_id")
    2: string token(api.query="token")
    3: i64 amount(api.query="amount")
}

struct mall_create_pay_response{
    1: i64 code
    2: string msg
    3: i64 pay_id
    4: string url
}

struct mall_pay_detail_request {
    1: i64 pay_id(api.query="pay_id")
}

struct mall_pay_detail_response{
    1: i64 code
    2: string msg
    3: common.Pay pay
}

struct mall_pay_return_request{
    1: i64 pay_id(api.form="pay_id")
    2: i64 order_id(api.form="order_id")
    3: string token(api.form="token")
    4: i64 amount(api.form="amount")
    5: i8 status(api.form="status")
}

struct mall_pay_return_response{
    1: i64 code
    2: string msg
}

struct mall_pay_notify_request {
    1: string order_id(api.form="out_trade_no")
    2: string pay_id(api.form="subject")
    3: string status(api.form="trade_status")
}

struct mall_pay_notify_response{
    1: string msg
}

service ApiService {
    mall_verification_response GetVerification(1: mall_verification_request req)(api.post="/api/verification/")
    mall_user_register_response Register(1: mall_user_register_request req)(api.post="/api/register/")
    mall_user_login_response Login(1: mall_user_login_request req)(api.get="/api/user/login/")
    mall_get_user_info_response GetUserInfo(1: mall_get_user_info_request req)
    mall_change_user_avatar_response ChangeAvatar(1: mall_change_user_avatar_request req)
    mall_change_user_background_response ChangeBackground(1: mall_change_user_background_request req)

    mall_merchant_register_response MerchantRegister(1: mall_merchant_register_request req)(api.post="/api/merchant/register/")
    mall_merchant_login_response MerchantLogin(1: mall_merchant_login_request req)(api.get="/api/merchant/login/")
    mall_merchant_get_info_response MerchantInfo(1: mall_merchant_get_info_request req)(api.get="api/merchant/info/")

    mall_publish_product_response PublishProduct(1: mall_publish_product_request req)(api.post="/api/product/")
    mall_update_product_response UpdateProduct(1: mall_update_product_request req)(api.put="/api/product/")
    mall_del_product_response DelProduct(1: mall_del_product_request req)(api.delete="/api/products/")
    mall_product_list_response ProductList(1: mall_product_list_request req)(api.get="/api/products/list/")
    mall_product_detail_response ProductDetail(1: mall_product_detail_request req)(api.get="/api/products/")
    mall_search_product_response SearchProduct(1: mall_search_product_request req)(api.get="/api/products/search/")
    mall_product_favorite_list_response ProductFavoriteList(1: mall_product_favorite_list_request req)(api.get="/api/products/favorite")
    mall_product_published_list_response PublishedProducts(1: mall_product_published_list_request req)(api.get="/api/products/published")

    mall_create_order_response CreateOrder(1:mall_create_order_request req)(api.post="/api/order/")
    mall_order_list_response OrderList(1: mall_order_list_requset req)(api.get="/api/orders/list/")
    mall_get_order_response GetOrder(1: mall_get_order_request req)(api.get="/api/orders/")

    mall_create_pay_response CreatePay(1: mall_create_pay_request req)(api.post="/api/pay/")
    mall_pay_detail_response PayDetail(1: mall_pay_detail_request req)(api.get="/api/pay/")
    mall_pay_return_response PayReturn(1: mall_pay_return_request req)(api.post="/api/pay/return/")
    mall_pay_notify_response PayNotify(1:mall_pay_notify_request req)(api.post="/api/pay/notify/")
}