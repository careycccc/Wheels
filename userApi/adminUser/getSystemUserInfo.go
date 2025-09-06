package adminUser

import (
	"fmt"
	"project/request"
	"project/utils"
)

type GetSysUserInfo struct {
	Random    int64  `json:"random"`
	Language  string `json:"language"`
	Signature string `json:"signature"`
	Timestamp int64  `json:"timestamp"`
}

func GetSystemUserInfo() {
	api := "/api/Home/GetSysUserInfo"
	// 初始化payload
	getSysUserInfo := GetSysUserInfo{
		Random:    request.RandmoNie(),
		Language:  "zh",
		Signature: " ",
		Timestamp: request.GetNowTime(),
	}
	var payload map[string]interface{}
	payload = map[string]interface{}{
		"random":    getSysUserInfo.Random,
		"language":  getSysUserInfo.Language,
		"signature": getSysUserInfo.Signature,
		"timestamp": getSysUserInfo.Timestamp,
	}
	// 设置请求头，和获取token
	headMap, base_url := GetHeaderUrl()
	respBody, _, err := request.PostRequestCofig(payload, base_url, api, headMap)
	if err != nil {
		fmt.Println("请求GetSystemUserInfo出错", respBody)
		return
	}
	// _, err := utils.HandlerMap(string(respBody), "userId")
	// if err != nil {
	// 	return
	// }
	// fmt.Printf("~~~~--------------------%v", string(respBody))
	result := utils.Unmarshal(string(respBody))
	innerMap, ok := result["data"].(map[string]interface{})
	if !ok {
		// fmt.Println("data 不存在")
		return
	}
	fmt.Printf("外层%T", innerMap)
}
