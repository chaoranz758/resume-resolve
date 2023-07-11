CREATE TABLE educational_experience(
     id         BIGINT PRIMARY KEY COMMENT'主键id',
     resume_id  BIGINT NOT NULL COMMENT'简历id',
     school     VARCHAR(32) NOT NULL COMMENT'学校',
     education  VARCHAR(16) NOT NULL COMMENT'学历',
     speciality VARCHAR(32) NOT NULL COMMENT'专业',
     ranking    VARCHAR(16) NOT NULL COMMENT'成绩排名',
     start_time date NOT NULL COMMENT'起时间',
     end_time   date NOT NULL COMMENT'止时间',
     created_at datetime(3) COMMENT'创建时间',
     updated_at datetime(3) COMMENT'更新时间',
     deleted_at datetime(3) COMMENT'删除时间',
     INDEX      idx_deleted_at(deleted_at) COMMENT'deleted_at字段对应的普通索引'
    -- INDEX idx_resume_id(resume_id) COMMENT'resume_id字段对应的普通索引',
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT'教育经历表';