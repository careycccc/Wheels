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

// 总代注册 请求  返回token
func GeneralRegiterFunc(userName string) string {
	api := "/api/Home/MobileAutoLogin"
	base_url := common.SIT_WEB_API
	// 获取请求头
	headerStruct := &common.DeskHeaderConfig2{}
	headerUrl := common.PLANT_H5
	headerList := []interface{}{"3003", headerUrl, headerUrl, headerUrl}
	headerMap, _ := common.InitStructToMap(headerStruct, headerList)
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
		devices := utils.GenerateCryptoRandomString(16)
		fmt.Println(" 注册设备：", devices)
		fmt.Println(" 浏览器指纹：", cry)
		randmo := request.RandmoNie()
		timestamp := request.GetNowTime()
		payloadList := []interface{}{userName, verifyCode, devices, cry, "", "", "", randmo, "en", "", timestamp}
		payloadMap, err := common.StructToMap(payloadStruct, payloadList)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		// 将嵌套map进行平铺
		flatMap := common.FlattenMap(payloadMap)
		respBody, _, err, _ := request.PostRequestCofigProxy(flatMap, base_url, api, headerMap)
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
		fmt.Printf("%s注册成功并且已经登录++++++++\n", userName)
		return resp.Data.Token
	}
	return ""
}

// GeneralRegiterFuncProxy 使用代理的总代注册函数
// 如果代理失效会自动尝试下一个可用的代理
// 返回格式: "token,ip" (直连)
func GeneralRegiterFuncProxy(userName string) (string, string) {
	api := "/api/Home/MobileAutoLogin"
	base_url := common.SIT_WEB_API

	fmt.Printf("开始为 %s 进行总代注册（使用代理）...\n", userName)

	// 获取请求头
	headerStruct := &common.DeskHeaderConfig2{}
	headerUrl := common.PLANT_H5
	headerList := []interface{}{"3003", headerUrl, headerUrl, headerUrl}
	headerMap, err := common.InitStructToMap(headerStruct, headerList)
	if err != nil {
		fmt.Printf("初始化请求头失败: %v\n", err)
		return "", ""
	}

	// 初始化payload
	payloadStruct := &GeneralRegiterStruct{}

	// 发送验证码
	fmt.Println("发送验证码...")
	actingFy.SendVerifiyCodeFunc(userName)

	// 获取验证码
	fmt.Println("等待验证码...")
	time.Sleep(time.Second * 2)
	verifyCode := actingFy.QueryTifyFunc2()

	if len(verifyCode) == 0 {
		fmt.Println("获取验证码失败")
		return "", ""
	}

	fmt.Printf("获取到验证码: %s\n", verifyCode)

	// 准备请求参数
	cry := utils.GenerateCryptoRandomString(32)
	devices := utils.GenerateCryptoRandomString(16)
	randmo := request.RandmoNie()
	timestamp := request.GetNowTime()
	payloadList := []interface{}{userName, verifyCode, devices, cry, "", "", "", randmo, "en", "", timestamp}

	payloadMap, err := common.StructToMap(payloadStruct, payloadList)
	if err != nil {
		fmt.Printf("构建请求参数失败: %v\n", err)
		return "", ""
	}

	// 将嵌套map进行平铺
	flatMap := common.FlattenMap(payloadMap)

	fmt.Println("使用代理发送注册请求...")
	respBody, _, err, ipInfo := request.PostRequestCofigProxy(flatMap, base_url, api, headerMap)
	if err != nil {
		fmt.Printf("注册请求失败: %v\n", err)
		return "", ""
	}

	fmt.Printf("使用的IP: %s\n", ipInfo)

	// 解析 JSON 数据
	var resp Response
	err = json.Unmarshal([]byte(string(respBody)), &resp)
	if err != nil {
		fmt.Printf("JSON 解析错误: %v\n", err)
		return "", ""
	}

	// 输出 token 值和使用的IP
	fmt.Printf("%s 注册成功并且已经登录（使用代理）++++++++\n", userName)
	fmt.Printf("使用的IP: %s\n", ipInfo)

	// 返回格式: "token|proxy_ip:port" 或 "token|direct"
	return resp.Data.Token, ipInfo
}
