package deskApi

import (
	"encoding/json"
	"fmt"
	"project/common"
	"project/request"
	"project/utils"
)

//   "userName": "919091997115",
//   "verifyCode": "614377",
//   "inviteCode": "YDWY52N",
//   "registerFingerprint": "e21f226e2db8717cc38568331be44cf3",
//   "track": { "isTrusted": true, "_vts": 1757399740332 },

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

/*
注册接口
userName  用户名
verifyCode 验证码
inviteCode 邀请码
*
*/
func RegisterFunc(userName, verifyCode, inviteCode string) {
	api := "/api/Home/MobileAutoLogin"
	base_url := common.SIT_WEB_API
	random := request.RandmoNie()
	timestamp := request.GetNowTime()
	generate := utils.GenerateCryptoRandomString(32)
	RegisterList := []interface{}{userName, verifyCode, inviteCode, generate, true, timestamp * 1000, "en", random, "", timestamp}
	registerMap, _ := InitializeRegisterStruct(RegisterList)
	register_url := common.REGISTER_url
	registreList := []interface{}{register_url, register_url, register_url}
	headerconfig := &common.DeskHeaderConfig2{}
	headerMap := common.InitStructToMap(headerconfig, registreList)
	resp, _, err := request.PostRequestCofig(registerMap, base_url, api, headerMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	println(string(resp))
}

func InitializeRegisterStruct(data []interface{}) (map[string]interface{}, error) {
	// Check if slice has enough elements (10 expected)
	if len(data) != 10 {
		return nil, fmt.Errorf("expected 10 elements in slice, got %d", len(data))
	}

	// Create TrackStruct
	track := TrackStruct{
		IsTrusted: data[4].(bool),
		Vts:       data[5].(int64),
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
	bytes, err := json.Marshal(register)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal struct: %v", err)
	}

	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal to map: %v", err)
	}

	return result, nil
}
