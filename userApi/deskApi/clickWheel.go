package deskApi

import (
	"encoding/json"
	"fmt"
	"project/common"
	"project/request"
)

// 定义结构体来映射 JSON 数据
type ClickResponse struct {
	Data struct {
		Amount bool `json:"amount"`
	} `json:"data"`
}

// 点击4个礼物盒
func ClickWheelFunc(userName, token string) bool {
	api := "/api/Activity/SpinInvitedWheel"
	base_url := common.SIT_WEB_API
	payloadStruct := &common.BaseStruct{}
	randmo := request.RandmoNie()
	timestamp := request.GetNowTime()
	payloadList := []interface{}{randmo, "en", "", timestamp}
	payloadMap, _ := common.InitStructToMap(payloadStruct, payloadList)
	// 请求头
	headerStruct := &common.DeskHeaderAstruct2{}
	headerUrl := common.PLANT_H5
	headerList := []interface{}{"3003", headerUrl, headerUrl, headerUrl, token}
	headerMap, _ := common.AssignSliceToStructMap(headerStruct, headerList)
	// 发送请求
	respBody, _, err := request.PostRequestCofig(payloadMap, base_url, api, headerMap)
	if err != nil {
		fmt.Println(err)
		return false
	}
	// fmt.Println(string(respBody))
	// 解析 JSON 数据
	var resp ClickResponse
	err = json.Unmarshal([]byte(string(respBody)), &resp)
	if err != nil {
		fmt.Printf("JSON 解析错误: %v", err)
		return false
	}
	fmt.Println("点击这个4个礼物盒的一个", resp.Data.Amount)
	return resp.Data.Amount
}
