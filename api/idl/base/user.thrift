namespace go base

include "department.thrift"

struct UserInfo {
    1: i64 user_id
    2: string username
    3: i64 department_id
    4: string department_name
}

struct CommonUserInfo {
    1: i64 user_id
    2: string username
}