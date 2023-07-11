package main

import (
	"context"
	user "resume-resolving/api/idl/service/user/kitex_gen/user"
)

// UserRPCServiceImpl implements the last service interface defined in the IDL.
type UserRPCServiceImpl struct{}

// UserChangePassword implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) UserChangePassword(ctx context.Context, request *user.UserChangePasswordRPCRequest) (resp *user.UserChangePasswordRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// UserLogin implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) UserLogin(ctx context.Context, request *user.UserLoginRPCRequest) (resp *user.UserLoginRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// UserRegister implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) UserRegister(ctx context.Context, request *user.UserRegisterRPCRequest) (resp *user.UserRegisterRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// HRRegister implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) HRRegister(ctx context.Context, request *user.HRRegisterRPCRequest) (resp *user.HRRegisterRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// HRDelete implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) HRDelete(ctx context.Context, request *user.HRDeleteRPCRequest) (resp *user.HRDeleteRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// GetHRByDepartment implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) GetHRByDepartment(ctx context.Context, request *user.GetHRByDepartmentRPCRequest) (resp *user.GetHRByDepartmentRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserByName implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) GetUserByName(ctx context.Context, request *user.GetUserByNameRPCRequest) (resp *user.GetUserByNameRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// GetHRByName implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) GetHRByName(ctx context.Context, request *user.GetHRByNameRPCRequest) (resp *user.GetHRByNameRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteHRByDepartmentId implements the UserRPCServiceImpl interface.
func (s *UserRPCServiceImpl) DeleteHRByDepartmentId(ctx context.Context, request *user.DeleteHRByDepartmentIdRPCRequest) (resp *user.DeleteHRByDepartmentIdRPCResponse, err error) {
	// TODO: Your code here...
	return
}
