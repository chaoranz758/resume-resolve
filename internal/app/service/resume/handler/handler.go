package handler

import (
	"context"
	resume "resume-resolving/api/idl/service/resume/kitex_gen/resume"
	"resume-resolving/internal/app/service/resume/service"
)

// ResumeRPCServiceImpl implements the last service interface defined in the IDL.
type ResumeRPCServiceImpl struct{}

// UploadStructResume implements the ResumeRPCServiceImpl interface.
func (s *ResumeRPCServiceImpl) UploadStructResume(ctx context.Context, request *resume.UploadStructResumeRPCRequest) (resp *resume.UploadStructResumeRPCResponse, err error) {
	code, message, err := service.UploadStructResume(request)
	return &resume.UploadStructResumeRPCResponse{
		Code:    code,
		Message: message,
	}, nil
}

// GetResumeById implements the ResumeRPCServiceImpl interface.
func (s *ResumeRPCServiceImpl) GetResumeById(ctx context.Context, request *resume.GetResumeByIdRPCRequest) (resp *resume.GetResumeByIdRPCResponse, err error) {
	code, message, data, err := service.GetResumeById(request)
	return &resume.GetResumeByIdRPCResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}, nil
}

// GetResumeByPost implements the ResumeRPCServiceImpl interface.
func (s *ResumeRPCServiceImpl) GetResumeByPost(ctx context.Context, request *resume.GetResumeByPostRPCRequest) (resp *resume.GetResumeByPostRPCResponse, err error) {
	code, message, data, err := service.GetResumeByPost(request)
	return &resume.GetResumeByPostRPCResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}, nil
}
