namespace go base

//基本信息
struct BasicInfo {
    1: string name
    2: string phone
    3: string email
    4: string self_evaluaton
    5: i64 birthday
    6: i64 user_id
    7: string resume_url
}

//基本信息
struct BasicInfoWithResumeId {
    1: string name
    2: string phone
    3: string email
    4: string self_evaluaton
    5: i64 birthday
    6: i64 user_id
    7: string resume_url
    8: i64 resume_id
}

//教育经历
struct EducationalExperienceList {
    1: string school
    2: string education
    3: string speciality
    4: i64 start_time
    5: i64 end_time
    6: string rank
}

//实习经历
struct InternshipExperienceList {
    1: string company
    2: string position
    3: i64 start_time
    4: i64 end_time
    5: string description
}

//工作经历
struct WorkExperienceList {
    1: string company
    2: string position
    3: i64 start_time
    4: i64 end_time
    5: string description
}

//项目经历
struct ProjectExperienceList {
    1: string project_name
    2: string project_role
    3: i64 start_time
    4: i64 end_time
    5: string project_description
    6: string project_url
}

//竞赛
struct ContestList {
    1: string name
    2: string description
}

//证书
struct CertificateList {
    1: string name
    2: string description
}

//语言
struct LanguageList {
    1: string language
    2: string proficiency_level
}

//社交账号
struct SocialList {
    1: string social_platform
    2: string platform_url
}