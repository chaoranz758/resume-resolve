CREATE TABLE user(
    user_id       BIGINT PRIMARY KEY COMMENT'用户id',
    username      VARCHAR(16) NOT NULL UNIQUE COMMENT'用户名',
    password      VARCHAR(16) NOT NULL COMMENT'密码',
    role          TINYINT NOT NULL DEFAULT 0 COMMENT'0-普通用户 1-HR 2-管理员',
    department_id BIGINT NOT NULL DEFAULT 0 COMMENT'所属部门id',
    created_at    datetime(3) COMMENT'创建时间',
    updated_at    datetime(3) COMMENT'更新时间',
    deleted_at    datetime(3) COMMENT'删除时间',
    INDEX         idx_deleted_at(deleted_at) COMMENT'deleted_at字段对应的普通索引'
    -- 用户名和role的联合索引; create_at role; department_id create_at role;
    -- role的单独索引？
    -- INDEX idx_department_role(department_id, role) COMMENT'department_id, role字段对应的联合索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT'用户表';