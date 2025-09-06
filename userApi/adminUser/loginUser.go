package adminUser

import (
	"encoding/json"
	"errors"
	"fmt"
	"project/common"
	"project/request"
	"project/userApi"
	"project/utils"
)

// 登录请求结构体
type LoginRequest struct {
	UserName  string `json:"userName"` // 账号
	Pwd       string `json:"pwd"`      // 密码
	Language  string `json:"language"`
	Random    int64  `json:"random"`
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"` //签名
}

// token的结构体
type Config struct {
	Token string `yaml:"token"`
}

/*
*
商户后台的登录
username 商户后台的账号
pwd 商户后台密码
*/
func Login(username, pwd string) error {
	api := "/api/Login/Login"
	loginData := LoginRequest{
		UserName:  username,
		Pwd:       pwd,
		Language:  "zh",
		Random:    request.RandmoNie(),
		Signature: "",
		Timestamp: request.GetNowTime(),
	}

	paylaodSignature := map[string]interface{}{
		"userName":  loginData.UserName,
		"pwd":       loginData.Pwd,
		"language":  loginData.Language,
		"random":    loginData.Random,
		"signature": loginData.Signature,
		"timestamp": loginData.Timestamp,
	}
	var baseurl common.CofingURL
	base_url := baseurl.ConfigUrlInit().ADMIN_URL
	// 获取请求头
	var head common.AdminHeaderConfig
	headMap := head.AdminHeaderConfigFunc()
	// respBody, err := request.PostRequest(paylaodSignature, api)
	respBody, _, err := request.PostRequestCofig(paylaodSignature, base_url, api, headMap)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// fmt.Printf("用户%v登录成功--------------------+\n", loginData.UserName)
	strResbody := string(respBody)
	var response userApi.Response
	error := json.Unmarshal([]byte(strResbody), &response)
	if error != nil {
		fmt.Println(error)
		return error
	}
	// fmt.Printf("登录结果%v", response)
	token, err := utils.HandlerMap(strResbody, "token")
	if err != nil {

		return err
	}
	config := Config{}
	config.Token = token
	errs := utils.WriteYAML(Token_addr_yaml_local, &config)
	if errs != nil {
		// fmt.Printf("token写入失败%v", errs)
		err := errors.New("token写入失败")
		fmt.Printf("%s", errs)
		return err
	}
	fmt.Printf("token写入成功.......\n")
	return nil
}
