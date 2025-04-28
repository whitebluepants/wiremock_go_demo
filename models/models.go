package models

import (
	"wiremock_go_demo/config"
)

// CcamApiResponse 对应ccam_api_responses表结构
type CcamApiResponse struct {
	ResponseID          int64  `gorm:"primaryKey;column:response_id"`
	ApiID               int64  `gorm:"column:api_id"`
	ResponseCode        string `gorm:"column:response_code"`
	ResponseDescription string `gorm:"column:response_description"`
	ContentType         string `gorm:"column:content_type"`
	CreateUser          string `gorm:"column:create_user"`
	CreateDate          string `gorm:"column:create_date"`
	UpdateUser          string `gorm:"column:update_user"`
	UpdateDate          string `gorm:"column:update_date"`
	Deleted             int    `gorm:"column:deleted"`
}

// CcamApiResponseParam 对应ccam_api_response_params表结构
type CcamApiResponseParam struct {
	ResponseParamID  int64  `gorm:"primaryKey;column:response_param_id"`
	ResponseID       int64  `gorm:"column:response_id"`
	ApiID            int64  `gorm:"column:api_id"`
	CollectionID     int64  `gorm:"column:collection_id"`
	ParamName        string `gorm:"column:param_name"`
	ParamType        string `gorm:"column:param_type"`
	ParamFormat      string `gorm:"column:param_format"`
	ParamDescription string `gorm:"column:param_description"`
	SchemaKey        string `gorm:"column:schema_key"`
	OriginalRef      string `gorm:"column:original_ref"`
	RefSchemaID      int64  `gorm:"column:ref_schema_id"`
	CreateUser       string `gorm:"column:create_user"`
	CreateDate       string `gorm:"column:create_date"`
	UpdateUser       string `gorm:"column:update_user"`
	UpdateDate       string `gorm:"column:update_date"`
	Deleted          int    `gorm:"column:deleted"`
}

// CcamSchemaProperty 对应ccam_schema_properties表结构
type CcamSchemaProperty struct {
	PropertyID          int64  `gorm:"primaryKey;column:property_id"`
	SchemaID            int64  `gorm:"column:schema_id"`
	CollectionID        int64  `gorm:"column:collection_id"`
	PropertyName        string `gorm:"column:property_name"`
	PropertyType        string `gorm:"column:property_type"`
	PropertyFormat      string `gorm:"column:property_format"`
	Required            int    `gorm:"column:required"`
	PropertyDescription string `gorm:"column:property_description"`
	PropertyEnum        string `gorm:"column:property_enum"`
	SchemaKey           string `gorm:"column:schema_key"`
	OriginalRef         string `gorm:"column:original_ref"`
	RefSchemaID         int64  `gorm:"column:ref_schema_id"`
	CreateUser          string `gorm:"column:create_user"`
	CreateDate          string `gorm:"column:create_date"`
	UpdateUser          string `gorm:"column:update_user"`
	UpdateDate          string `gorm:"column:update_date"`
	Deleted             int    `gorm:"column:deleted"`
}

// 查询函数
func GetCcamApiResponsesByApiID(apiID int64) ([]CcamApiResponse, error) {
	var responses []CcamApiResponse
	err := config.DB.Where("api_id = ? AND deleted = 0", apiID).Find(&responses).Error
	return responses, err
}

func GetCcamApiResponseParamsByResponseID(responseID int64) ([]CcamApiResponseParam, error) {
	var params []CcamApiResponseParam
	err := config.DB.Where("response_id = ? AND deleted = 0", responseID).Find(&params).Error
	return params, err
}

func GetCcamSchemaPropertiesBySchemaID(schemaID int64) ([]CcamSchemaProperty, error) {
	var properties []CcamSchemaProperty
	err := config.DB.Where("schema_id = ? AND deleted = 0", schemaID).Find(&properties).Error
	return properties, err
}

// CcamApiInfo 对应ccam_api_info表结构
type CcamApiInfo struct {
	ApiID       int64  `gorm:"column:api_id;primaryKey"`
	ApiMethod   string `gorm:"column:api_method"`
	ApiPath     string `gorm:"column:api_path"`
	Description string `gorm:"column:description"`
}

func (s *CcamApiInfo) TableName() string {
	return "ccam_api_info" // 指定表名
}

// GetCcamApiInfoByApiID 根据API ID获取接口基本信息
func GetCcamApiInfoByApiID(apiID int64) (*CcamApiInfo, error) {
	var apiInfo CcamApiInfo
	result := config.DB.Where("api_id = ?", apiID).First(&apiInfo)
	return &apiInfo, result.Error
}
