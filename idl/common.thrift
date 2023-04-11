namespace go common

struct common_response{
    1: i64 code
    2: string msg
}

struct Comment{
    1:i64 id
    2:User user //comment user info
    3:string content
    4:string create_data // format mm-dd
}

struct User {
    1: i64 id
    2:string name
    3:string avatar
    4:string background
    6:string signature
}

struct Merchant {
    1: i64 id
    2: i64 alipay
    3: string name
    4: string description
}

struct Image{
    1: i16 id
    2: string path
}

struct Product {
    1: i64 id
    2: i64 m_id
    3: i64 price
    4: string name
    5: string description
    6: list<Image> iamges
    7: i64 stock
    8: i8 status
}

struct ProductOperateInfo {
    1: i64 sale_count
    2: i64 comment_count
    3: bool is_favorite
}

struct ProductDetail {
    1: Product basic_info
    2: Merchant merchant_info
    3: ProductOperateInfo operate_info
    4: i64 create_time
    5: i64 update_time
}

struct Order {
    1: i64 order_id
    2: i64 user_id
    3: i64 product_id
    4: i64 product_num
    5: i64 amount
    6: i8 status
    7: i64 create_time
    8: i64 update_time
    9: i64 exp_time
}