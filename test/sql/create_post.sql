CREATE TABLE post(
    post_id               BIGINT PRIMARY KEY COMMENT'岗位id',
    hr_id                 BIGINT NOT NULL COMMENT'发布岗位的HRid',
    post_brief            VARCHAR(64) NOT NULL COMMENT'岗位简介',
    post_description      VARCHAR(1024) NOT NULL COMMENT'岗位描述',
    post_require          VARCHAR(512) NOT NULL COMMENT'岗位要求',
    is_school_recruitment TINYINT NOT NULL DEFAULT 0 COMMENT'是否校招 1-非校招 2-校招',
    is_internship         TINYINT NOT NULL DEFAULT 0 COMMENT'是否实习 1-非实习 2-实习',
    post_category_id      BIGINT NOT NULL COMMENT'岗位类别id',
    department_id         BIGINT NOT NULL COMMENT'部门id',
    created_at            datetime(3) COMMENT'创建时间',
    updated_at            datetime(3) COMMENT'更新时间',
    deleted_at            datetime(3) COMMENT'删除时间',
    INDEX                 idx_deleted_at(deleted_at) COMMENT'deleted_at字段对应的普通索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT'岗位信息表';