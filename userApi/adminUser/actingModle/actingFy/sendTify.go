package actingFy

import (
	"fmt"
	"project/common"
	"project/request"
)

// 发送验证码
type SendVerifiyCodeStruct struct {
	VerifyCodeType any `json:"verifyCodeType"`
	PhoneOrEmail   any `json:"phoneOrEmail"`
	CodeType       any `json:"codeType"`
	Language       any `json:"language"`
	Random         any `json:"random"`
	Signature      any `json:"signature"`
	Timestamp      any `json:"timestamp"`
}

// 传入用户账号加区号 返回验证码
func SendVerifiyCodeFunc(userName string) string {
	random := request.RandmoNie()
	timestamp := request.GetNowTime()
	api := "/api/Home/SendVerifiyCode"
	base_url := common.SIT_WEB_API
	registryUrl := common.REGISTER_url
	// 获取请求头
	config := &common.DeskHeaderConfig2{}
	headerList := []interface{}{"3003", registryUrl, registryUrl, registryUrl}
	heaerMap, _ := common.InitStructToMap(config, headerList)
	// 发送验证码的结构体初始化
	SendVerifiyCode := &SendVerifiyCodeStruct{}
	assiginSlice := []interface{}{1, userName, 18, "en", random, "", timestamp}
	requestMap, err := common.AssignSliceToStructMap(SendVerifiyCode, assiginSlice)
	if err != nil {
		fmt.Println("发送验证码的结构体初始化失败")
		return ""
	}
	resp, _, err := request.PostRequestCofig(requestMap, base_url, api, heaerMap)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Printf("%v,发送验证码成功%v\n", userName, string(resp))
	return ""
}
