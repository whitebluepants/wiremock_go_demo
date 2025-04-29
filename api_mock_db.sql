-- 创建数据库（已补充）
-- CREATE DATABASE IF NOT EXISTS `capability-backend` 
-- CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `capability-backend`;

-- ----------------------------
-- Table structure for ccam_mock_api (Mock基础表)
-- ----------------------------
CREATE TABLE IF NOT EXISTS `ccam_mock_api` (
    `mock_api_id` BIGINT AUTO_INCREMENT COMMENT 'ID' PRIMARY KEY,
    `api_id` BIGINT NOT NULL COMMENT '关联ccam_api_info.api_id',
    `api_description` VARCHAR(200) NULL COMMENT 'API描述，冗余字段',
    `api_path` VARCHAR(255) NOT NULL COMMENT '请求路径，如/qotho/api/sync，冗余字段',
    `api_method` VARCHAR(50) NOT NULL COMMENT '请求方式，冗余字段',
    `mock_type` TINYINT(1) DEFAULT 0 COMMENT 'Mock类型，0:默认Mock，1:期望Mock，如果是1则要查rule表',
    `mock_enabled` TINYINT(1) DEFAULT 1 COMMENT '是否启用Mock，0:禁用，1:启用',
    `default_response` TEXT NULL COMMENT '默认响应内容',
    `deprecated` TINYINT(1) DEFAULT 0 COMMENT '是否过期，过期置灰不建议使用',
    `create_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL COMMENT '创建时间',
    `update_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    CONSTRAINT `unique_api_path_method` UNIQUE (`api_path`, `api_method`),
    CONSTRAINT `ccam_mock_api_ibfk_1` FOREIGN KEY (`api_id`) 
    REFERENCES `ccam_api_info` (`api_id`)
) COMMENT 'Mock API基础表';

-- ----------------------------
-- Index for ccam_mock_api
-- ----------------------------
CREATE INDEX `api_id` ON `ccam_mock_api` (`api_id`);

-- ----------------------------
-- Table structure for ccam_mock_rule (Mock规则表)
-- ----------------------------
CREATE TABLE IF NOT EXISTS `ccam_mock_rule` (
    `rule_id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '规则ID',
    `mock_api_id` BIGINT NOT NULL COMMENT '关联ccam_mock_api.mock_api_id',
    `condition_type` VARCHAR(50) NOT NULL COMMENT '条件类型: header/query/body',
    `condition_key` VARCHAR(100) NOT NULL COMMENT '条件键（如X-Token）',
    `condition_value` VARCHAR(255) NOT NULL COMMENT '条件值（如Beaver abc123）',
    `priority` INT DEFAULT 1 COMMENT '规则优先级',
    `response_template` TEXT NULL COMMENT 'Handlebars模板内容',
    `status_code` INT NOT NULL COMMENT 'HTTP状态码（如200/400）',
    `delay_ms` INT DEFAULT 0 COMMENT '模拟延迟（毫秒）',
    `create_user` VARCHAR(256) NOT NULL,
    `create_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    `update_user` VARCHAR(256) NULL,
    `update_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted` INT DEFAULT 0 COMMENT '软删除标记，1表示已删除，0表示未删除'
) COMMENT 'Mock规则表';

-- ----------------------------
-- Table structure for ccam_mock_script (Mock脚本表)
-- ----------------------------
CREATE TABLE IF NOT EXISTS `ccam_mock_script` (
    `script_id` BIGINT AUTO_INCREMENT COMMENT '脚本唯一ID' PRIMARY KEY,
    `mock_api_id` BIGINT NOT NULL COMMENT '关联的API ID，指向 ccam_mock_api 表中的 ID', -- 修正外键描述（原pt_wire_mock_api应为ccam_mock_api）
    `script` TEXT COLLATE utf8mb4_unicode_ci NULL COMMENT 'Mock脚本内容，存储用于动态生成响应的脚本',
    `create_user` VARCHAR(256) NOT NULL COMMENT '创建人，记录创建脚本的用户',
    `create_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL COMMENT '创建时间，记录脚本的创建时间',
    `update_user` VARCHAR(256) NULL COMMENT '更新人，记录最后更新脚本的用户',
    `update_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间，记录脚本的最后更新时间'
) COMMENT 'Mock脚本表，用于存储与API相关的动态脚本';