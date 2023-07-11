namespace go resume

include "../../base/resume.thrift"

struct UploadStructResumeRPCRequest {
    1: resume.BasicInfo basic_information
    2: list<resume.EducationalExperienceList> educational_experience_list
    3: list<resume.InternshipExperienceList> internship_experience_list
    4: list<resume.WorkExperienceList> work_experience_list
    5: list<resume.ProjectExperienceList> project_experience_list
    6: list<resume.ContestList> contest_list
    7: list<resume.CertificateList> certificate_list
    8: list<resume.LanguageList> language_list
    9: list<resume.SocialList> social_list
}

struct UploadStructResumeRPCResponse {
    1: i32 code
    2: string message
}

struct GetResumeByIdRPCRequest {
    1: i64 user_id
}

struct GetResumeByIdRPCResponse {
    1: i32 code
    2: string message
    3: ResumeInfo data
}

struct GetResumeByPostRPCRequest {
    1: i64 post_id
    2: byte is_talent_pool
    3: i32 limit
    4: i32 offset
}

struct GetResumeByPostRPCResponse {
    1: i32 code
    2: string message
    3: list<ResumeInfo> data
}

struct ResumeInfo {
    1: resume.BasicInfoWithResumeId basic_information
    2: list<resume.EducationalExperienceList> educational_experience_list
    3: list<resume.InternshipExperienceList> internship_experience_list
    4: list<resume.WorkExperienceList> work_experience_list
    5: list<resume.ProjectExperienceList> project_experience_list
    6: list<resume.ContestList> contest_list
    7: list<resume.CertificateList> certificate_list
    8: list<resume.LanguageList> language_list
    9: list<resume.SocialList> social_list
    10: TalentPortrait talent_portrait
}

struct TalentPortrait {
    1: i16 age
    2: string max_education
    3: string graduated_school
    4: i8 school_level
    5: i16 working_seniority
}

service ResumeRPCService {
    //与web交互服务
    //提交结构化简历数据
    UploadStructResumeRPCResponse UploadStructResume(1: UploadStructResumeRPCRequest request)
    //获取单一用户的结构化简历数据
    GetResumeByIdRPCResponse GetResumeById(1: GetResumeByIdRPCRequest request)
    //获取某个岗位的简历投递或人才库推荐信息-附带人才画像
    GetResumeByPostRPCResponse GetResumeByPost(1: GetResumeByPostRPCRequest request)
}