CREATE TABLE collect_post(
    id            BIGINT PRIMARY KEY COMMENT'主键id',
    user_id       BIGINT NOT NULL COMMENT'用户id',
    post_id       BIGINT NOT NULL COMMENT'岗位id',
    created_at    datetime(3) COMMENT'创建时间',
    updated_at    datetime(3) COMMENT'更新时间',
    deleted_at    datetime(3) COMMENT'删除时间',
    INDEX         idx_deleted_at(deleted_at) COMMENT'deleted_at字段对应的普通索引',
    -- INDEX idx_user_id(user_id) COMMENT'user_id字段对应的普通索引'
    UNIQUE uniq_user_post(user_id, post_id) COMMENT'user_id, post_id字段对应的联合唯一约束'
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT'用户收藏岗位表';