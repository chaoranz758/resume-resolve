CREATE TABLE contest_certificate(
    id          BIGINT PRIMARY KEY COMMENT'主键id',
    resume_id   BIGINT NOT NULL COMMENT'简历id',
    name        VARCHAR(32) NOT NULL COMMENT'竞赛名称/证书名称',
    description VARCHAR(1024) NOT NULL COMMENT'描述',
    is_contest  TINYINT NOT NULL DEFAULT 0 COMMENT'0-竞赛 1-证书',
    created_at  datetime(3) COMMENT'创建时间',
    updated_at  datetime(3) COMMENT'更新时间',
    deleted_at  datetime(3) COMMENT'删除时间',
    INDEX       idx_deleted_at(deleted_at) COMMENT'deleted_at字段对应的普通索引'
    -- INDEX idx_resume_id(resume_id) COMMENT'resume_id字段对应的普通索引',
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT'竞赛表/证书表';