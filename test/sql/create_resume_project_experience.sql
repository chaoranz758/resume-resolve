CREATE TABLE project_experience(
    id                  BIGINT PRIMARY KEY COMMENT'主键id',
    resume_id           BIGINT NOT NULL COMMENT'简历id',
    project_name        VARCHAR(32) NOT NULL COMMENT'项目名称',
    project_role        VARCHAR(32) NOT NULL COMMENT'项目角色',
    project_description VARCHAR(1024) NOT NULL COMMENT'项目描述',
    project_url         VARCHAR(256) NOT NULL DEFAULT '' COMMENT'项目描述',
    start_time          date NOT NULL COMMENT'起时间',
    end_time            date NOT NULL COMMENT'止时间',
    created_at          datetime(3) COMMENT'创建时间',
    updated_at          datetime(3) COMMENT'更新时间',
    deleted_at          datetime(3) COMMENT'删除时间',
    INDEX               idx_deleted_at(deleted_at) COMMENT'deleted_at字段对应的普通索引'
    -- INDEX idx_resume_id(resume_id) COMMENT'resume_id字段对应的普通索引',
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT'项目经历表';