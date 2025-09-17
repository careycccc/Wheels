package actingFy

import (
	"encoding/json"
	"fmt"
	"project/common"
	"project/request"
	"project/userApi/adminUser"
)

// FieldExtractor 接口定义提取字段的通用方法
type FieldExtractor interface {
	GetField() string
}

// // NumberStruct 示例结构体，表示包含 Number 字段的类型
// type NumberStruct struct {
// 	Number string `json:"number"`
// }

// func (n NumberStruct) GetField() string {
// 	return n.Number
// }

// // AmountStruct 示例结构体，表示包含 Amount 字段的类型
// type AmountStruct struct {
// 	Amount string `json:"amount"`
// }

// func (a AmountStruct) GetField() string {
// 	return a.Amount
// }

// Response 泛型响应结构体
type Response[T FieldExtractor] struct {
	Data struct {
		List []T `json:"list"`
	} `json:"data"`
}

// 查询的泛型函数
func QueryTifyFuncGeneric[Q any, H any, T FieldExtractor](api string, query *Q, queryList []interface{}, headerStruct *H, headerList []interface{}) (string, error) {
	// 结构体转 Map
	queryMap, err := common.StructToMap(query, queryList)
	if err != nil {
		return "", fmt.Errorf("failed to convert query struct to map: %v", err)
	}
	// 获取 token 和 base_url
	token := adminUser.GetToken()
	base_url := common.ADMIN_SYSTEM_url

	// 确保 headerList 包含必要参数
	headerList = append(headerList, base_url, base_url, base_url, token)
	headerMap, err := common.AssignSliceToStructMap(headerStruct, headerList)
	if err != nil {
		return "", fmt.Errorf("failed to convert header struct to map: %v", err)
	}
	// 平铺嵌套 Map
	fietMap := common.FlattenMap(queryMap)

	// 发送 POST 请求
	resp, _, err := request.PostRequestCofig(fietMap, base_url, api, headerMap)
	if err != nil {
		return "", fmt.Errorf("POST request failed: %v", err)
	}

	// 解析 JSON
	var response Response[T]
	err = json.Unmarshal([]byte(string(resp)), &response)
	if err != nil {
		return "", fmt.Errorf("JSON parsing failed: %v", err)
	}

	// 提取字段值
	if len(response.Data.List) > 0 {
		return response.Data.List[0].GetField(), nil
	}

	return "", fmt.Errorf("list is empty, cannot extract field")
}
