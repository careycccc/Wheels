package deskApi

import (
	"encoding/json"
	"fmt"
	"project/common"
	"project/request"
)

// 定义结构体来映射 JSON 数据
type ClickShareResponse struct {
	Data struct {
		InviteCode string `json:"inviteCode"`
	} `json:"data"`
}

// 点击分享按钮 // 返回一个邀请码
func ClickShareFunc(userName, token string) string {
	api := "/api/Activity/GetUserInviteLinkAddress"
	base_url := common.SIT_WEB_API
	// 点击4个礼物盒
	payloadStruct := &common.BaseStruct{}
	randmo := request.RandmoNie()
	timestamp := request.GetNowTime()
	payloadList := []interface{}{randmo, "en", "", timestamp}
	payloadMap := common.InitStructToMap(payloadStruct, payloadList)
	// 请求头
	headerStruct := &common.DeskHeaderAstruct2{}
	headerUrl := common.PLANT_H5
	headerList := []interface{}{"3003", headerUrl, headerUrl, headerUrl, token}
	headerMap, _ := common.AssignSliceToStructMap(headerStruct, headerList)
	// 发送请求
	respBody, _, err := request.PostRequestCofig(payloadMap, base_url, api, headerMap)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 解析 JSON 数据
	var resp ClickShareResponse
	err = json.Unmarshal([]byte(string(respBody)), &resp)
	if err != nil {
		fmt.Printf("JSON 解析错误: %v", err)
		return ""
	}
	fmt.Printf("%s点击了一次分享", userName)
	return resp.Data.InviteCode
}
