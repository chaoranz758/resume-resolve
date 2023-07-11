CREATE TABLE talent_portrait(
    id                BIGINT PRIMARY KEY COMMENT'主键id',
    user_id           BIGINT NOT NULL COMMENT'用户id',
    age               TINYINT NOT NULL DEFAULT 0 COMMENT'年龄',
    max_education     VARCHAR(32) NOT NULL DEFAULT '' COMMENT'最高学历',
    graduated_school  VARCHAR(32) NOT NULL DEFAULT '' COMMENT'毕业院校',
    school_level      TINYINT NOT NULL DEFAULT 0 COMMENT'学校档次 1-985/211/双一流 2-211/双一流 3-双一流',
    working_seniority TINYINT NOT NULL DEFAULT 0 COMMENT'工作年限',
    created_at        datetime(3) COMMENT'创建时间',
    updated_at        datetime(3) COMMENT'更新时间',
    deleted_at        datetime(3) COMMENT'删除时间',
    INDEX             idx_deleted_at(deleted_at) COMMENT'deleted_at字段对应的普通索引'
    -- INDEX idx_resume_id(resume_id) COMMENT'resume_id字段对应的普通索引',
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT'用户画像表';