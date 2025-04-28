package main

import (
	"encoding/json"
	"fmt"
	"wiremock_go_demo/faker"
	"wiremock_go_demo/models"
	"wiremock_go_demo/response"
)

// response_demo/main.go: 根据ccam库表数据生成api响应体
func main() {
	// 数据库通过init初始化
	var apiID int64 = 196 // 查询地址信息接口
	var libFaker faker.Faker = &faker.LibraryFaker{}

	responseList, _ := models.GetCcamApiResponsesByApiID(apiID)

	for _, rsp := range responseList {
		responseJsonMap := response.CreateRspJsonMap(rsp, libFaker)
		jsonData, _ := json.MarshalIndent(responseJsonMap, "", "  ")
		fmt.Println(string(jsonData))
	}
}
