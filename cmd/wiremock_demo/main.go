package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"wiremock_go_demo/faker"
	"wiremock_go_demo/models"
	"wiremock_go_demo/response"

	"github.com/wiremock/go-wiremock"
)

var wiremockClient *wiremock.Client
var wireDomain = "localhost"
var wirePort = 8099

// wiremock_demo/main.go: 根据ccam库表数据生成wireMock stub
func main() {
	// 数据库通过init初始化
	wiremockClient = wiremock.NewClient(fmt.Sprintf("http://%s:%v", wireDomain, wirePort)) // 初始化WireMock客户端
	err := wiremockClient.Reset()                                                          // 清除mapping
	if err != nil {
		log.Fatalf("error ==>> wiremockClient.Reset() error=[%v]", err.Error())
		return
	}

	// 默认mock已验证, 先注释
	// defaultMockInit()
	expectationMockInit()

	// 阻塞主进程，防止程序退出
	// stop := make(chan os.Signal, 1)
	// signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	// go func() {
	// 	<-stop
	// 	fmt.Println("\n正在清理资源并退出...")
	//  wiremockClient.Reset()
	// }()

	// select {}
}

// defaultMockInit: 默认Mock
func defaultMockInit() {
	var apiID int64 = 196
	var libFaker faker.Faker = &faker.LibraryFaker{}

	apiInfo, _ := models.GetCcamApiInfoByApiID(apiID)
	responseList, _ := models.GetCcamApiResponsesByApiID(apiID)

	for _, rsp := range responseList {
		responseJsonMap := response.CreateRspJsonMap(rsp, libFaker)
		errCode, _ := strconv.Atoi(rsp.ResponseCode)

		// 这里需要生成一个接口唯一路径, 并且不同响应码要通过query参数区分. 对于正常的200默认没有query参数
		stubPath := fmt.Sprintf("/%v%s", apiInfo.ApiID, apiInfo.ApiPath)          // 接口主键ID + 接口路径
		realPath := fmt.Sprintf("http://%s:%v%s", wireDomain, wirePort, stubPath) // debug用
		if errCode != 200 {
			// stub注册的时候?前面要加\, 但是访问的时候不需要.
			stubPath += fmt.Sprintf("\\?rspID=%v", errCode)
			realPath += fmt.Sprintf("?rspID=%v", errCode)
		}

		// 默认stub
		// todo: 暂时写死Get方便测试, apiInfo.ApiMethod
		stubRule := wiremock.NewStubRule("GET", wiremock.URLMatching(stubPath)).
			WillReturnResponse(
				wiremock.NewResponse().
					WithStatus(int64(errCode)).
					WithJSONBody(responseJsonMap), // 不需要我们marshal, 传json map即可
				// WithHeader("Content-Type", "application/json"), // 默认mock没有header
			)

		err := wiremockClient.StubFor(stubRule)
		if err != nil {
			fmt.Printf("error ==>> stub规则=[%v], stub错误=[%v]\n", stubRule, err.Error())
		} else {
			log.Printf("stub success. 访问路径=[%s]", realPath)
		}
	}
}

