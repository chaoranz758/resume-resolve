package handler

import (
	"context"
	post_category "resume-resolving/api/idl/service/post_category/kitex_gen/post_category"
	"resume-resolving/internal/app/service/post_category/service"
)

// PostCategoryRPCServiceImpl implements the last service interface defined in the IDL.
type PostCategoryRPCServiceImpl struct{}

// AppendPostCategory implements the PostCategoryRPCServiceImpl interface.
func (s *PostCategoryRPCServiceImpl) AppendPostCategory(ctx context.Context, request *post_category.AppendPostCategoryRPCRequest) (resp *post_category.AppendPostCategoryRPCResponse, err error) {
	code, message, data, err := service.AppendPostCategory(request)
	return &post_category.AppendPostCategoryRPCResponse{
		Code:    int32(code),
		Message: message,
		Data:    data,
	}, nil
}

// UpdatePostCategory implements the PostCategoryRPCServiceImpl interface.
func (s *PostCategoryRPCServiceImpl) UpdatePostCategory(ctx context.Context, request *post_category.UpdatePostCategoryRPCRequest) (resp *post_category.UpdatePostCategoryRPCResponse, err error) {
	code, message, data, err := service.UpdatePostCategory(request)
	return &post_category.UpdatePostCategoryRPCResponse{
		Code:    int32(code),
		Message: message,
		Data:    data,
	}, nil
}

// DeletePostCategory implements the PostCategoryRPCServiceImpl interface.
func (s *PostCategoryRPCServiceImpl) DeletePostCategory(ctx context.Context, request *post_category.DeletePostCategoryRPCRequest) (resp *post_category.DeletePostCategoryRPCResponse, err error) {
	code, message, err := service.DeletePostCategory(request)
	return &post_category.DeletePostCategoryRPCResponse{
		Code:    int32(code),
		Message: message,
	}, nil
}

// GetsPostCategory implements the PostCategoryRPCServiceImpl interface.
func (s *PostCategoryRPCServiceImpl) GetsPostCategory(ctx context.Context, request *post_category.GetsPostCategoryRPCRequest) (resp *post_category.GetsPostCategoryRPCResponse, err error) {
	code, message, data, err := service.GetsPostCategory(request)
	return &post_category.GetsPostCategoryRPCResponse{
		Code:    int32(code),
		Message: message,
		Data:    data,
	}, nil
}

// GetPostCategoryById implements the PostCategoryRPCServiceImpl interface.
func (s *PostCategoryRPCServiceImpl) GetPostCategoryById(ctx context.Context, request *post_category.GetPostCategoryByIdRPCRequest) (resp *post_category.GetPostCategoryByIdRPCResponse, err error) {
	code, message, data, err := service.GetPostCategoryById(request)
	return &post_category.GetPostCategoryByIdRPCResponse{
		Code:                 int32(code),
		Message:              message,
		PostCategoryInfoList: data,
	}, nil
}
