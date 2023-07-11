CREATE TABLE basic_info(
    resume_id       BIGINT PRIMARY KEY COMMENT'简历id',
    user_id         BIGINT NOT NULL UNIQUE COMMENT'用户id',
    name            VARCHAR(8) NOT NULL COMMENT'姓名',
    phone VARCHAR(20) NOT NULL COMMENT'电话号码',
    resume_url      VARCHAR(256) NOT NULL COMMENT'简历文件url',
    email           VARCHAR(32) NOT NULL COMMENT'邮箱',
    self_evaluation VARCHAR(1024) NOT NULL DEFAULT '' COMMENT'自我评价',
    birthday        date NOT NULL COMMENT'出生日期',
    created_at      datetime(3) COMMENT'创建时间',
    updated_at      datetime(3) COMMENT'更新时间',
    deleted_at      datetime(3) COMMENT'删除时间',
    INDEX           idx_deleted_at(deleted_at) COMMENT'deleted_at字段对应的普通索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT'简历基础信息表';