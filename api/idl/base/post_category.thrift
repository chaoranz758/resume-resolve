namespace go post_category

struct PostCategoryInformation {
    1: i64 post_category_id
    2: string post_category_name
}

struct PostCategoryAllInformation {
    1: i64 post_category_id
    2: string post_category_name
    3: i64 post_category_parent_id
    4: string post_category_parent_name
}


