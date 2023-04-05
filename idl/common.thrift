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