package handler

import (
	"context"
	"resume-resolving/api/idl/service/user/kitex_gen/user"
	"resume-resolving/internal/app/service/user/service"
)

// UserRPCServiceImpl implements the last service interface defined in the IDL.
type UserRPCServiceImpl struct{}

// UserChangePassword implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) UserChangePassword(ctx context.Context, request *user.UserChangePasswordRPCRequest) (resp *user.UserChangePasswordRPCResponse, err error) {
	code, message, err := service.UserChangePassword(request)
	return &user.UserChangePasswordRPCResponse{
		Code:    int32(code),
		Message: message,
	}, nil
}

// UserLogin implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) UserLogin(ctx context.Context, request *user.UserLoginRPCRequest) (resp *user.UserLoginRPCResponse, err error) {
	code, message, err := service.UserLogin(request)
	return &user.UserLoginRPCResponse{
		Code:    int32(code),
		Message: message,
	}, nil
}

// UserRegister implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) UserRegister(ctx context.Context, request *user.UserRegisterRPCRequest) (resp *user.UserRegisterRPCResponse, err error) {
	code, message, err := service.UserRegister(request)
	return &user.UserRegisterRPCResponse{
		Code:    int32(code),
		Message: message,
	}, nil
}

// HRRegister implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) HRRegister(ctx context.Context, request *user.HRRegisterRPCRequest) (resp *user.HRRegisterRPCResponse, err error) {
	code, message, err := service.HRRegister(request)
	return &user.HRRegisterRPCResponse{
		Code:    int32(code),
		Message: message,
	}, nil
}

// HRDelete implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) HRDelete(ctx context.Context, request *user.HRDeleteRPCRequest) (resp *user.HRDeleteRPCResponse, err error) {
	code, message, err := service.HRDelete(request)
	return &user.HRDeleteRPCResponse{
		Code:    int32(code),
		Message: message,
	}, nil
}

// GetHRByDepartment implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) GetHRByDepartment(ctx context.Context, request *user.GetHRByDepartmentRPCRequest) (resp *user.GetHRByDepartmentRPCResponse, err error) {
	code, message, data, err := service.GetHRByDepartment(request)
	return &user.GetHRByDepartmentRPCResponse{
		Code:    int32(code),
		Message: message,
		Data:    data,
	}, nil
}

// GetUserByName implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) GetUserByName(ctx context.Context, request *user.GetUserByNameRPCRequest) (resp *user.GetUserByNameRPCResponse, err error) {
	code, message, data, err := service.GetUserByName(request)
	return &user.GetUserByNameRPCResponse{
		Code:    int32(code),
		Message: message,
		Data:    data,
	}, nil
}

// GetHRByName implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) GetHRByName(ctx context.Context, request *user.GetHRByNameRPCRequest) (resp *user.GetHRByNameRPCResponse, err error) {
	code, message, data, err := service.GetHRByName(request)
	return &user.GetHRByNameRPCResponse{
		Code:    int32(code),
		Message: message,
		Data:    data,
	}, nil
}

// DeleteHRByDepartmentId implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) DeleteHRByDepartmentId(ctx context.Context, request *user.DeleteHRByDepartmentIdRPCRequest) (resp *user.DeleteHRByDepartmentIdRPCResponse, err error) {
	code, message, err := service.DeleteHRByDepartmentId(request)
	return &user.DeleteHRByDepartmentIdRPCResponse{
		Code:    int32(code),
		Message: message,
	}, nil
}
