package main

import (
	"context"
	post "resume-resolving/api/idl/service/post/kitex_gen/post"
)

// PostRPCServiceImpl implements the last service interface defined in the IDL.
type PostRPCServiceImpl struct{}

// AppendPost implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) AppendPost(ctx context.Context, request *post.AppendPostRPCRequest) (resp *post.AppendPostRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdatePost implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) UpdatePost(ctx context.Context, request *post.UpdatePostRPCRequest) (resp *post.UpdatePostRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// DeletePost implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) DeletePost(ctx context.Context, request *post.DeletePostRPCRequest) (resp *post.DeletePostRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// GetsPostInUser implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) GetsPostInUser(ctx context.Context, request *post.GetsPostInUserRPCRequest) (resp *post.GetsPostInUserRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// GetsPostInHR implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) GetsPostInHR(ctx context.Context, request *post.GetsPostInHRRPCRequest) (resp *post.GetsPostInHRRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// DeliveryPost implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) DeliveryPost(ctx context.Context, request *post.DeliveryPostRPCRequest) (resp *post.DeliveryPostRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// CollectPost implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) CollectPost(ctx context.Context, request *post.CollectPostRPCRequest) (resp *post.CollectPostRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// GetDeliveryPost implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) GetDeliveryPost(ctx context.Context, request *post.GetDeliveryPostRPCRequest) (resp *post.GetDeliveryPostRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// GetCollectPost implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) GetCollectPost(ctx context.Context, request *post.GetCollectPostRPCRequest) (resp *post.GetCollectPostRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateResumeStatus implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) UpdateResumeStatus(ctx context.Context, request *post.UpdateResumeStatusRPCRequest) (resp *post.UpdateResumeStatusRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserIdByPostId implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) GetUserIdByPostId(ctx context.Context, request *post.GetUserIdByPostIdRPCRequest) (resp *post.GetUserIdByPostIdRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteResumeRelativeInfoByHRId implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) DeleteResumeRelativeInfoByHRId(ctx context.Context, request *post.DeleteResumeRelativeInfoByHRIdRPCRequest) (resp *post.DeleteResumeRelativeInfoByHRIdRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteResumeRelativeInfoByDepartmentId implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) DeleteResumeRelativeInfoByDepartmentId(ctx context.Context, request *post.DeleteResumeRelativeInfoByDepartmentIdRPCRequest) (resp *post.DeleteResumeRelativeInfoByDepartmentIdRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteResumeRelativeInfoByPostCategoryIdList implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) DeleteResumeRelativeInfoByPostCategoryIdList(ctx context.Context, request *post.DeleteResumeRelativeInfoByPostCategoryIdListRPCRequest) (resp *post.DeleteResumeRelativeInfoByPostCategoryIdListRPCResponse, err error) {
	// TODO: Your code here...
	return
}

// DeletePostCityByCityId implements the PostRPCServiceImpl interface.
func (s *PostRPCServiceImpl) DeletePostCityByCityId(ctx context.Context, request *post.DeletePostCityByCityIdRPCRequest) (resp *post.DeletePostCityByCityIdRPCResposne, err error) {
	// TODO: Your code here...
	return
}
