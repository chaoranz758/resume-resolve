package service

var (
	errCreateDepartment          = "create department or create department city map relative information failed"
	errUpdateDepartment          = "update department information failed"
	errDeleteDepartmentFromRedis = "delete department information from redis failed"
	errDeleteDepartment          = "delete department and department city map relative information failed"
	errGetsDepartment            = "gets department information failed"
	errConcurrent                = "concurrent errors"
	errJsonUnMarshal             = "json unmarshal failed"
	errCreateCity                = "create city information failed"
	errDeleteCityFromRedis       = "delete city information from redis failed"
	errDeleteCity                = "delete city information failed"
	errGetsCity                  = "gets city information failed"
	errGetsCityByDepartment      = "gets city by department failed"
	errGetDepartmentInfosById    = "get department information by id failed"
	errGetCityInfoById           = "get city information by id failed"
	errRpcService                = "rpc service error"
	errRpcBizService             = "rpc service biz error"
)
