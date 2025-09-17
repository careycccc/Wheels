package deskApi

// 旋转邀请转盘的免费次数

import (
	"fmt"
	"project/common"
	"project/request"
)

// 点击转盘提现

type ClickFreeWheelWithdraw struct {
	common.BaseStruct
}

type ClickFreeWheelWithdrawResponse struct {
	Code    int `json:"code"`
	MsgCode int `json:"msgCode"`
	Data    any `json:"data"`
}

// 需要传入token
func ClickFreeWheelFunc(token string) (string, error) {
	api := "/api/Activity/SpinInvitedWheel"
	base_url := common.SIT_WEB_API
	randmo := request.RandmoNie()
	timeStamp := request.GetNowTime()
	// 组装请求体
	payloadStruct := &ClickFreeWheelWithdraw{}
	paylaodList := []interface{}{randmo, "en", "", timeStamp}
	payloadMap, err := common.StructToMap(payloadStruct, paylaodList)
	if err != nil {
		return "", fmt.Errorf("请求体组装失败: %v", err)
	}
	fildMap := common.FlattenMap(payloadMap)
	// 获取请求头
	headerStruct := &common.DeskHeaderAstruct2{}
	headerUrl := common.PLANT_H5
	headerList := []interface{}{"3003", headerUrl, headerUrl, headerUrl, token}
	headMap, err := common.AssignSliceToStructMap(headerStruct, headerList)
	if err != nil {
		fmt.Println("请求头组装失败", err)
		return "", err
	}
	// fmt.Println("免费请求头的参数-----", payloadMap)
	respBoy, _, err := request.PostRequestCofig(fildMap, base_url, api, headMap)
	if err != nil {
		return "", fmt.Errorf("请求失败: %v", err)
	}

	fmt.Println("点击免费的次数的旋转的结果", string(respBoy))

	return string(respBoy), nil
}
