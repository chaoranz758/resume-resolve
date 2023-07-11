namespace go user

include "../../base/common.thrift"

struct UserLoginRequest {
    1: string username (api.body = "username")
    2: string password (api.body = "password")
    3: i8 role (api.body = "role")
}

struct UserChangePasswordRequest {
    1: string username (api.body = "username")
    2: string password (api.body = "password")
    3: string new_password (api.body = "new_password")
    4: i8 role (api.body = "role")
}

struct UserRegisterRequest {
    1: string username (api.body = "username")
    2: string password (api.body = "password")
    3: string re_password (api.body = "re_password")
}

struct HRRegisterRequest {
    1: string username (api.body = "username")
    2: i64 department_id (api.body = "department_id")
}

struct HRDeleteRequest {
    1: i64 user_id (api.body = "user_id")
}

struct GetHRByDepartmentRequest {
    1: i64 department_id (api.query = "department_id") //有id按部门查，没有id查全部
    2: i32 limit (api.query = "limit")
    3: i64 offset (api.query = "offset")
}

struct GetUserByNameRequest {
    1: string username (api.query = "username")
}

struct GetHRByNameRequest {
    1: string username (api.query = "username")
}

service UserService {
    //登录
     common.NilResponse UserLogin(1: UserLoginRequest request) (api.post = "/api/v1/user/login")
    //修改管理员或HR的密码
    common.NilResponse UserChangePassword(1: UserChangePasswordRequest request) (api.post = "/api/v1/user/change-password")
    //普通用户注册
    common.NilResponse UserRegister(1: UserRegisterRequest request) (api.post = "/api/v1/user/register")
    //新增HR
    common.NilResponse HRRegister(1: HRRegisterRequest request) (api.post = "/api/v1/user/hr-register")
    //删除HR
    common.NilResponse HRDelete(1: HRDeleteRequest request) (api.post = "/api/v1/user/hr-delete")
    //查询全部HR信息或按部门查询HR信息
    common.NilResponse GetHRByDepartment(1: GetHRByDepartmentRequest request) (api.get = "/api/v1/user/gets-department")
    //根据HR用户名查询HR信息
    common.NilResponse GetHRByName(1: GetHRByNameRequest request) (api.get = "/api/v1/user/get-hr")
    //根据用户名查询用户信息
    common.NilResponse GetUserByName(1: GetUserByNameRequest request) (api.get = "/api/v1/user/get")
}

