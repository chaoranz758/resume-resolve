namespace go post

include "post_category.thrift"
include "department.thrift"

struct PostInfo {
    1: i64 post_id
    2: string post_brief
    3: string post_description
    4: string post_require
    5: byte is_school_recruitment
    6: byte is_internship
    7: post_category.PostCategoryAllInformation post_category_information
    8: list<department.CityInformation> city_information
    9: department.DepartmentInformation department_information
}



