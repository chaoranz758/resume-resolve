package handler

import (
	"context"
	department "resume-resolving/api/idl/service/department/kitex_gen/department"
	"resume-resolving/internal/app/service/department/service"
)

// DepartmentRPCServiceImpl implements the last service interface defined in the IDL.
type DepartmentRPCServiceImpl struct{}

// AppendDepartment implements the DepartmentRPCServiceImpl interface.
func (s *DepartmentRPCServiceImpl) AppendDepartment(ctx context.Context, request *department.AppendDepartmentRPCRequest) (resp *department.AppendDepartmentRPCResponse, err error) {
	code, message, data, err := service.AppendDepartment(request)
	return &department.AppendDepartmentRPCResponse{
		Code:    int32(code),
		Message: message,
		Data:    data,
	}, nil
}

// UpdateDepartment implements the DepartmentRPCServiceImpl interface.
func (s *DepartmentRPCServiceImpl) UpdateDepartment(ctx context.Context, request *department.UpdateDepartmentRPCRequest) (resp *department.UpdateDepartmentRPCResponse, err error) {
	code, message, data, err := service.UpdateDepartment(request)
	return &department.UpdateDepartmentRPCResponse{
		Code:    int32(code),
		Message: message,
		Data:    data,
	}, nil
}

// DeleteDepartment implements the DepartmentRPCServiceImpl interface.
func (s *DepartmentRPCServiceImpl) DeleteDepartment(ctx context.Context, request *department.DeleteDepartmentRPCRequest) (resp *department.DeleteDepartmentRPCResponse, err error) {
	code, message, err := service.DeleteDepartment(request)
	return &department.DeleteDepartmentRPCResponse{
		Code:    int32(code),
		Message: message,
	}, nil
}

// GetsDepartment implements the DepartmentRPCServiceImpl interface.
func (s *DepartmentRPCServiceImpl) GetsDepartment(ctx context.Context, request *department.GetsDepartmentRPCRequest) (resp *department.GetsDepartmentRPCResponse, err error) {
	code, message, data, err := service.GetsDepartment(request)
	return &department.GetsDepartmentRPCResponse{
		Code:    int32(code),
		Message: message,
		Data:    data,
	}, nil
}

// AppendCity implements the DepartmentRPCServiceImpl interface.
func (s *DepartmentRPCServiceImpl) AppendCity(ctx context.Context, request *department.AppendCityRPCRequest) (resp *department.AppendCityRPCResponse, err error) {
	code, message, data, err := service.AppendCity(request)
	return &department.AppendCityRPCResponse{
		Code:    int32(code),
		Message: message,
		Data:    data,
	}, nil
}

// DeleteCity implements the DepartmentRPCServiceImpl interface.
func (s *DepartmentRPCServiceImpl) DeleteCity(ctx context.Context, request *department.DeleteCityRPCRequest) (resp *department.DeleteCityRPCResponse, err error) {
	code, message, err := service.DeleteCity(request)
	return &department.DeleteCityRPCResponse{
		Code:    int32(code),
		Message: message,
	}, nil
}

// GetsCity implements the DepartmentRPCServiceImpl interface.
func (s *DepartmentRPCServiceImpl) GetsCity(ctx context.Context, request *department.GetsCityRPCRequest) (resp *department.GetsCityRPCResponse, err error) {
	code, message, data, err := service.GetsCity(request)
	return &department.GetsCityRPCResponse{
		Code:    int32(code),
		Message: message,
		Data:    data,
	}, nil
}

// GetsCityByDepartment implements the DepartmentRPCServiceImpl interface.
func (s *DepartmentRPCServiceImpl) GetsCityByDepartment(ctx context.Context, request *department.GetsCityByDepartmentRPCRequest) (resp *department.GetsCityByDepartmentRPCResponse, err error) {
	code, message, data, err := service.GetsCityByDepartment(request)
	return &department.GetsCityByDepartmentRPCResponse{
		Code:    int32(code),
		Message: message,
		Data:    data,
	}, nil
}

// GetDepartmentInfosById implements the DepartmentRPCServiceImpl interface.
func (s *DepartmentRPCServiceImpl) GetDepartmentInfosById(ctx context.Context, request *department.GetDepartmentInfosByIdRPCRequest) (resp *department.GetDepartmentInfosByIdRPCResponse, err error) {
	code, message, data, err := service.GetDepartmentInfosById(request)
	return &department.GetDepartmentInfosByIdRPCResponse{
		Code:               int32(code),
		Message:            message,
		DepartmentInfoList: data,
	}, nil
}

// GetCityInfoById implements the DepartmentRPCServiceImpl interface.
func (s *DepartmentRPCServiceImpl) GetCityInfoById(ctx context.Context, request *department.GetCityInfoByIdRPCRequest) (resp *department.GetCityInfoByIdRPCResponse, err error) {
	code, message, data, err := service.GetCityInfoById(request)
	return &department.GetCityInfoByIdRPCResponse{
		Code:         int32(code),
		Message:      message,
		CityInfoList: data,
	}, nil
}

// GetDepartmentCityInfoById implements the DepartmentRPCServiceImpl interface.
func (s *DepartmentRPCServiceImpl) GetDepartmentCityInfoById(ctx context.Context, request *department.GetDepartmentCityInfoByIdRPCRequest) (resp *department.GetDepartmentCityInfoByIdRPCResponse, err error) {
	code, message, departmentData, cityData, err := service.GetDepartmentCityInfoById(request)
	return &department.GetDepartmentCityInfoByIdRPCResponse{
		Code:               int32(code),
		Message:            message,
		DepartmentInfoList: departmentData,
		CityInfoList:       cityData,
	}, nil
}
