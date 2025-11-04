DROP SCHEMA IF EXISTS `gorm`;
CREATE SCHEMA IF NOT EXISTS `gorm`;
USE `gorm`;

CREATE TABLE `station`
(
    `id`           bigint unsigned                            NOT NULL,
    `longitude`    decimal(23, 20)  DEFAULT -1                NOT NULL COMMENT '经度',
    `latitude`     decimal(23, 20)  DEFAULT -1                NOT NULL COMMENT '纬度',
    `name`         varchar(32)      DEFAULT ''                NOT NULL COMMENT '测站名称',
    `type`         tinyint unsigned DEFAULT 0                 NOT NULL COMMENT '站类(1-流量，2-雨量，3-河道水位，4-水库水位)',
    `code`         varchar(32)      DEFAULT ''                NOT NULL COMMENT '测站编码',
    PRIMARY KEY (`id`),
    CONSTRAINT UNIQUE (`code`, `type`)
)
    COMMENT '测站信息';



CREATE TABLE `station_river_water_level`
(
    `station_id`            bigint unsigned                         NOT NULL,
    `alarm_water_level`     decimal(8, 3) DEFAULT -1                NOT NULL COMMENT '警戒水位，左开右闭',
    `guarantee_water_level` decimal(8, 3) DEFAULT -1                NOT NULL COMMENT '保证水位，左开右闭',
    PRIMARY KEY (`station_id`),
    CONSTRAINT FOREIGN KEY (`station_id`) REFERENCES `station` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
)
    COMMENT '水位测站扩展信息';

CREATE TABLE `station_rain`
(
    `station_id` bigint unsigned                         NOT NULL,
    `level1`     decimal(6, 2) DEFAULT 30                NOT NULL COMMENT '一级警戒雨量阈值[0-lv1)',
    `level2`     decimal(6, 2) DEFAULT 50                NOT NULL COMMENT '一级警戒雨量阈值[lv1-lv2)',
    PRIMARY KEY (`station_id`),
    CONSTRAINT FOREIGN KEY (`station_id`) REFERENCES `station` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
)
    COMMENT '雨量测站扩展信息';

CREATE TABLE `reservoir`
(
    `id`                      bigint unsigned                            NOT NULL,
    `wl_station_id`           bigint unsigned  DEFAULT 0                 NOT NULL COMMENT '水库水位测站id'
        REFERENCES `station` (`id`),
    `rain_station_id`         bigint unsigned  DEFAULT 0                 NOT NULL COMMENT '水库雨量测站id'
        REFERENCES `station_river_water_level` (`station_id`),
    `name`                    varchar(32)      DEFAULT ''                NOT NULL COMMENT '水库名称',
    `longitude`               decimal(23, 20)  DEFAULT -1                NOT NULL COMMENT '地理位置经度坐标',
    `latitude`                decimal(23, 20)  DEFAULT -1                NOT NULL COMMENT '地理位置纬度坐标',
    PRIMARY KEY (`id`)
)
    COMMENT '水库信息';
