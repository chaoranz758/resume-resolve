namespace go department

include "../../base/common.thrift"

struct AppendDepartmentRequest {
    1: string department_name (api.body = "department_name")
    2: string department_description (api.body = "department_description")
    3: list<i64> city_list (api.body = "city_list")
}

struct UpdateDepartmentRequest {
    1: string department_name (api.body = "department_name")
    2: string department_description (api.body = "department_description")
    3: i64 department_id (api.body = "department_id")
}

struct DeleteDepartmentRequest {
    1: i64 department_id (api.body = "department_id")
}

struct GetsDepartmentRequest {}

struct AppendCityRequest {
    1: string city_name (api.body = "city_name")
}

struct DeleteCityRequest {
    1: i64 city_id (api.body = "city_id")
}

struct GetsCityRequest {}

struct GetsCityByDepartmentRequest {
    1: i64 department_id (api.query = "department_id")
}

service DepartmentService {
    //增加部门
    common.NilResponse AppendDepartment(1: AppendDepartmentRequest request) (api.post = "/api/v1/department/append")
    //修改部门信息
    common.NilResponse UpdateDepartment(1: UpdateDepartmentRequest request) (api.post = "/api/v1/department/update")
    //删除部门
    common.NilResponse DeleteDepartment(1: DeleteDepartmentRequest request) (api.post = "/api/v1/department/delete")
    //查询所有部门信息
    common.NilResponse GetsDepartment(1: GetsDepartmentRequest request) (api.get = "/api/v1/department/gets")
    //增加城市
    common.NilResponse AppendCity(1: AppendCityRequest request) (api.post = "/api/v1/department/city/append")
    //删除城市
    common.NilResponse DeleteCity(1: DeleteCityRequest request) (api.post = "/api/v1/department/city/delete")
    //查询所有城市
    common.NilResponse GetsCity(1: GetsCityRequest request) (api.get = "/api/v1/department/city/gets")
    //查询部门所在城市
    common.NilResponse GetsCityByDepartment(1: GetsCityByDepartmentRequest request) (api.get = "/api/v1/department/city/gets-department")
}