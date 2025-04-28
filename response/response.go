package response

import (
	"wiremock_go_demo/faker"
	"wiremock_go_demo/models"
)

func CreateRspJsonMap(response models.CcamApiResponse, faker faker.Faker) map[string]interface{} {
	var responseJsonMap = make(map[string]interface{})
	paramsList, _ := models.GetCcamApiResponseParamsByResponseID(response.ResponseID)

	if len(paramsList) == 0 {
		// fmt.Printf("rspID=[%v] rspCode=[%v], 该响应情况没有响应字段\n", response.ResponseID, response.ResponseCode)
		return map[string]interface{}{
			"code": response.ResponseCode,
			"msg":  response.ResponseDescription,
		}
	}

	for _, param := range paramsList {
		key := param.ParamName
		t := param.ParamType
		if t == "integer" {
			responseJsonMap[key] = faker.FakerInteger(param.ParamFormat)
		} else if t == "float" {
			responseJsonMap[key] = faker.FakerFloat(param.ParamFormat)
		} else if t == "string" {
			responseJsonMap[key] = faker.FakerString(param.ParamFormat)
		} else if t == "ref" {
			responseJsonMap[key] = faker.ObjectFaker(param.RefSchemaID)
		}
	}
	return responseJsonMap
}
