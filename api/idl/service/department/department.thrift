namespace go department

include "../../base/department.thrift"

struct AppendDepartmentRPCRequest {
    1: string department_name
    2: string department_description
    3: list<i64> city_list
}

struct AppendDepartmentRPCResponse {
    1: i32 code
    2: string message
    3: department.DepartmentInformation data
}

struct UpdateDepartmentRPCRequest {
    1: string department_name
    2: string department_description
    3: i64 department_id
}

struct UpdateDepartmentRPCResponse {
    1: i32 code
    2: string message
    3: department.DepartmentInformation data
}

struct DeleteDepartmentRPCRequest {
    1: i64 department_id
}

struct DeleteDepartmentRPCResponse {
    1: i32 code
    2: string message
}

struct GetsDepartmentRPCRequest {}

struct GetsDepartmentRPCResponse {
    1: i32 code
    2: string message
    3: list<department.DepartmentInformation> data
}

struct AppendCityRPCRequest {
    1: string city_name
}

struct AppendCityRPCResponse {
    1: i32 code
    2: string message
    3: department.CityInformation data
}

struct DeleteCityRPCRequest {
    1: i64 city_id
}

struct DeleteCityRPCResponse {
    1: i32 code
    2: string message
}

struct GetsCityRPCRequest {}

struct GetsCityRPCResponse {
    1: i32 code
    2: string message
    3: list<department.CityInformation> data
}

struct GetsCityByDepartmentRPCRequest {
    1: i64 department_id
}

struct GetsCityByDepartmentRPCResponse {
    1: i32 code
    2: string message
    3: list<department.CityInformation> data
}

struct GetDepartmentInfosByIdRPCRequest {
    1: list<i64> department_id_list
}

struct GetDepartmentInfosByIdRPCResponse {
    1: i32 code
    2: string message
    3: list<department.DepartmentInformation> department_info_list
}

struct GetCityInfoByIdRPCRequest {
    1: list<i64> city_id
}

struct GetCityInfoByIdRPCResponse {
    1: i32 code
    2: string message
    3: list<department.CityInformation> city_info_list
}

struct GetDepartmentCityInfoByIdRPCRequest {
    1: list<i64> department_id_list
    2: list<i64> city_id_list
}

struct GetDepartmentCityInfoByIdRPCResponse {
    1: i32 code
    2: string message
    3: list<department.DepartmentInformation> department_info_list
    4: list<department.CityInformation> city_info_list
}

service DepartmentRPCService {
    //与web层交互的接口
    //增加部门
    AppendDepartmentRPCResponse AppendDepartment(1: AppendDepartmentRPCRequest request)
    //修改部门信息
    UpdateDepartmentRPCResponse UpdateDepartment(1: UpdateDepartmentRPCRequest request)
    //删除部门
    DeleteDepartmentRPCResponse DeleteDepartment(1: DeleteDepartmentRPCRequest request)
    //查询所有部门信息
    GetsDepartmentRPCResponse GetsDepartment(1: GetsDepartmentRPCRequest request)
    //增加城市
    AppendCityRPCResponse AppendCity(1: AppendCityRPCRequest request)
    //删除城市
    DeleteCityRPCResponse DeleteCity(1: DeleteCityRPCRequest request)
    //查询所有城市
    GetsCityRPCResponse GetsCity(1: GetsCityRPCRequest request)
    //查询部门所在城市
    GetsCityByDepartmentRPCResponse GetsCityByDepartment(1: GetsCityByDepartmentRPCRequest request)

    //内部接口
    //根据部门id列表查部门信息
    GetDepartmentInfosByIdRPCResponse GetDepartmentInfosById(1: GetDepartmentInfosByIdRPCRequest request)
    //根据城市id列表查询城市信息
    GetCityInfoByIdRPCResponse GetCityInfoById(1: GetCityInfoByIdRPCRequest request)
    //根据部门id和城市id查询部门信息和城市信息
    GetDepartmentCityInfoByIdRPCResponse GetDepartmentCityInfoById(1: GetDepartmentCityInfoByIdRPCRequest request)
}