use gorm;

CREATE TABLE `user_follow_0`
(
    `id`           bigint NOT NULL,
    `user_id`      bigint NOT NULL DEFAULT '0' COMMENT '发起关注的人',
    `followed_uid` bigint NOT NULL DEFAULT '0' COMMENT '被关注用户的uid',
    `created_at`   datetime        DEFAULT NULL,
    `updated_at`   datetime        DEFAULT NULL,
    `deleted_at`   datetime        DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `udx_uid_fuid` (`user_id`, `followed_uid`)
) COMMENT ='用户关注表';

CREATE TABLE `user_follow_1`
(
    `id`           bigint NOT NULL,
    `user_id`      bigint NOT NULL DEFAULT '0' COMMENT '发起关注的人',
    `followed_uid` bigint NOT NULL DEFAULT '0' COMMENT '被关注用户的uid',
    `created_at`   datetime        DEFAULT NULL,
    `updated_at`   datetime        DEFAULT NULL,
    `deleted_at`   datetime        DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `udx_uid_fuid` (`user_id`, `followed_uid`)
) COMMENT ='用户关注表';

INSERT INTO `user_follow_0` (`created_at`,`updated_at`,`deleted_at`,`user_id`,`followed_uid`) VALUES ('2023-02-19 20:52:26.939','2023-02-19 20:52:26.939',NULL,79535114761158382,79542521267813869)