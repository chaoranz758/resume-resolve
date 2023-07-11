namespace go post

include "../../base/common.thrift"

struct AppendPostRequest {
    1: i64 hr_id (api.body = "hr_id")
    2: string post_brief (api.body = "post_brief")
    3: string post_description (api.body = "post_description")
    4: string post_require (api.body = "post_require")
    5: byte is_school_recruitment (api.body = "is_school_recruitment")
    6: byte is_internship (api.body = "is_internship")
    7: i64 post_category_id (api.body = "post_category_id")
    8: i64 department_id (api.body = "department_id")
    9: list<i64> city_list (api.body = "city_list")
}

struct UpdatePostRequest {
    1: i64 post_id (api.body = "post_id")
    2: string post_brief (api.body = "post_brief")
    3: string post_description (api.body = "post_description")
    4: string post_require (api.body = "post_require")
}

struct DeletePostRequest {
    1: i64 post_id (api.body = "post_id")
}

struct GetsPostInUserRequest {
    1: i64 limit (api.body = "limit")
    2: i64 offset (api.body = "offset")
    3: byte is_new (api.body = "is_new")
    4: byte is_school_recruitment (api.body = "is_school_recruitment")
    5: byte is_internship (api.body = "is_internship")
    6: list<i64> department_id_list (api.body = "department_id_list")
    7: list<i64> post_category_id_list (api.body = "post_category_id_list")
}

struct GetsPostInHRRequest {
    1: i64 limit (api.query = "limit")
    2: i64 offset (api.query = "offset")
    3: i64 hr_id (api.query = "hr_id")
    4: byte is_school_recruitment (api.query = "is_school_recruitment")
    5: byte is_internship (api.query = "is_internship")
}

struct DeliveryPostRequest {
    1: byte is_delivery (api.body = "is_delivery")
    2: i64 user_id (api.body = "user_id")
    3: i64 post_id (api.body = "post_id")
}

struct CollectPostRequest {
    1: byte is_collect (api.body = "is_collect")
    2: i64 user_id (api.body = "user_id")
    3: i64 post_id (api.body = "post_id")
}

struct GetDeliveryPostRequest {
    1: i64 user_id (api.query = "user_id")
}

struct GetCollectPostRequest {
    1: i64 user_id (api.query = "user_id")
}

struct UpdateResumeStatusRequest {
    1: byte resume_operate (api.body = "resume_operate") //0-通 1-不通过并放入人才库 2-不通过不放入人才库
    2: i64 user_id (api.body = "user_id")
    3: i64 post_id (api.body = "post_id")
}

service PostService {
    //增加岗位
    common.NilResponse AppendPost(1: AppendPostRequest request) (api.post = "/api/v1/post/append")
    //修改岗位信息
    common.NilResponse UpdatePost(1: UpdatePostRequest request) (api.post = "/api/v1/post/update")
    //删除岗位
    common.NilResponse DeletePost(1: DeletePostRequest request) (api.post = "/api/v1/post/delete")
    //查询岗位信息-求职者界面
    common.NilResponse GetsPostInUser(1: GetsPostInUserRequest request) (api.post = "/api/v1/post/gets-user")
    //查询岗位信息-HR界面
    common.NilResponse GetsPostInHR(1: GetsPostInHRRequest request) (api.get = "/api/v1/post/gets-hr")
    //投递岗位
    common.NilResponse DeliveryPost(1: DeliveryPostRequest request) (api.post = "/api/v1/post/delivery")
    //岗位收藏
    common.NilResponse CollectPost(1: CollectPostRequest request) (api.post = "/api/v1/post/collect")
    //查看投递情况
    common.NilResponse GetDeliveryPost(1: GetDeliveryPostRequest request) (api.get = "/api/v1/post/get-delivery")
    //查看收藏情况
    common.NilResponse GetCollectPost(1: GetCollectPostRequest request) (api.get = "/api/v1/post/get-collect")
    //HR修改投递的简历状态
    common.NilResponse UpdateResumeStatus(1: UpdateResumeStatusRequest request) (api.post = "/api/v1/post/update-resume")
}