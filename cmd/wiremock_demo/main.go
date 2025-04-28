package main

import (
	"fmt"
	"log"
	"strconv"
	"wiremock_go_demo/faker"
	"wiremock_go_demo/models"
	"wiremock_go_demo/response"

	"github.com/wiremock/go-wiremock"
)

var wiremockClient *wiremock.Client
var wireDomain = "localhost"
var wirePort = 8080

// wiremock_demo/main.go: 根据ccam库表数据生成wireMock stub
func main() {
	// 数据库通过init初始化
	wiremockClient = wiremock.NewClient(fmt.Sprintf("http://%s:%v", wireDomain, wirePort)) // 初始化WireMock客户端
	err := wiremockClient.Reset()                                                          // 清除mapping
	if err != nil {
		log.Fatalf("error ==>> wiremockClient.Reset() error=[%v]", err.Error())
		return
	}
	defaultMockInit()
	// expectationMockInit()

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
