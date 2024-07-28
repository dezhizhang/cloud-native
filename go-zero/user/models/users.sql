CREATE TABLE users
(
    id        bigint AUTO_INCREMENT,
    name      varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
    password  varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
    mobile    varchar(255) NOT NULL DEFAULT '' COMMENT '手机号',
    gender    char(10)     NOT NULL DEFAULT 'male' COMMENT 'gender,male|female|unknown',
    create_at timestamp NULL,
    update_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE mobile_index (mobile),
    UNIQUE name_index (name),
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'users table';