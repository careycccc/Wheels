package actingFy

import (
	"encoding/json"
	"fmt"
	"project/common"
	"project/request"
	"project/userApi/adminUser"
)

// 查询验证码

type QueryTifyStruct struct {
	MobileOrEmail string `json:"mobileOrEmail"`
	common.BaseOrderByStruct
}

type QueryTifyStruct2 struct {
	common.BaseOrderByStruct
}

// 定义与 JSON 结构对应的 Go 结构体
type Response struct {
	Data struct {
		List []struct {
			Number string `json:"number"`
		} `json:"list"`
	} `json:"data"`
}

// 验证码查询 返回验证码
func QueryTifyFunc(userName string) string {
	api := "/api/Users/GetVerifyCodePageList"
	randmo := request.RandmoNie()
	timestamp := request.GetNowTime()
	// 结构体初始化
	query := &QueryTifyStruct{}
	queryList := []interface{}{userName, 1, 20, "Desc", randmo, "en", "", timestamp}
	queryMap, _ := common.StructToMap(query, queryList)
	// 请求头
	token, err := adminUser.GetToken()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	base_url := common.ADMIN_SYSTEM_url
	headerStruct := &common.AdminHeaderAuthorizationConfig2{}
	headerList := []interface{}{base_url, base_url, base_url, token}
	headerMap, _ := common.AssignSliceToStructMap(headerStruct, headerList)
	// 把嵌套的的map进行平铺
	fietMap := common.FlattenMap(queryMap)
	resp, _, err := request.PostRequestCofig(fietMap, base_url, api, headerMap)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 解析 JSON
	var response Response
	err = json.Unmarshal([]byte(string(resp)), &response)
	if err != nil {
		fmt.Printf("JSON 解析失败: %v", err)
	}

	// 提取 number 的值
	if len(response.Data.List) > 0 {
		number := response.Data.List[0].Number
		fmt.Println("Number 的值:", number)
		return number
	} else {
		fmt.Println("List 为空，无法提取 number")
	}
	return ""
}

func QueryTifyFunc2() string {
	api := "/api/Users/GetVerifyCodePageList"
	randmo := request.RandmoNie()
	timestamp := request.GetNowTime()
	// 结构体初始化
	query := &QueryTifyStruct2{}
	queryList := []interface{}{1, 20, "Desc", randmo, "en", "", timestamp}
	queryMap, _ := common.StructToMap(query, queryList)
	// 请求头
	token, err := adminUser.GetToken()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	base_url := common.ADMIN_SYSTEM_url
	headerStruct := &common.AdminHeaderAuthorizationConfig2{}
	headerList := []interface{}{base_url, base_url, base_url, token}
	headerMap, _ := common.AssignSliceToStructMap(headerStruct, headerList)
	// 把嵌套的的map进行平铺
	fietMap := common.FlattenMap(queryMap)
	resp, _, err := request.PostRequestCofig(fietMap, base_url, api, headerMap)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 解析 JSON
	var response Response
	err = json.Unmarshal([]byte(string(resp)), &response)
	if err != nil {
		fmt.Printf("JSON 解析失败: %v", err)
	}

	// 提取 number 的值
	if len(response.Data.List) > 0 {
		number := response.Data.List[0].Number
		// fmt.Println("Number 的值:", number)
		return number
	} else {
		fmt.Println("List 为空，无法提取 number")
	}
	return ""
}
