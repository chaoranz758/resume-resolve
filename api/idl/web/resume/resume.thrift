namespace go resume

include "../../base/common.thrift"
include "../../base/resume.thrift"

struct UploadResumeFileRequest {
    1: i64 user_id (api.form = "user_id")
}

struct UploadStructResumeRequest {
    1: resume.BasicInfo basic_information (api.body = "basic_information")
    2: list<resume.EducationalExperienceList> educational_experience_list (api.body = "educational_experience_list")
    3: list<resume.InternshipExperienceList> internship_experience_list (api.body = "internship_experience_list")
    4: list<resume.WorkExperienceList> work_experience_list (api.body = "work_experience_list")
    5: list<resume.ProjectExperienceList> project_experience_list (api.body = "project_experience_list")
    6: list<resume.ContestList> contest_list (api.body = "contest_list")
    7: list<resume.CertificateList> certificate_list (api.body = "certificate_list")
    8: list<resume.LanguageList> language_list (api.body = "language_list")
    9: list<resume.SocialList> social_list (api.body = "social_list")
}

struct GetResumeByIdRequest {
    1: i64 user_id (api.query = "user_id")
}

struct GetResumeByPostRequest {
    1: i64 post_id (api.query = "post_id")
    2: byte is_talent_pool (api.query = "is_talent_pool")
    3: i32 limit (api.query = "limit")
    4: i32 offset (api.query = "offset")
}

service ResumeService {
    //上传简历文件并返回解析结果
    common.NilResponse UploadResumeFile(1: UploadResumeFileRequest request) (api.post = "/api/v1/resume/upload-file")
    //提交结构化简历数据
    common.NilResponse UploadStructResume(1: UploadStructResumeRequest request) (api.post = "/api/v1/resume/upload")
    //获取单一用户的结构化简历数据
    common.NilResponse GetResumeById(1: GetResumeByIdRequest request) (api.get = "/api/v1/resume/get-id")
    //获取某个岗位的简历投递或人才库推荐信息-附带人才画像
    common.NilResponse GetResumeByPost(1: GetResumeByPostRequest request) (api.get = "/api/v1/resume/gets-post")
}