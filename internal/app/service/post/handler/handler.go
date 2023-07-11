package handler

import (
	"context"
	post "resume-resolving/api/idl/service/post/kitex_gen/post"
	"resume-resolving/internal/app/service/post/service"
)

// PostRPCServiceImpl implements the last service interface defined in the IDL.
type PostRPCServiceImpl struct{}

// AppendPost implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) AppendPost(ctx context.Context, request *post.AppendPostRPCRequest) (resp *post.AppendPostRPCResponse, err error) {
	code, message, err := service.AppendPost(request)
	return &post.AppendPostRPCResponse{
		Code:    code,
		Message: message,
	}, nil
}

// UpdatePost implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) UpdatePost(ctx context.Context, request *post.UpdatePostRPCRequest) (resp *post.UpdatePostRPCResponse, err error) {
	code, message, err := service.UpdatePost(request)
	return &post.UpdatePostRPCResponse{
		Code:    code,
		Message: message,
	}, nil
}

// DeletePost implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) DeletePost(ctx context.Context, request *post.DeletePostRPCRequest) (resp *post.DeletePostRPCResponse, err error) {
	code, message, err := service.DeletePost(request)
	return &post.DeletePostRPCResponse{
		Code:    code,
		Message: message,
	}, nil
}

// GetsPostInUser implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) GetsPostInUser(ctx context.Context, request *post.GetsPostInUserRPCRequest) (resp *post.GetsPostInUserRPCResponse, err error) {
	code, message, data, err := service.GetsPostInUser(request)
	return &post.GetsPostInUserRPCResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}, nil
}

// GetsPostInHR implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) GetsPostInHR(ctx context.Context, request *post.GetsPostInHRRPCRequest) (resp *post.GetsPostInHRRPCResponse, err error) {
	code, message, data, err := service.GetsPostInHR(request)
	return &post.GetsPostInHRRPCResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}, nil
}

// DeliveryPost implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) DeliveryPost(ctx context.Context, request *post.DeliveryPostRPCRequest) (resp *post.DeliveryPostRPCResponse, err error) {
	code, message, err := service.DeliveryPost(request)
	return &post.DeliveryPostRPCResponse{
		Code:    code,
		Message: message,
	}, nil
}

// CollectPost implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) CollectPost(ctx context.Context, request *post.CollectPostRPCRequest) (resp *post.CollectPostRPCResponse, err error) {
	code, message, err := service.CollectPost(request)
	return &post.CollectPostRPCResponse{
		Code:    code,
		Message: message,
	}, nil
}

// GetDeliveryPost implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) GetDeliveryPost(ctx context.Context, request *post.GetDeliveryPostRPCRequest) (resp *post.GetDeliveryPostRPCResponse, err error) {
	code, message, data, err := service.GetDeliveryPost(request)
	return &post.GetDeliveryPostRPCResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}, nil
}

// GetCollectPost implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) GetCollectPost(ctx context.Context, request *post.GetCollectPostRPCRequest) (resp *post.GetCollectPostRPCResponse, err error) {
	code, message, data, err := service.GetCollectPost(request)
	return &post.GetCollectPostRPCResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}, nil
}

// UpdateResumeStatus implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) UpdateResumeStatus(ctx context.Context, request *post.UpdateResumeStatusRPCRequest) (resp *post.UpdateResumeStatusRPCResponse, err error) {
	code, message, err := service.UpdateResumeStatus(request)
	return &post.UpdateResumeStatusRPCResponse{
		Code:    code,
		Message: message,
	}, nil
}

// GetUserIdByPostId implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) GetUserIdByPostId(ctx context.Context, request *post.GetUserIdByPostIdRPCRequest) (resp *post.GetUserIdByPostIdRPCResponse, err error) {
	code, message, data, err := service.GetUserIdByPostId(request)
	return &post.GetUserIdByPostIdRPCResponse{
		Code:       code,
		Message:    message,
		UserIdList: data,
	}, nil
}

// DeleteResumeRelativeInfoByHRId implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) DeleteResumeRelativeInfoByHRId(ctx context.Context, request *post.DeleteResumeRelativeInfoByHRIdRPCRequest) (resp *post.DeleteResumeRelativeInfoByHRIdRPCResponse, err error) {
	code, message, err := service.DeleteResumeRelativeInfoByHRId(request)
	return &post.DeleteResumeRelativeInfoByHRIdRPCResponse{
		Code:    code,
		Message: message,
	}, nil
}

// DeleteResumeRelativeInfoByDepartmentId implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) DeleteResumeRelativeInfoByDepartmentId(ctx context.Context, request *post.DeleteResumeRelativeInfoByDepartmentIdRPCRequest) (resp *post.DeleteResumeRelativeInfoByDepartmentIdRPCResponse, err error) {
	code, message, err := service.DeleteResumeRelativeInfoByDepartmentId(request)
	return &post.DeleteResumeRelativeInfoByDepartmentIdRPCResponse{
		Code:    code,
		Message: message,
	}, nil
}

// DeleteResumeRelativeInfoByPostCategoryIdList implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) DeleteResumeRelativeInfoByPostCategoryIdList(ctx context.Context, request *post.DeleteResumeRelativeInfoByPostCategoryIdListRPCRequest) (resp *post.DeleteResumeRelativeInfoByPostCategoryIdListRPCResponse, err error) {
	code, message, err := service.DeleteResumeRelativeInfoByPostCategoryIdList(request)
	return &post.DeleteResumeRelativeInfoByPostCategoryIdListRPCResponse{
		Code:    code,
		Message: message,
	}, nil
}

// DeletePostCityByCityId implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) DeletePostCityByCityId(ctx context.Context, request *post.DeletePostCityByCityIdRPCRequest) (resp *post.DeletePostCityByCityIdRPCResposne, err error) {
	code, message, err := service.DeletePostCityByCityId(request)
	return &post.DeletePostCityByCityIdRPCResposne{
		Code:    code,
		Message: message,
	}, nil
}
