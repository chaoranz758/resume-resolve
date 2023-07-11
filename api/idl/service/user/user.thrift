namespace go user

include "../../base/common.thrift"
include "../../base/user.thrift"

struct UserChangePasswordRPCRequest {
    1: string username
    2: string password
    3: string new_password
    4: i8 role
}

struct UserChangePasswordRPCResponse {
    1: i32 code
    2: string message
}

struct UserLoginRPCRequest {
    1: string username
    2: string password
    3: i8 role
}

struct UserLoginRPCResponse {
    1: i32 code
    2: string message
}

struct UserRegisterRPCRequest {
    1: string username
    2: string password
    3: string re_password
}

struct UserRegisterRPCResponse {
    1: i32 code
    2: string message
}

struct HRRegisterRPCRequest {
    1: string username
    2: i64 department_id
}

struct HRRegisterRPCResponse {
    1: i32 code
    2: string message
}

struct HRDeleteRPCRequest {
    1: i64 user_id
}

struct HRDeleteRPCResponse {
    1: i32 code
    2: string message
}

struct GetHRByDepartmentRPCRequest {
    1: i64 department_id //有id按部门查，没有id查全部
    2: i32 limit
    3: i32 offset
}

struct GetHRByDepartmentRPCResponse {
    1: i32 code
    2: string message
    3: list<user.UserInfo> data
}

struct GetUserByNameRPCRequest {
    1: string username
}

struct GetUserByNameRPCResponse {
    1: i32 code
    2: string message
    3: user.CommonUserInfo data
}

struct GetHRByNameRPCRequest {
    1: string username
}

struct GetHRByNameRPCResponse {
    1: i32 code
    2: string message
    3: user.UserInfo data
}

struct DeleteHRByDepartmentIdRPCRequest {
    1: i64 department_id
}

struct DeleteHRByDepartmentIdRPCResponse {
    1: i32 code
    2: string message
}

service UserRPCService {
    //与web层交互
    //修改管理员或HR的密码
    UserChangePasswordRPCResponse UserChangePassword(1: UserChangePasswordRPCRequest request)
    //管理员或HR登录
    UserLoginRPCResponse UserLogin(1: UserLoginRPCRequest request)
    //普通用户注册
    UserRegisterRPCResponse UserRegister(1: UserRegisterRPCRequest request)
    //新增HR
    HRRegisterRPCResponse HRRegister(1: HRRegisterRPCRequest request)
    //删除HR
    HRDeleteRPCResponse HRDelete(1: HRDeleteRPCRequest request)
    //查询全部HR信息或按部门查询HR信息
    GetHRByDepartmentRPCResponse GetHRByDepartment(1: GetHRByDepartmentRPCRequest request)
    //根据普通用户名查询普通用户信息
    GetUserByNameRPCResponse GetUserByName(1: GetUserByNameRPCRequest request)
    //根据HR用户名查询HR信息
    GetHRByNameRPCResponse GetHRByName(1: GetHRByNameRPCRequest request)

    //内部接口
    //根据部门id删除hr
    DeleteHRByDepartmentIdRPCResponse DeleteHRByDepartmentId(1: DeleteHRByDepartmentIdRPCRequest request)
}