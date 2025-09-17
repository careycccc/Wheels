package actingFy

import (
	"fmt"
	"project/common"
	"project/request"
)

// 查询验证码
type QueryTifyStruct struct {
	MobileOrEmail string `json:"mobileOrEmail"`
	common.BaseOrderByStruct
}

type QueryTifyStruct2 struct {
	common.BaseOrderByStruct
}

// NumberStruct 示例结构体，表示包含 Number 字段的类型
type NumberStruct struct {
	Number string `json:"number"`
}

func (n NumberStruct) GetField() string {
	return n.Number
}

// 输入用户账号，查询验证码
func QueryTifyFunc(userName string) string {
	api := "/api/Users/GetVerifyCodePageList"
	randmo := request.RandmoNie()
	timestamp := request.GetNowTime()
	// 结构体初始化
	query := &QueryTifyStruct{}
	queryList := []interface{}{userName, 1, 20, "Desc", randmo, "en", "", timestamp}
	headerStruct := &common.AdminHeaderAuthorizationConfig2{}
	headerList := []interface{}{}
	number, err := QueryTifyFuncGeneric[QueryTifyStruct, common.AdminHeaderAuthorizationConfig2, NumberStruct](api, query, queryList, headerStruct, headerList)
	if err != nil {
		fmt.Println("泛型调用失败", err)
		return ""
	}
	fmt.Println("验证码", number)
	return number
}

// / 点击了查询验证码的按钮
func QueryTifyFunc2() string {
	api := "/api/Users/GetVerifyCodePageList"
	query := &QueryTifyStruct2{}
	// 生成随机数和时间戳
	randmo := request.RandmoNie()
	timestamp := request.GetNowTime()
	// PageNo, PageSize, OrderBy, Random, Language, Signature, Timestamp
	queryList := []interface{}{1, 20, "Desc", randmo, "en", "", timestamp}
	headerStruct := &common.AdminHeaderAuthorizationConfig2{}
	headerList := []interface{}{}
	number, err := QueryTifyFuncGeneric[QueryTifyStruct2, common.AdminHeaderAuthorizationConfig2, NumberStruct](api, query, queryList, headerStruct, headerList)
	if err != nil {
		fmt.Println("泛型调用失败", err)
		return ""
	}
	fmt.Println("验证码", number)
	return number
}
