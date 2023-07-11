namespace go post

include "../../base/post.thrift"

struct AppendPostRPCRequest {
    1: i64 hr_id
    2: string post_brief
    3: string post_description
    4: string post_require
    5: byte is_school_recruitment
    6: byte is_internship
    7: i64 post_category_id
    8: i64 department_id
    9: list<i64> city_list
}

struct AppendPostRPCResponse {
    1: i32 code
    2: string message
}

struct UpdatePostRPCRequest {
    1: i64 post_id
    2: string post_brief
    3: string post_description
    4: string post_require
}

struct UpdatePostRPCResponse {
    1: i32 code
    2: string message
}

struct DeletePostRPCRequest {
    1: i64 post_id
}

struct DeletePostRPCResponse {
    1: i32 code
    2: string message
}

struct GetsPostInUserRPCRequest {
    1: i64 limit
    2: i64 offset
    3: byte is_new
    4: byte is_school_recruitment
    5: byte is_internship
    6: list<i64> department_id_list
    7: list<i64> post_category_id_list
}

struct GetsPostInUserRPCResponse {
    1: i32 code
    2: string message
    3: list<post.PostInfo> data
}

struct GetsPostInHRRPCRequest {
    1: i64 limit
    2: i64 offset
    3: i64 hr_id
    4: byte is_school_recruitment
    5: byte is_internship
}

struct GetsPostInHRRPCResponse {
    1: i32 code
    2: string message
    3: list<post.PostInfo> data
}

struct DeliveryPostRPCRequest {
    1: byte is_delivery
    2: i64 user_id
    3: i64 post_id
}

struct DeliveryPostRPCResponse {
    1: i32 code
    2: string message
}

struct CollectPostRPCRequest {
    1: byte is_collect
    2: i64 user_id
    3: i64 post_id
}

struct CollectPostRPCResponse {
    1: i32 code
    2: string message
}

struct GetDeliveryPostRPCRequest {
    1: i64 user_id
}

struct GetDeliveryPostRPCResponse {
    1: i32 code
    2: string message
    3: list<GetDeliveryPostRPCData> data
}

struct GetDeliveryPostRPCData {
    1: post.PostInfo post_information
    2: string resume_status
}

struct GetCollectPostRPCRequest {
    1: i64 user_id
}

struct GetCollectPostRPCResponse {
    1: i32 code
    2: string message
    3: list<post.PostInfo> data
}

struct UpdateResumeStatusRPCRequest {
    1: byte resume_operate //0-通 1-不通过并放入人才库 2-不通过不放入人才库
    2: i64 user_id
    3: i64 post_id
}

struct UpdateResumeStatusRPCResponse {
    1: i32 code
    2: string message
}

struct DeleteResumeRelativeInfoByHRIdRPCRequest {
    1: i64 hr_id
}

struct DeleteResumeRelativeInfoByHRIdRPCResponse {
    1: i32 code
    2: string message
}

struct DeleteResumeRelativeInfoByDepartmentIdRPCRequest {
    1: i64 department_id
}

struct DeleteResumeRelativeInfoByDepartmentIdRPCResponse {
    1: i32 code
    2: string message
}

struct DeleteResumeRelativeInfoByPostCategoryIdListRPCRequest {
    1: list<i64> post_category_id_list
}

struct DeleteResumeRelativeInfoByPostCategoryIdListRPCResponse {
    1: i32 code
    2: string message
}

struct GetUserIdByPostIdRPCRequest {
    1: i8 is_talent_pool
    2: i64 post_id
    3: i32 limit
    4: i32 offset
}

struct GetUserIdByPostIdRPCResponse {
    1: i32 code
    2: string message
    3: list<i64> user_id_list
}

struct DeletePostCityByCityIdRPCRequest {
    1: i64 city_id
}

struct DeletePostCityByCityIdRPCResposne {
     1: i32 code
    2: string message
}

service PostRPCService {
    //外部接口
    //增加岗位
    AppendPostRPCResponse AppendPost(1: AppendPostRPCRequest request)
    //修改岗位信息
    UpdatePostRPCResponse UpdatePost(1: UpdatePostRPCRequest request)
    //删除岗位
    DeletePostRPCResponse DeletePost(1: DeletePostRPCRequest request)
    //查询岗位信息-求职者界面
    GetsPostInUserRPCResponse GetsPostInUser(1: GetsPostInUserRPCRequest request)
    //查询岗位信息-HR界面
    GetsPostInHRRPCResponse GetsPostInHR(1: GetsPostInHRRPCRequest request)
    //投递岗位
    DeliveryPostRPCResponse DeliveryPost(1: DeliveryPostRPCRequest request)
    //岗位收藏
    CollectPostRPCResponse CollectPost(1: CollectPostRPCRequest request)
    //查看投递情况
    GetDeliveryPostRPCResponse GetDeliveryPost(1: GetDeliveryPostRPCRequest request)
    //查看收藏情况
    GetCollectPostRPCResponse GetCollectPost(1: GetCollectPostRPCRequest request)
    //HR修改投递的简历状态
    UpdateResumeStatusRPCResponse UpdateResumeStatus(1: UpdateResumeStatusRPCRequest request)

    //内部接口
    //查询某个岗位或人才库投递的用户id
    GetUserIdByPostIdRPCResponse GetUserIdByPostId(1: GetUserIdByPostIdRPCRequest request)
    //根据hrid删除岗位表、岗位投递表、岗位收藏表相关信息
    DeleteResumeRelativeInfoByHRIdRPCResponse DeleteResumeRelativeInfoByHRId(1: DeleteResumeRelativeInfoByHRIdRPCRequest request)
    //根据部门id删除岗位表、岗位投递表、岗位收藏表相关信息
    DeleteResumeRelativeInfoByDepartmentIdRPCResponse DeleteResumeRelativeInfoByDepartmentId(1: DeleteResumeRelativeInfoByDepartmentIdRPCRequest request)
    //根据岗位类别id列表删除岗位表、岗位投递表、岗位收藏表相关信息
    DeleteResumeRelativeInfoByPostCategoryIdListRPCResponse DeleteResumeRelativeInfoByPostCategoryIdList(1: DeleteResumeRelativeInfoByPostCategoryIdListRPCRequest request)
    //根据城市id删除岗位城市映射表相关信息
    DeletePostCityByCityIdRPCResposne DeletePostCityByCityId(1: DeletePostCityByCityIdRPCRequest request)
}