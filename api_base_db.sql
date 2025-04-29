-- 创建数据库（已补充）
-- CREATE DATABASE IF NOT EXISTS `capability-backend` 
-- CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `capability-backend`;

-- ----------------------------
-- Table structure for ccam_collections
-- ----------------------------
CREATE TABLE `ccam_collections` (
    `collection_id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Unique ID',
    `fix_name` VARCHAR(100) NOT NULL COMMENT 'ApplicationName CMDB唯一标识',
    `component_name` VARCHAR(100) NOT NULL COMMENT 'Eureka注册的微服务名称',
    `component_cluster` VARCHAR(100) NOT NULL COMMENT 'Eureka注册的微服务所属集群',
    `collection_name` VARCHAR(100) NOT NULL COMMENT 'API集合名称',
    `collection_description` VARCHAR(1000) COMMENT 'API集合描述',
    `collection_type` VARCHAR(20) NOT NULL COMMENT 'API集合类型，例如openapi',
    `collection_version` VARCHAR(50) NOT NULL COMMENT 'API集合版本，默认main', -- 修正默认值约束
    -- 通用字段
    `create_user` VARCHAR(256) NOT NULL COMMENT '创建人',
    `create_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_user` VARCHAR(256) COMMENT '更新人',
    `update_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted` INTEGER DEFAULT 0 COMMENT '软删除，1表示删除，0表示存在',
    -- 唯一约束
    CONSTRAINT `uq_fixname_componentname_cluster_version` UNIQUE 
    (`fix_name`, `component_name`, `component_cluster`, `collection_version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='API集合表，基本等于微服务维度';

-- ----------------------------
-- Table structure for ccam_collection_schemas
-- ----------------------------
CREATE TABLE `ccam_collection_schemas` (
    `schema_id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Unique ID',
    `collection_id` BIGINT NOT NULL COMMENT '关联的 cc_api_collections 表的 ID',
    `schema_key` VARCHAR(100) NOT NULL COMMENT '结构体标识，例如：User',
    `schema_type` VARCHAR(100) COMMENT '结构体类型，例如：object',
    `schema_description` VARCHAR(100) COMMENT '结构体描述，例如：用户请求对象',
    -- 通用字段
    `create_user` VARCHAR(256) NOT NULL COMMENT '创建人',
    `create_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_user` VARCHAR(256) COMMENT '更新人',
    `update_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted` INTEGER DEFAULT 0 COMMENT '软删除，1 表示删除，0 表示存在',
    -- 索引
    INDEX `idx_collection_id_schema_key` (`collection_id`, `schema_key`),
    -- 唯一约束
    CONSTRAINT `uq_collection_id_schema_key` UNIQUE (`collection_id`, `schema_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='API 通用结构体表';

-- ----------------------------
-- Table structure for ccam_schema_properties
-- ----------------------------
CREATE TABLE `ccam_schema_properties` (
    `property_id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'Unique ID',
    `schema_id` BIGINT NOT NULL COMMENT '关联的 ccam_collection_schemas 表的 ID',
    `collection_id` BIGINT NOT NULL COMMENT '冗余字段，关联 ccam_collections 表的 ID',
    `property_name` VARCHAR(50) NOT NULL COMMENT '属性名称',
    `property_type` VARCHAR(20) NOT NULL COMMENT '属性类型，如：integer、string 或 ref',
    `property_format` VARCHAR(50) DEFAULT NULL COMMENT '属性格式，如：int64、INTERNAL',
    `required` TINYINT(1) DEFAULT '1' COMMENT '是否必须，true/false，默认 true',
    `property_description` VARCHAR(500) COMMENT '属性描述',
    `property_enum` VARCHAR(2000) DEFAULT NULL COMMENT '枚举列表，property_type=string 且 enum 不为空时需设值',
    `schema_key` VARCHAR(100) DEFAULT NULL COMMENT '等于 Swagger 中的 simple_ref，property_type=ref 时不可为空',
    `original_ref` VARCHAR(100) DEFAULT NULL COMMENT '完整的引用描述，property_type=ref 时不可为空',
    `ref_schema_id` BIGINT DEFAULT NULL COMMENT '引用 ID，来源于 ccam_collection_schemas 表的 schema_id，property_type=ref 时不可为空',
    -- 通用字段
    `create_user` VARCHAR(256) NOT NULL COMMENT '创建人',
    `create_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_user` VARCHAR(256) DEFAULT NULL COMMENT '更新人',
    `update_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted` INT DEFAULT '0' COMMENT '软删除，1 表示删除，0 表示存在',
    PRIMARY KEY (`property_id`),
    KEY `ccam_schema_properties_collection_id_IDX` (`collection_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3831 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='属性表';

-- ----------------------------
-- Table structure for ccam_api_tags
-- ----------------------------
CREATE TABLE `ccam_api_tags` (
    `tag_id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Unique ID',
    `collection_id` BIGINT NOT NULL COMMENT '关联的 cc_api_collections 表的 ID',
    `tag_name` VARCHAR(100) NOT NULL COMMENT 'API 分组名称',
    `tag_description` VARCHAR(200) COMMENT 'API 分组描述',
    -- 通用字段
    `create_user` VARCHAR(256) NOT NULL COMMENT '创建人',
    `create_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_user` VARCHAR(256) COMMENT '更新人',
    `update_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted` INTEGER DEFAULT 0 COMMENT '软删除，1 表示删除，0 表示存在'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='API 分组表，记录 Controller';

-- ----------------------------
-- Table structure for ccam_api_info
-- ----------------------------
CREATE TABLE `ccam_api_info` (
    `api_id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Unique ID',
    `tag_id` BIGINT NOT NULL COMMENT '关联的 cc_api_tags 表的 ID',
    `collection_id` BIGINT NOT NULL COMMENT '冗余字段，关联 cc_api_collections 表的 ID',
    `api_summary` VARCHAR(1000) NOT NULL COMMENT 'API 名称',
    `operation_id` VARCHAR(100) NOT NULL COMMENT '操作标识',
    `api_description` VARCHAR(1000) COMMENT 'API 描述',
    `api_method` VARCHAR(50) NOT NULL COMMENT '请求方式（如：GET、POST 等）',
    `api_path` VARCHAR(255) NOT NULL COMMENT '请求路径，例如：/qotho/api/sync',
    `deprecated` BOOLEAN DEFAULT FALSE COMMENT '是否过期，过期置灰不建议使用，默认false',
    -- 通用字段
    `create_user` VARCHAR(256) NOT NULL COMMENT '创建人',
    `create_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_user` VARCHAR(256) COMMENT '更新人',
    `update_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted` INTEGER DEFAULT 0 COMMENT '软删除，1 表示删除，0 表示存在',
    -- 唯一约束
    CONSTRAINT `uq_collection_id_api_method_api_path` UNIQUE 
    (`collection_id`, `api_method`, `api_path`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='API 基础信息表';

-- ----------------------------
-- Table structure for ccam_api_request_params
-- ----------------------------
CREATE TABLE `ccam_api_request_params` (
    `request_param_id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Unique ID',
    `api_id` BIGINT NOT NULL COMMENT '关联的 cc_api_info 表的 API ID',
    `collection_id` BIGINT NOT NULL COMMENT '关联的 ccam_collection 表的 ID',
    `param_in` VARCHAR(100) NOT NULL COMMENT '参数所属类别，query/body/path/header',
    `content_type` VARCHAR(200) DEFAULT '*/*' COMMENT '请求类型，默认为 * /*，如 application/json',
    `param_name` VARCHAR(200) NOT NULL COMMENT '参数名称',
    `required` BOOLEAN DEFAULT TRUE COMMENT '是否必须，true/false，默认 true',
    `param_type` VARCHAR(20) COMMENT '参数类型，如：integer、string',
    `param_description` VARCHAR(500) COMMENT '参数描述',
    `param_format` VARCHAR(50) COMMENT '参数格式，如：int64',
    `schema_key` VARCHAR(100) COMMENT '等于 Swagger 中的 simple_ref，param_in=body 时不可为空',
    `original_ref` VARCHAR(100) COMMENT '完整的引用描述，param_in=body 且 param_type=ref 时不可为空',
    `ref_schema_id` BIGINT COMMENT '引用 ID，来源于 ccam_collection_schemas 表的 schema_id，param_in=body 且 param_type=ref 时不可为空',
    -- 通用字段
    `create_user` VARCHAR(256) NOT NULL COMMENT '创建人',
    `create_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_user` VARCHAR(256) COMMENT '更新人',
    `update_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted` INTEGER DEFAULT 0 COMMENT '软删除，1 表示删除，0 表示存在'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='API 请求参数表，记录 Swagger 的 parameters 列表';

-- ----------------------------
-- Table structure for ccam_api_responses
-- ----------------------------
CREATE TABLE `ccam_api_responses` (
    `response_id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Unique ID',
    `api_id` BIGINT NOT NULL COMMENT '关联的 cc_api_info 表的 API ID',
    `response_code` VARCHAR(100) NOT NULL COMMENT '响应 Code，如：200、400、401',
    `response_description` VARCHAR(500) COMMENT '响应描述',
    `content_type` VARCHAR(200) DEFAULT '*/*' COMMENT '响应类型，默认 * /*，如 application/json',
    -- 通用字段
    `create_user` VARCHAR(256) NOT NULL COMMENT '创建人',
    `create_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_user` VARCHAR(256) COMMENT '更新人',
    `update_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted` INTEGER DEFAULT 0 COMMENT '软删除，1 表示删除，0 表示存在'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='API 响应信息表，记录 Swagger 的 responses';

-- ----------------------------
-- Table structure for ccam_api_response_params
-- ----------------------------
CREATE TABLE `ccam_api_response_params` (
    `response_param_id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Unique ID',
    `response_id` BIGINT NOT NULL COMMENT '关联的 ccam_api_responses 表的 ID',
    `api_id` BIGINT NOT NULL COMMENT '冗余字段，关联 cc_api_info 表的 API ID',
    `collection_id` BIGINT NOT NULL COMMENT '关联的 ccam_collection 表的 ID',
    `param_name` VARCHAR(200) NOT NULL COMMENT '参数名称',
    `param_type` VARCHAR(20) NOT NULL COMMENT '参数类型，如：integer、string、ref',
    `param_format` VARCHAR(50) COMMENT '参数格式，如：int64',
    `param_description` VARCHAR(500) COMMENT '参数描述',
    `schema_key` VARCHAR(100) COMMENT '等于 Swagger 中的 simple_ref，param_type=ref 时不可为空',
    `original_ref` VARCHAR(100) COMMENT '完整的引用描述，param_type=ref 时不可为空',
    `ref_schema_id` BIGINT COMMENT '引用 ID，来源于 ccam_collection_schemas 表的 schema_id，param_type=ref 时不可为空',
    -- 通用字段
    `create_user` VARCHAR(256) NOT NULL COMMENT '创建人',
    `create_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_user` VARCHAR(256) COMMENT '更新人',
    `update_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted` INTEGER DEFAULT 0 COMMENT '软删除，1 表示删除，0 表示存在'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='API 响应参数表，记录 Swagger 的 response 参数列表';

-- ----------------------------
-- Index for ccam_api_responses
-- ----------------------------
CREATE INDEX `ccam_api_responses_response_code_IDX` 
ON `ccam_api_responses` (`response_code`);

-- ----------------------------
-- Alter table for ccam_api_info (add api_remark)
-- ----------------------------
ALTER TABLE `ccam_api_info` 
ADD `api_remark` VARCHAR(200) NULL COMMENT '用户增加的API额外描述';

-- ----------------------------
-- Alter table for ccam_collections (modify collection_version)
-- ----------------------------
ALTER TABLE `ccam_collections` 
MODIFY COLUMN `collection_version` VARCHAR(50) NOT NULL 
COMMENT 'API集合版本';

-- ----------------------------
-- Update collection_type from "openapi" to "Swagger2"
-- ----------------------------
UPDATE `ccam_collections` 
SET `collection_type` = 'Swagger2' 
WHERE `collection_type` = 'openapi';