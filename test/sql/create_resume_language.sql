CREATE TABLE language(
    id                BIGINT PRIMARY KEY COMMENT'主键id',
    resume_id         BIGINT NOT NULL COMMENT'简历id',
    language_name          VARCHAR(16) NOT NULL COMMENT'语言',
    proficiency_level VARCHAR(8) NOT NULL COMMENT'精通程度',
    created_at        datetime(3) COMMENT'创建时间',
    updated_at        datetime(3) COMMENT'更新时间',
    deleted_at        datetime(3) COMMENT'删除时间',
    INDEX             idx_deleted_at(deleted_at) COMMENT'deleted_at字段对应的普通索引'
    -- INDEX idx_resume_id(resume_id) COMMENT'resume_id字段对应的普通索引',
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT'语言能力表';