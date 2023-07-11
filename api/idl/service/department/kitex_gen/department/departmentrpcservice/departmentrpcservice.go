// Code generated by Kitex v0.5.2. DO NOT EDIT.

package departmentrpcservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	department "resume-resolving/api/idl/service/department/kitex_gen/department"
)

func serviceInfo() *kitex.ServiceInfo {
	return departmentRPCServiceServiceInfo
}

var departmentRPCServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "DepartmentRPCService"
	handlerType := (*department.DepartmentRPCService)(nil)
	methods := map[string]kitex.MethodInfo{
		"AppendDepartment":          kitex.NewMethodInfo(appendDepartmentHandler, newDepartmentRPCServiceAppendDepartmentArgs, newDepartmentRPCServiceAppendDepartmentResult, false),
		"UpdateDepartment":          kitex.NewMethodInfo(updateDepartmentHandler, newDepartmentRPCServiceUpdateDepartmentArgs, newDepartmentRPCServiceUpdateDepartmentResult, false),
		"DeleteDepartment":          kitex.NewMethodInfo(deleteDepartmentHandler, newDepartmentRPCServiceDeleteDepartmentArgs, newDepartmentRPCServiceDeleteDepartmentResult, false),
		"GetsDepartment":            kitex.NewMethodInfo(getsDepartmentHandler, newDepartmentRPCServiceGetsDepartmentArgs, newDepartmentRPCServiceGetsDepartmentResult, false),
		"AppendCity":                kitex.NewMethodInfo(appendCityHandler, newDepartmentRPCServiceAppendCityArgs, newDepartmentRPCServiceAppendCityResult, false),
		"DeleteCity":                kitex.NewMethodInfo(deleteCityHandler, newDepartmentRPCServiceDeleteCityArgs, newDepartmentRPCServiceDeleteCityResult, false),
		"GetsCity":                  kitex.NewMethodInfo(getsCityHandler, newDepartmentRPCServiceGetsCityArgs, newDepartmentRPCServiceGetsCityResult, false),
		"GetsCityByDepartment":      kitex.NewMethodInfo(getsCityByDepartmentHandler, newDepartmentRPCServiceGetsCityByDepartmentArgs, newDepartmentRPCServiceGetsCityByDepartmentResult, false),
		"GetDepartmentInfosById":    kitex.NewMethodInfo(getDepartmentInfosByIdHandler, newDepartmentRPCServiceGetDepartmentInfosByIdArgs, newDepartmentRPCServiceGetDepartmentInfosByIdResult, false),
		"GetCityInfoById":           kitex.NewMethodInfo(getCityInfoByIdHandler, newDepartmentRPCServiceGetCityInfoByIdArgs, newDepartmentRPCServiceGetCityInfoByIdResult, false),
		"GetDepartmentCityInfoById": kitex.NewMethodInfo(getDepartmentCityInfoByIdHandler, newDepartmentRPCServiceGetDepartmentCityInfoByIdArgs, newDepartmentRPCServiceGetDepartmentCityInfoByIdResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "department",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func appendDepartmentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*department.DepartmentRPCServiceAppendDepartmentArgs)
	realResult := result.(*department.DepartmentRPCServiceAppendDepartmentResult)
	success, err := handler.(department.DepartmentRPCService).AppendDepartment(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDepartmentRPCServiceAppendDepartmentArgs() interface{} {
	return department.NewDepartmentRPCServiceAppendDepartmentArgs()
}

func newDepartmentRPCServiceAppendDepartmentResult() interface{} {
	return department.NewDepartmentRPCServiceAppendDepartmentResult()
}

func updateDepartmentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*department.DepartmentRPCServiceUpdateDepartmentArgs)
	realResult := result.(*department.DepartmentRPCServiceUpdateDepartmentResult)
	success, err := handler.(department.DepartmentRPCService).UpdateDepartment(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDepartmentRPCServiceUpdateDepartmentArgs() interface{} {
	return department.NewDepartmentRPCServiceUpdateDepartmentArgs()
}

func newDepartmentRPCServiceUpdateDepartmentResult() interface{} {
	return department.NewDepartmentRPCServiceUpdateDepartmentResult()
}

func deleteDepartmentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*department.DepartmentRPCServiceDeleteDepartmentArgs)
	realResult := result.(*department.DepartmentRPCServiceDeleteDepartmentResult)
	success, err := handler.(department.DepartmentRPCService).DeleteDepartment(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDepartmentRPCServiceDeleteDepartmentArgs() interface{} {
	return department.NewDepartmentRPCServiceDeleteDepartmentArgs()
}

func newDepartmentRPCServiceDeleteDepartmentResult() interface{} {
	return department.NewDepartmentRPCServiceDeleteDepartmentResult()
}

func getsDepartmentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*department.DepartmentRPCServiceGetsDepartmentArgs)
	realResult := result.(*department.DepartmentRPCServiceGetsDepartmentResult)
	success, err := handler.(department.DepartmentRPCService).GetsDepartment(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDepartmentRPCServiceGetsDepartmentArgs() interface{} {
	return department.NewDepartmentRPCServiceGetsDepartmentArgs()
}

func newDepartmentRPCServiceGetsDepartmentResult() interface{} {
	return department.NewDepartmentRPCServiceGetsDepartmentResult()
}

func appendCityHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*department.DepartmentRPCServiceAppendCityArgs)
	realResult := result.(*department.DepartmentRPCServiceAppendCityResult)
	success, err := handler.(department.DepartmentRPCService).AppendCity(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDepartmentRPCServiceAppendCityArgs() interface{} {
	return department.NewDepartmentRPCServiceAppendCityArgs()
}

func newDepartmentRPCServiceAppendCityResult() interface{} {
	return department.NewDepartmentRPCServiceAppendCityResult()
}

func deleteCityHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*department.DepartmentRPCServiceDeleteCityArgs)
	realResult := result.(*department.DepartmentRPCServiceDeleteCityResult)
	success, err := handler.(department.DepartmentRPCService).DeleteCity(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDepartmentRPCServiceDeleteCityArgs() interface{} {
	return department.NewDepartmentRPCServiceDeleteCityArgs()
}

func newDepartmentRPCServiceDeleteCityResult() interface{} {
	return department.NewDepartmentRPCServiceDeleteCityResult()
}

func getsCityHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*department.DepartmentRPCServiceGetsCityArgs)
	realResult := result.(*department.DepartmentRPCServiceGetsCityResult)
	success, err := handler.(department.DepartmentRPCService).GetsCity(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDepartmentRPCServiceGetsCityArgs() interface{} {
	return department.NewDepartmentRPCServiceGetsCityArgs()
}

func newDepartmentRPCServiceGetsCityResult() interface{} {
	return department.NewDepartmentRPCServiceGetsCityResult()
}

func getsCityByDepartmentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*department.DepartmentRPCServiceGetsCityByDepartmentArgs)
	realResult := result.(*department.DepartmentRPCServiceGetsCityByDepartmentResult)
	success, err := handler.(department.DepartmentRPCService).GetsCityByDepartment(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDepartmentRPCServiceGetsCityByDepartmentArgs() interface{} {
	return department.NewDepartmentRPCServiceGetsCityByDepartmentArgs()
}

func newDepartmentRPCServiceGetsCityByDepartmentResult() interface{} {
	return department.NewDepartmentRPCServiceGetsCityByDepartmentResult()
}

func getDepartmentInfosByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*department.DepartmentRPCServiceGetDepartmentInfosByIdArgs)
	realResult := result.(*department.DepartmentRPCServiceGetDepartmentInfosByIdResult)
	success, err := handler.(department.DepartmentRPCService).GetDepartmentInfosById(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDepartmentRPCServiceGetDepartmentInfosByIdArgs() interface{} {
	return department.NewDepartmentRPCServiceGetDepartmentInfosByIdArgs()
}

func newDepartmentRPCServiceGetDepartmentInfosByIdResult() interface{} {
	return department.NewDepartmentRPCServiceGetDepartmentInfosByIdResult()
}

func getCityInfoByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*department.DepartmentRPCServiceGetCityInfoByIdArgs)
	realResult := result.(*department.DepartmentRPCServiceGetCityInfoByIdResult)
	success, err := handler.(department.DepartmentRPCService).GetCityInfoById(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDepartmentRPCServiceGetCityInfoByIdArgs() interface{} {
	return department.NewDepartmentRPCServiceGetCityInfoByIdArgs()
}

func newDepartmentRPCServiceGetCityInfoByIdResult() interface{} {
	return department.NewDepartmentRPCServiceGetCityInfoByIdResult()
}

func getDepartmentCityInfoByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*department.DepartmentRPCServiceGetDepartmentCityInfoByIdArgs)
	realResult := result.(*department.DepartmentRPCServiceGetDepartmentCityInfoByIdResult)
	success, err := handler.(department.DepartmentRPCService).GetDepartmentCityInfoById(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDepartmentRPCServiceGetDepartmentCityInfoByIdArgs() interface{} {
	return department.NewDepartmentRPCServiceGetDepartmentCityInfoByIdArgs()
}

func newDepartmentRPCServiceGetDepartmentCityInfoByIdResult() interface{} {
	return department.NewDepartmentRPCServiceGetDepartmentCityInfoByIdResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) AppendDepartment(ctx context.Context, request *department.AppendDepartmentRPCRequest) (r *department.AppendDepartmentRPCResponse, err error) {
	var _args department.DepartmentRPCServiceAppendDepartmentArgs
	_args.Request = request
	var _result department.DepartmentRPCServiceAppendDepartmentResult
	if err = p.c.Call(ctx, "AppendDepartment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateDepartment(ctx context.Context, request *department.UpdateDepartmentRPCRequest) (r *department.UpdateDepartmentRPCResponse, err error) {
	var _args department.DepartmentRPCServiceUpdateDepartmentArgs
	_args.Request = request
	var _result department.DepartmentRPCServiceUpdateDepartmentResult
	if err = p.c.Call(ctx, "UpdateDepartment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteDepartment(ctx context.Context, request *department.DeleteDepartmentRPCRequest) (r *department.DeleteDepartmentRPCResponse, err error) {
	var _args department.DepartmentRPCServiceDeleteDepartmentArgs
	_args.Request = request
	var _result department.DepartmentRPCServiceDeleteDepartmentResult
	if err = p.c.Call(ctx, "DeleteDepartment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetsDepartment(ctx context.Context, request *department.GetsDepartmentRPCRequest) (r *department.GetsDepartmentRPCResponse, err error) {
	var _args department.DepartmentRPCServiceGetsDepartmentArgs
	_args.Request = request
	var _result department.DepartmentRPCServiceGetsDepartmentResult
	if err = p.c.Call(ctx, "GetsDepartment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AppendCity(ctx context.Context, request *department.AppendCityRPCRequest) (r *department.AppendCityRPCResponse, err error) {
	var _args department.DepartmentRPCServiceAppendCityArgs
	_args.Request = request
	var _result department.DepartmentRPCServiceAppendCityResult
	if err = p.c.Call(ctx, "AppendCity", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteCity(ctx context.Context, request *department.DeleteCityRPCRequest) (r *department.DeleteCityRPCResponse, err error) {
	var _args department.DepartmentRPCServiceDeleteCityArgs
	_args.Request = request
	var _result department.DepartmentRPCServiceDeleteCityResult
	if err = p.c.Call(ctx, "DeleteCity", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetsCity(ctx context.Context, request *department.GetsCityRPCRequest) (r *department.GetsCityRPCResponse, err error) {
	var _args department.DepartmentRPCServiceGetsCityArgs
	_args.Request = request
	var _result department.DepartmentRPCServiceGetsCityResult
	if err = p.c.Call(ctx, "GetsCity", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetsCityByDepartment(ctx context.Context, request *department.GetsCityByDepartmentRPCRequest) (r *department.GetsCityByDepartmentRPCResponse, err error) {
	var _args department.DepartmentRPCServiceGetsCityByDepartmentArgs
	_args.Request = request
	var _result department.DepartmentRPCServiceGetsCityByDepartmentResult
	if err = p.c.Call(ctx, "GetsCityByDepartment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetDepartmentInfosById(ctx context.Context, request *department.GetDepartmentInfosByIdRPCRequest) (r *department.GetDepartmentInfosByIdRPCResponse, err error) {
	var _args department.DepartmentRPCServiceGetDepartmentInfosByIdArgs
	_args.Request = request
	var _result department.DepartmentRPCServiceGetDepartmentInfosByIdResult
	if err = p.c.Call(ctx, "GetDepartmentInfosById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCityInfoById(ctx context.Context, request *department.GetCityInfoByIdRPCRequest) (r *department.GetCityInfoByIdRPCResponse, err error) {
	var _args department.DepartmentRPCServiceGetCityInfoByIdArgs
	_args.Request = request
	var _result department.DepartmentRPCServiceGetCityInfoByIdResult
	if err = p.c.Call(ctx, "GetCityInfoById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetDepartmentCityInfoById(ctx context.Context, request *department.GetDepartmentCityInfoByIdRPCRequest) (r *department.GetDepartmentCityInfoByIdRPCResponse, err error) {
	var _args department.DepartmentRPCServiceGetDepartmentCityInfoByIdArgs
	_args.Request = request
	var _result department.DepartmentRPCServiceGetDepartmentCityInfoByIdResult
	if err = p.c.Call(ctx, "GetDepartmentCityInfoById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
