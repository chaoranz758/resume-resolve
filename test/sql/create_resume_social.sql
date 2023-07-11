CREATE TABLE social(
    id              BIGINT PRIMARY KEY COMMENT'主键id',
    resume_id       BIGINT NOT NULL COMMENT'简历id',
    social_platform VARCHAR(256) NOT NULL COMMENT'社交平台',
    platform_url    VARCHAR(256) NOT NULL COMMENT'url/id',
    created_at      datetime(3) COMMENT'创建时间',
    updated_at      datetime(3) COMMENT'更新时间',
    deleted_at      datetime(3) COMMENT'删除时间',
    INDEX           idx_deleted_at(deleted_at) COMMENT'deleted_at字段对应的普通索引'
    -- INDEX idx_resume_id(resume_id) COMMENT'resume_id字段对应的普通索引',
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT'社交账号表';