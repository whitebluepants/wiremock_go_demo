SELECT * FROM ccam_collections WHERE deleted = 0;

-- 1. 导出集合信息 collection_id = 14
select * from ccam_collections where collection_id = 14

-- 2. 导出该集合的 API 分组（ccam_api_tags）
SELECT * FROM ccam_api_tags WHERE collection_id = 14 AND deleted = 0;

-- 3. 导出该集合下的所有 API（ccam_api_info）
SELECT * FROM ccam_api_info WHERE collection_id = 14 AND deleted = 0 

-- 4. 导出该集合的 Schema 定义（ccam_collection_schemas）
SELECT * 
FROM ccam_collection_schemas 
WHERE collection_id = 14 
  AND deleted = 0;
  
  
-- 5. 导出 Schema 的属性（ccam_schema_properties）
SELECT * 
FROM ccam_schema_properties 
WHERE schema_id IN (
    SELECT schema_id 
    FROM ccam_collection_schemas 
    WHERE collection_id = 14 
      AND deleted = 0
) 
  AND deleted = 0;

-- 6. 导出 API 请求参数（ccam_api_request_params）
SELECT * 
FROM ccam_api_request_params 
WHERE collection_id = 14 
  AND api_id IN (
    SELECT api_id 
    FROM ccam_api_info 
    WHERE collection_id = 14 
      AND deleted = 0
  ) 
  AND deleted = 0;

-- 7. 导出 API 响应定义（ccam_api_responses）
SELECT * 
FROM ccam_api_responses 
WHERE api_id IN (
    SELECT api_id 
    FROM ccam_api_info 
    WHERE collection_id = 14 
      AND deleted = 0
) 
  AND deleted = 0;

-- 8. 导出响应参数（ccam_api_response_params）
SELECT * 
FROM ccam_api_response_params 
WHERE api_id IN (
    SELECT api_id 
    FROM ccam_api_info 
    WHERE collection_id = 14 
      AND deleted = 0
) 
  AND deleted = 0;

-- 9. 导出 Mock API 基础信息（ccam_mock_api）
SELECT * 
FROM ccam_mock_api 
WHERE api_id IN (
    SELECT api_id 
    FROM ccam_api_info 
    WHERE collection_id = 14 
      AND deleted = 0
) 
  AND deleted = 0;

-- 10. 导出 Mock 规则（ccam_mock_rule）
SELECT * 
FROM ccam_mock_rule 
WHERE api_id IN (
    SELECT api_id 
    FROM ccam_api_info 
    WHERE collection_id = 14 
      AND deleted = 0
) 
  AND deleted = 0;

-- 11. 导出 Mock 脚本（ccam_mock_script）
SELECT * 
FROM ccam_mock_script 
WHERE api_id IN (
    SELECT api_id 
    FROM ccam_api_info 
    WHERE collection_id = 14 
      AND deleted = 0
) 
  AND deleted = 0;