// expectationMockInit: 期望Mock
func expectationMockInit() {
	/*
		前端业务逻辑
		用户配置了一个mock 期望以后, 应该会在ccam_mock_api表里生成一个记录后
		拿到mock_api_id主键后再到ccam_mock_rule表里生成规则记录

		而这里我们要实现mock期望的stub mapping
		捞取到属于mock期望的mock_api记录(mock_type = 1)后, 每条记录都去ccam_mock_rule表捞取规则
		然后根据规则进行stub注册
	*/

	// 测试接口API ID
	var apiID int64 = 196
	// 捞取mock期望记录
	mockExpectations, _ := models.GetExpectedMockApi(apiID)

	// 遍历mock期望记录, 每个mock期望记录都去ccam_mock_rule表捞取规则
	for _, mockApi := range mockExpectations {
		// 捞取规则
		mockRules, _ := models.GetCcamMockRulesByMockApiID(mockApi.MockApiID)

		// stubPath: stub路径, 如果要使用URL Matching, 需要把query参数拼接到stubPath一起注册, 否则无法匹配
		stubPath := fmt.Sprintf("/%v%s", mockApi.ApiID, mockApi.ApiPath)

		// realPath: 真实访问路径, debug用
		realPath := fmt.Sprintf("http://%s:%v%s", wireDomain, wirePort, stubPath)
		var queryParams []string
		// 遍历mock期望规则，收集query参数
		for _, rule := range mockRules {
			if rule.ConditionType == "query" {
				queryParams = append(queryParams, fmt.Sprintf("%s=%s", rule.ConditionKey, rule.ConditionValue))
			}
		}
		// 突然想到, stubPath要加query参数只是因为默认mock要区分不同响应码才需要
		// 而期望mock我们要判断query, 直接通过sdk即可
		if len(queryParams) > 0 {
			realPath += "?" + strings.Join(queryParams, "&")
		}

		// stubRule stub的规则, 遍历每条期望规则完善rule
		var stubRule *wiremock.StubRule
		stubRule = wiremock.NewStubRule(mockApi.ApiMethod, wiremock.URLPathMatching(stubPath))

		// 目前响应规则是写在rule表, 并且rule有多条记录, 通过flag判断只加一次响应
		var isFirst bool = true

		// 遍历mock期望规则，构建匹配规则
		for _, rule := range mockRules {
			switch rule.ConditionType {
			case "header":
				stubRule = stubRule.WithHeader(rule.ConditionKey, wiremock.EqualTo(rule.ConditionValue))
			case "query":
				stubRule = stubRule.WithQueryParam(rule.ConditionKey, wiremock.EqualTo(rule.ConditionValue))
			case "path":
				stubRule = stubRule.WithPathParam(rule.ConditionKey, wiremock.EqualTo(rule.ConditionValue))
			case "body":
				// body的判断函数都是只有一个入参, 难道是要把key&value一起marshal后传入?
				s, _ := json.Marshal(map[string]interface{}{rule.ConditionKey: rule.ConditionValue})
				stubRule = stubRule.WithBodyPattern(wiremock.MatchingJsonPath(string(s)))
			}

			if isFirst && stubRule != nil {
				// 数据库里存的是json字符串, 要反序列化成map
				var responseTemplate map[string]interface{}
				json.Unmarshal([]byte(rule.ResponseTemplate), &responseTemplate)
				s, _ := json.Marshal(responseTemplate)

				stubRule = stubRule.WillReturnResponse(
					wiremock.NewResponse().
						WithStatus(200).
						WithBody(string(s)).
						WithHeader("Content-Type", "application/json"),
				)
				isFirst = false
			}
		}

		fmt.Printf("-----------------------------------------\n")
		fmt.Println("stubPath ==>>", stubPath)
		fmt.Println("realPath ==>>", realPath)
		fmt.Printf("stubRule ==>> [%v]\n", stubRule)
		fmt.Printf("-----------------------------------------\n")

		// 注册stub
		err := wiremockClient.StubFor(stubRule)
		if err != nil {
			fmt.Printf("error ==>> stub规则=[%v], stub错误=[%v]\n", stubRule, err.Error())
		} else {
			log.Printf("stub success. 访问路径=[%s]", realPath)
		}
	}
}

func useless() {
	// map[conditionType]map[conditionKey]conditionValue
	// mockRule := map[string]map[string]string{
	// 	"header": {
	// 		"Content-Type": "application/json", // 期望header
	// 	},
	// 	"body": {
	// 		"name": "John", // 期望body
	// 	},
	// 	"query": {
	// 		"age": "30", // 期望query
	// 	},
	// 	"path": {
	// 		"id": "123", // 期望path
	// 	},
	// }
	// wiremock.NewStubRule("GET", wiremock.URLPathMatching("/testExpectation"))
}
