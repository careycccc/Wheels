package deskApi

import (
	"encoding/json"
	"fmt"
	"project/common"
	"project/request"
	"project/utils"
)

// 注册接口

type RegisterStruct struct {
	UserName            string `json:"userName"`
	VerifyCode          string `json:"verifyCode"`
	InviteCode          string `json:"inviteCode"`
	RegisterFingerprint string `json:"registerFingerprint"`
	Random              int64  `json:"randmo"`
	Language            string `json:"language"`
	Signature           string `json:"signature"`
	Timestamp           int64  `json:"timestamp"`
	TrackStruct
}

type TrackStruct struct {
	IsTrusted bool  `json:"isTrusted"`
	Vts       int64 `json:"_vts"`
}

type ResponseResiter struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

/*
注册接口
userName  用户名
verifyCode 验证码
inviteCode 邀请码
return token  返回token
*
*/
func RegisterFunc(userName, verifyCode, inviteCode string) string {
	api := "/api/Home/MobileAutoLogin"
	base_url := common.SIT_WEB_API
	random := request.RandmoNie()
	timestamp := request.GetNowTime()
	generate := utils.GenerateCryptoRandomString(32)
	RegisterList := []interface{}{userName, verifyCode, inviteCode, generate, true, timestamp, "en", random, "", timestamp}
	registerMap, _ := InitializeRegisterStruct(RegisterList)
	register_url := common.REGISTER_url
	registreList := []interface{}{"3003", register_url, register_url, register_url}
	// 初始化请求头
	headerconfig := &common.DeskHeaderConfig2{}
	headerMap, _ := common.InitStructToMap(headerconfig, registreList)
	respBoy, _, err := request.PostRequestCofig(registerMap, base_url, api, headerMap)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 返回token
	// fmt.Println("注册返回的结果", string(respBoy))
	var response ResponseResiter
	err = json.Unmarshal(respBoy, &response)
	if err != nil {
		fmt.Println("解析响应失败: %v", err)
		return ""
	}
	return response.Data.Token

}

func InitializeRegisterStruct(data []interface{}) (map[string]interface{}, error) {
	// Check if slice has enough elements (10 expected)
	if len(data) != 10 {
		return nil, fmt.Errorf("expected 10 elements in slice, got %d", len(data))
	}

	// Create TrackStruct
	track := TrackStruct{
		IsTrusted: data[4].(bool),
		Vts:       data[5].(int64), // Empty string from slice
	}

	// Create RegisterStruct
	register := RegisterStruct{
		UserName:            data[0].(string),
		VerifyCode:          data[1].(string),
		InviteCode:          data[2].(string),
		RegisterFingerprint: data[3].(string),
		Random:              data[7].(int64),
		Language:            data[6].(string),
		Signature:           data[8].(string),
		Timestamp:           data[9].(int64),
		TrackStruct:         track,
	}

	// Convert to map[string]interface{}
	result := make(map[string]interface{})
	// Manually populate the map to control number formatting
	result["userName"] = register.UserName
	result["verifyCode"] = register.VerifyCode
	result["inviteCode"] = register.InviteCode
	result["registerFingerprint"] = register.RegisterFingerprint
	result["random"] = register.Random // int64, no scientific notation
	result["language"] = register.Language
	result["signature"] = register.Signature
	result["timestamp"] = register.Timestamp // int64, no scientific notation
	result["isTrusted"] = register.IsTrusted
	if register.Vts != 0 {
		result["_vts"] = register.Vts // Only include if non-empty due to omitempty
	} else {
		result["_vts"] = "" // Explicitly set to empty string as per slice
	}

	return result, nil
}
