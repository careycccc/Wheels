package deskApi

import (
	"encoding/json"
	"fmt"
	"project/common"
	"project/request"
)

type PostReponsestruct struct {
	Data    any    `json:"data"`
	Code    int    `json:"code"`    // 0
	Msg     string `json:"msg"`     // Succeed
	MsgCode int    `json:"msgCode"` // 0
}

/*
注册的泛型
P,H是需要传入的两个结构体
p 表示的是 payload的结构体
H 表示的是 header的结构体
api 表示接口地址
payloadList 表示需要赋值的 payload
headerList 表示需要赋值的 header
payloadFunc 表示需要处理的pay的func
headerFunc 表示需要处理的header的func
token 需要传入的token值
*
*/
func PostGenericsFunc[P, H any](api string, payload *P, payloadList []interface{}, headerStruct *H,
	headerList []interface{}, payloadFunc func(structType interface{}, slice []interface{}) (map[string]interface{}, error), headerFunc func(structType interface{}, slice []interface{}) (map[string]interface{}, error)) *PostReponsestruct {
	// 结构体转 Map
	payloadMap, err := payloadFunc(payload, payloadList)
	if err != nil {
		return &PostReponsestruct{
			Data:    "",
			Code:    1,
			Msg:     "failed to convert payloadMap struct to map",
			MsgCode: 1,
		}
	}
	// 获取 token 和 base_url
	base_url := common.SIT_WEB_API
	// 确保 headerList 包含必要参数
	headerMap, err := headerFunc(headerStruct, headerList)
	if err != nil {
		return &PostReponsestruct{
			Data:    "",
			Code:    1,
			Msg:     "failed to convert headerMap struct to map",
			MsgCode: 1,
		}
	}

	respBody, _, err := request.PostRequestCofig(payloadMap, base_url, api, headerMap)
	if err != nil {
		return &PostReponsestruct{
			Data:    "",
			Code:    1,
			Msg:     fmt.Sprintf("错误代码:%s", err),
			MsgCode: 1,
		}
	}
	var result PostReponsestruct
	err = json.Unmarshal([]byte(string(respBody)), &result)
	if err != nil {
		return &PostReponsestruct{
			Data:    "",
			Code:    1,
			Msg:     fmt.Sprintf("错误代码:%s", err),
			MsgCode: 1,
		}
	}
	return &result
}
