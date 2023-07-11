CREATE TABLE department(
    department_id          BIGINT PRIMARY KEY COMMENT'部门id',
    department_name        VARCHAR(32) NOT NULL UNIQUE COMMENT'部门名称',
    department_description VARCHAR(1024) NOT NULL COMMENT'部门描述信息',
    created_at    datetime(3) COMMENT'创建时间',
    updated_at    datetime(3) COMMENT'更新时间',
    deleted_at    datetime(3) COMMENT'删除时间',
    INDEX         idx_deleted_at(deleted_at) COMMENT'deleted_at字段对应的普通索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT'部门表';