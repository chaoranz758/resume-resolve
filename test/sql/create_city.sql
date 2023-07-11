CREATE TABLE city(
    city_id   BIGINT PRIMARY KEY COMMENT'城市id',
    city_name VARCHAR(16) NOT NULL UNIQUE COMMENT'城市名称',
    created_at    datetime(3) COMMENT'创建时间',
    updated_at    datetime(3) COMMENT'更新时间',
    deleted_at    datetime(3) COMMENT'删除时间',
    INDEX         idx_deleted_at(deleted_at) COMMENT'deleted_at字段对应的普通索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT'城市表';