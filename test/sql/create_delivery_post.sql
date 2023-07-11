CREATE TABLE delivery_post(
    id            BIGINT PRIMARY KEY COMMENT'主键id',
    user_id       BIGINT NOT NULL COMMENT'用户id',
    post_id       BIGINT NOT NULL COMMENT'岗位id',
    resume_status TINYINT NOT NULL DEFAULT 0 COMMENT'简历投递状态 0-未处理 1-通过 2-不通过',
    is_talent_pool TINYINT NOT NULL DEFAULT 0 COMMENT'是否位于人才库 0-不在 1-在',
    created_at    datetime(3) COMMENT'创建时间',
    updated_at    datetime(3) COMMENT'更新时间',
    deleted_at    datetime(3) COMMENT'删除时间',
    INDEX         idx_deleted_at(deleted_at) COMMENT'deleted_at字段对应的普通索引',
    -- INDEX idx_user_id(user_id) COMMENT'user_id字段对应的普通索引',
    -- INDEX idx_post_order_status_pool(post_id, resume_order, resume_status, is_talent_pool) COMMENT'post_id, resume_order, resume_status, is_talent_pool字段对应的联合索引'
    UNIQUE uniq_user_post(user_id, post_id) COMMENT'user_id, post_id字段对应的联合唯一约束'
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT'用户投递岗位表';