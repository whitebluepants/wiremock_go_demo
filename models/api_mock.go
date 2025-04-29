package models

// CcamMockApi 对应ccam_mock_api表结构
type CcamMockApi struct {
	MockApiID       int64  `gorm:"primaryKey;column:mock_api_id"`
	ApiID           int64  `gorm:"column:api_id"`
	ApiDescription  string `gorm:"column:api_description"`
	ApiPath         string `gorm:"column:api_path"`
	ApiMethod       string `gorm:"column:api_method"`
	MockType        int    `gorm:"column:mock_type"`
	MockEnabled     int    `gorm:"column:mock_enabled"`
	DefaultResponse string `gorm:"column:default_response"`
	Deprecated      int    `gorm:"column:deprecated"`
	CreateTime      string `gorm:"column:create_time"`
	UpdateTime      string `gorm:"column:update_time"`
}

func (s *CcamMockApi) TableName() string {
	return "ccam_mock_api"
}

// GetExpectedMockApiRecords 根据API ID捞取mock期望的api记录
func GetExpectedMockApi(apiID int64) ([]CcamMockApi, error) {
	return []CcamMockApi{
		{
			// 正常Api Path情况
			MockApiID:       1,
			ApiID:           apiID,
			ApiDescription:  "Mock API Description",
			ApiPath:         "/mock/api/path",
			ApiMethod:       "GET",
			MockType:        1,
			MockEnabled:     1,
			DefaultResponse: `{\"message\": \"Mock response\"}`,
			Deprecated:      0,
			CreateTime:      "2024-01-01 00:00:00",
			UpdateTime:      "2024-01-01 00:00:00",
		},
		// 占位符Path 情况
		// {
		// 	MockApiID:       2,
		// 	ApiID:           apiID,
		// 	ApiDescription:  "Mock API Description 2",
		// 	ApiPath:         "/mock/api/{path_test}/path2",
		// 	ApiMethod:       "POST",
		// 	MockType:        1,
		// 	MockEnabled:     1,
		// 	DefaultResponse: `{\"message\": \"Mock response 2\"}`,
		// 	Deprecated:      0,
		// 	CreateTime:      "2024-01-02 00:00:00",
		// 	UpdateTime:      "2024-01-02 00:00:00",
		// },
	}, nil

	// var mockApis []CcamMockApi
	// err := config.DB.Where("api_id = ? AND deleted = 0 AND mock_type = 1", apiID).Find(&mockApis).Error
	// return mockApis, err
}

// CcamMockRule 对应ccam_mock_rule表结构
type CcamMockRule struct {
	RuleID           int64  `gorm:"primaryKey;column:rule_id"`
	MockApiID        int64  `gorm:"column:mock_api_id"`
	ConditionType    string `gorm:"column:condition_type"`
	ConditionKey     string `gorm:"column:condition_key"`
	ConditionValue   string `gorm:"column:condition_value"`
	Priority         int    `gorm:"column:priority"`
	ResponseTemplate string `gorm:"column:response_template"`
	StatusCode       int    `gorm:"column:status_code"`
	DelayMs          int    `gorm:"column:delay_ms"`
	CreateUser       string `gorm:"column:create_user"`
	CreateDate       string `gorm:"column:create_date"`
	UpdateUser       string `gorm:"column:update_user"`
	UpdateDate       string `gorm:"column:update_date"`
	Deleted          int    `gorm:"column:deleted"`
}

func (s *CcamMockRule) TableName() string {
	return "ccam_mock_rule"
}

// GetCcamMockRulesByMockApiID 根据模拟接口ID获取模拟规则，先硬编码返回
func GetCcamMockRulesByMockApiID(mockApiID int64) ([]CcamMockRule, error) {
	// 硬编码返回示例
	return []CcamMockRule{
		// {
		// 	RuleID:           1,
		// 	MockApiID:        mockApiID,
		// 	ConditionType:    "header",
		// 	ConditionKey:     "header-test",
		// 	ConditionValue:   "test",
		// 	ResponseTemplate: `{\"message\": \"Mock expectation response\"}`,
		// 	Deleted:          0,
		// },
		{
			RuleID:           2,
			MockApiID:        mockApiID,
			ConditionType:    "query",
			ConditionKey:     "query-test",
			ConditionValue:   "123",
			ResponseTemplate: `{\"message\": \"Mock expectation response\"}`,
			Deleted:          0,
		},
		// {
		// 	RuleID:           3,
		// 	MockApiID:        mockApiID,
		// 	ConditionType:    "body",
		// 	ConditionKey:     "body-test",
		// 	ConditionValue:   "abc",
		// 	ResponseTemplate: `{\"message\": \"Mock expectation response\"}`,
		// 	Deleted:          0,
		// },
	}, nil

	// var mockRules []CcamMockRule
	// err := config.DB.Where("mock_api_id = ? AND deleted = 0", mockApiID).Find(&mockRules).Error
	// return mockRules, err
}
