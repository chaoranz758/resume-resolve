namespace go post_category

include "../../base/common.thrift"

struct AppendPostCategoryRequest {
    1: string post_category_name (api.body = "post_category_name")
    2: i8 post_category_level (api.body = "post_category_level")
    3: i64 post_category_parent_id (api.body = "post_category_parent_id")
}

struct UpdatePostCategoryRequest {
    1: string post_category_name (api.body = "post_category_name")
    2: i64 post_category_id (api.body = "post_category_id")
}

struct DeletePostCategoryRequest {
    1: i64 post_category_id (api.body = "post_category_id")
    2: i8 level (api.body = "level")
    3: i64 post_category_parent_id (api.body = "post_category_parent_id")
}

struct GetsPostCategoryRequest {
    1: i8 level (api.query = "level")
    2: i64 post_category_id (api.query = "post_category_id")
}

service PostCategoryService {
    //增加岗位类别
    common.NilResponse AppendPostCategory(1: AppendPostCategoryRequest request) (api.post = "/api/v1/post-category/append")
    //修改岗位类别
    common.NilResponse UpdatePostCategory(1: UpdatePostCategoryRequest request) (api.post = "/api/v1/post-category/update")
    //删除岗位类别
    common.NilResponse DeletePostCategory(1: DeletePostCategoryRequest request) (api.post = "/api/v1/post-category/delete")
    //查询所有岗位类别
    common.NilResponse GetsPostCategory(1: GetsPostCategoryRequest request) (api.get = "/api/v1/post-category/gets")
}