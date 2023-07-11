CREATE TABLE post_city_map(
    id            BIGINT PRIMARY KEY COMMENT'主键id',
    post_id       BIGINT NOT NULL COMMENT'部门id',
    city_id       BIGINT NOT NULL COMMENT'城市id',
    created_at    datetime(3) COMMENT'创建时间',
    updated_at    datetime(3) COMMENT'更新时间',
    deleted_at    datetime(3) COMMENT'删除时间',
    INDEX         idx_deleted_at(deleted_at) COMMENT'deleted_at字段对应的普通索引',
    -- INDEX idx_post_id(post_id) COMMENT'post_id字段对应的普通索引',
    -- INDEX idx_city_id(city_id) COMMENT'city_id字段对应的普通索引',
    UNIQUE uniq_post_city(post_id, city_id) COMMENT'post_id, city_id字段对应的联合唯一约束'
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT'岗位城市映射表';