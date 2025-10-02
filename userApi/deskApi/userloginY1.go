package deskApi

import (
	"encoding/json"
	"fmt"
	"project/common"
	"project/request"
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

func UserloginY1(username, password string) (string, error) {
	api := "/api/Home/Login"
	payloadStruct := &userloginY1{}
	browserid := utils.GenerateCryptoRandomString(32)
	randmo := request.RandmoNie()
	timestamp := request.GetNowTime()
	payloadListe := []interface{}{username, password, "Mobile", "", browserid, "", "en", randmo, "", timestamp}
	header_url := common.PLANT_H5
	headerStruct := &common.DeskHeaderConfig2{}
	headerList := []interface{}{"3003", header_url, header_url, header_url}
	response := request.PostGenericsFunc[userloginY1, common.DeskHeaderConfig2](api, payloadStruct, payloadListe, headerStruct, headerList, common.InitStructToMap, common.InitStructToMap)
	if response.Code != 0 {
		fmt.Printf("请求失败: %s\n", response.Msg)
		return "", fmt.Errorf("请求失败: %s", response.Msg)
	}
	// fmt.Println(response)

	// 将 response.Data 转换为字符串
	if dataStr, ok := response.Data.(string); ok {
		return dataStr, nil
	}

	// 如果不是字符串，尝试转换为JSON字符串
	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return "", fmt.Errorf("转换响应数据失败: %v", err)
	}

	type Response struct {
		Token string `json:"token"`
	}

	// Parse the JSON
	var resp Response
	err = json.Unmarshal(jsonData, &resp)
	if err != nil {
		return "", fmt.Errorf("Error parsing JSON: %v", err)
	}

	return resp.Token, nil
}
