package deskApi

import (
	"fmt"
	"project/request"
	"project/utils"
)

// 自动登录结构体
type autoLoginstruct struct {
	RegisterDevice      string `json:"registerDevice"`
	RegisterFingerprint string `json:"registerFingerprint"`
	InviteCode          string `json:"inviteCode"`
	PackageName         string `json:"packageName"`
	Language            string `json:"language"`
	Random              int64  `json:"random"`
	Signature           string `json:"signature"`
	Timestamp           int64  `json:"timestamp"`
}

func AutoLogin() {
	api := "/api/Home/AutoLogin"
	// 初始化结构体
	autologinData := autoLoginstruct{
		RegisterDevice:      "",
		RegisterFingerprint: utils.GenerateCryptoRandomString(32),
		InviteCode:          "",
		PackageName:         "",
		Language:            "en",
		Random:              request.RandmoNie(),
		Signature:           "",
		Timestamp:           request.GetNowTime(),
	}

	autoPayload := map[string]interface{}{
		"registerDevice":      autologinData.RegisterDevice,
		"registerFingerprint": autologinData.RegisterFingerprint,
		"inviteCode":          autologinData.InviteCode,
		"packageName":         autologinData.PackageName,
		"language":            autologinData.Language,
		"random":              autologinData.Random,
		"signature":           autologinData.Signature,
		"timestamp":           autologinData.Timestamp,
	}
	base_url := NewUserUrlFunc().userUrl
	response, _, err := request.PostRequestCofig(autoPayload, base_url, api)
	if err != nil {
		fmt.Printf("自动登录接口的post请求失败", err)
		return
	}
	respBody := string(response)
	fmt.Printf("响应结果:%v", respBody)
}
