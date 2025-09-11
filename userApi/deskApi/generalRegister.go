package deskApi

import (
	"encoding/json"
	"fmt"
	"project/common"
	"project/request"
	"project/userApi/adminUser/actingModle/actingFy"
	"project/utils"
	"time"
)

// 总代注册，没有上级的注册方式

type GeneralRegiterStruct struct {
	UserName            string `json:"userName"`
	VerifyCode          string `json:"verifyCode"`
	RegisterDevice      string `json:"registerDevice"`
	RegisterFingerprint string `json:"registerFingerprint"`
	InviteCode          string `json:"inviteCode"`
	Rrack               string `json:"track"`
	PackageName         string `json:"packageName"`
	common.BaseStruct
}

// 定义结构体来映射 JSON 数据
type Response struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

// 总代注册 请求
func GeneralRegiterFunc(userName string) string {
	api := "/api/Home/MobileAutoLogin"
	base_url := common.SIT_WEB_API
	// 获取请求头
	headerStruct := &common.DeskHeaderConfig2{}
	headerUrl := common.PLANT_H5
	headerList := []interface{}{"3003", headerUrl, headerUrl, headerUrl}
	headerMap := common.InitStructToMap(headerStruct, headerList)
	// 初始化payload
	payloadStruct := &GeneralRegiterStruct{}
	// 发送验证码
	actingFy.SendVerifiyCodeFunc(userName) // 发送验证码
	// 获取验证码
	time.Sleep(time.Second * 2)
	verifyCode := actingFy.QueryTifyFunc2() //获取验证码
	if len(verifyCode) > 0 {
		// 获取验证码
		// 随机浏览器指纹
		cry := utils.GenerateCryptoRandomString(32)
		randmo := request.RandmoNie()
		timestamp := request.GetNowTime()
		payloadList := []interface{}{userName, verifyCode, "", cry, "", "", "", randmo, "en", "", timestamp}
		payloadMap, err := common.StructToMap(payloadStruct, payloadList)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		// 将嵌套map进行平铺
		flatMap := common.FlattenMap(payloadMap)
		respBody, _, err := request.PostRequestCofig(flatMap, base_url, api, headerMap)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		// 解析 JSON 数据
		var resp Response
		err = json.Unmarshal([]byte(string(respBody)), &resp)
		if err != nil {
			fmt.Printf("JSON 解析错误: %v", err)
		}

		// 输出 token 值
		fmt.Printf("%s注册成功并且已经登录++++++++", userName)
		return resp.Data.Token
	}
	return ""
}
