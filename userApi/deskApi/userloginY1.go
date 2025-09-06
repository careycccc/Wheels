package deskApi

import (
	"encoding/json"
	"fmt"
	"project/common"
	"project/request"
	"project/userApi"
	"project/utils"
)

type userloginY1 struct {
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	LoginType   string `json:"loginType"`
	DeviceId    string `json:"deviceId"`
	BrowserId   string `json:"browserId"`
	PackageName string `json:"packageName"`
	Language    string `json:"language"`
	Random      int64  `json:"random"`
	Signature   string `json:"signature"`
	Timestamp   int64  `json:"timestamp"`
}

// 账号，密码 登录
func UserloginY1(username, password string) (string, error) {
	api := "/api/Home/Login"
	userloginInit := userloginY1{
		UserName:    username,
		Password:    password,
		LoginType:   "Mobile",
		DeviceId:    "",
		BrowserId:   utils.GenerateCryptoRandomString(32),
		PackageName: "",
		Language:    "en",
		Random:      request.RandmoNie(),
		Signature:   "",
		Timestamp:   request.GetNowTime(),
	}

	userloginMap := map[string]interface{}{
		"userName":    userloginInit.UserName,
		"password":    userloginInit.Password,
		"loginType":   userloginInit.LoginType,
		"deviceId":    userloginInit.DeviceId,
		"browserId":   userloginInit.BrowserId,
		"packageName": userloginInit.PackageName,
		"language":    userloginInit.Language,
		"random":      userloginInit.Random,
		"signature":   userloginInit.Signature,
		"timestamp":   userloginInit.Timestamp,
	}
	base_url := NewUserUrlFunc().userUrl
	// 获取请求头
	headMap := common.NewDeskHeaderConfig().DeskHeaderConfigFunc()
	resp, _, err := request.PostRequestCofig(userloginMap, base_url, api, headMap)
	if err != nil {
		fmt.Println("输入账号和密码登录的post请求失败")
		return "输入账号和密码登录的post请求失败", err
	}

	strResbody := string(resp)
	var response userApi.Response
	error := json.Unmarshal([]byte(strResbody), &response)
	if error != nil {
		fmt.Println(error)
		return "登录请求失败", error
	}
	// fmt.Printf("登录结果%v", response)
	token, err := utils.HandlerMap(strResbody, "token")
	if err != nil {
		return "寻找token失败", err
	}
	// fmt.Printf("token==>%v\n", token)
	fmt.Printf("%s登录成功--------------------------------------------", username)
	return token, nil
}
