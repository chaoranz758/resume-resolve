namespace go post_category

include "../../base/post_category.thrift"

struct AppendPostCategoryRPCRequest {
    1: string post_category_name
    2: i8 post_category_level
    3: i64 post_category_parent_id
}

struct AppendPostCategoryRPCResponse {
    1: i32 code
    2: string message
    3: post_category.PostCategoryInformation data
}

struct UpdatePostCategoryRPCRequest {
    1: string post_category_name
    2: i64 post_category_id
}

struct UpdatePostCategoryRPCResponse {
    1: i32 code
    2: string message
    3: post_category.PostCategoryInformation data
}

struct DeletePostCategoryRPCRequest {
    1: i64 post_category_id
    2: i8 level
    3: i64 post_category_parent_id
}

struct DeletePostCategoryRPCResponse {
    1: i32 code
    2: string message
}

struct GetsPostCategoryRPCRequest {
    1: i8 level
    2: i64 post_category_id
}

struct GetsPostCategoryRPCResponse {
    1: i32 code
    2: string message
    3: list<post_category.PostCategoryAllInformation> data
}

struct GetPostCategoryByIdRPCRequest {
    1: list<i64> post_category_id
}

struct GetPostCategoryByIdRPCResponse {
    1: i32 code
    2: string message
    3: list<post_category.PostCategoryAllInformation> post_category_info_list
}

service PostCategoryRPCService {
    //外部接口
    //增加岗位类别
    AppendPostCategoryRPCResponse AppendPostCategory(1: AppendPostCategoryRPCRequest request)
    //修改岗位类别
    UpdatePostCategoryRPCResponse UpdatePostCategory(1: UpdatePostCategoryRPCRequest request)
    //删除岗位类别
    DeletePostCategoryRPCResponse DeletePostCategory(1: DeletePostCategoryRPCRequest request)
    //查询所有岗位类别
    GetsPostCategoryRPCResponse GetsPostCategory(1: GetsPostCategoryRPCRequest request)

    //内部接口
    //根据岗位类别id查一二级岗位名称
    GetPostCategoryByIdRPCResponse GetPostCategoryById(1: GetPostCategoryByIdRPCRequest request)
}