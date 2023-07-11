CREATE TABLE post_category(
    post_category_id        BIGINT PRIMARY KEY COMMENT'岗位类别id',
    post_category_name      VARCHAR(32) NOT NULL UNIQUE COMMENT'岗位类别名称',
    post_category_level     tinyint NOT NULL DEFAULT 0 COMMENT'岗位级别 0-1级 1-2级',
    post_category_parent_id BIGINT NOT NULL DEFAULT 0 COMMENT'父岗位类别id',
    created_at              datetime(3) COMMENT'创建时间',
    updated_at              datetime(3) COMMENT'更新时间',
    deleted_at              datetime(3) COMMENT'删除时间',
    INDEX                   idx_deleted_at(deleted_at) COMMENT'deleted_at字段对应的普通索引'
    -- level;post_category_parent_id单独索引
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT'岗位类别表';