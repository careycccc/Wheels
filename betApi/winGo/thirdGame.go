package winGo

import (
	"encoding/json"
	"fmt"
	"project/common"
	"project/request"
	"project/userApi"
	"project/utils"
	"regexp"
)

type ThirdGameStruct struct {
	GameCode     string `json:"gameCode"`
	VendorCode   string `json:"vendorCode"`
	GameId       int64  `json:"gameId"`
	ReturnUrl    string `json:"returnUrl"`
	DeviceType   string `json:"deviceType"`
	DeviceTypeId string `json:"deviceTypeId"`
	Language     string `json:"language"`
	Random       int64  `json:"random"`
	Signature    string `json:"signature"`
	Timestamp    int64  `json:"timestamp"`
}

/*
token 传入登录成功后台的用户的token
gameCode WinGo_5M
return token
*
*/
func ThirdGameFunc(token, gameCode string) string {
	api := "/api/ThirdGame/GetGameUrl"
	var baseUrl common.CofingURL
	base_url := baseUrl.ConfigUrlInit().H5_URL
	// 初始化结构体
	thirdGameStructinit := ThirdGameStruct{
		GameCode:     gameCode,
		VendorCode:   "ARLottery",
		GameId:       10003,
		ReturnUrl:    "https://sit-plath5-y1.mggametransit.com/game?categoryCode=C202505280608510046",
		DeviceType:   "PC",
		DeviceTypeId: utils.GenerateCryptoRandomString(32),
		Language:     "en",
		Random:       request.RandmoNie(),
		Signature:    "",
		Timestamp:    request.GetNowTime(),
	}
	thirdGameMap := map[string]interface{}{
		"gameCode":     thirdGameStructinit.GameCode,
		"vendorCode":   thirdGameStructinit.VendorCode,
		"gameId":       thirdGameStructinit.GameId,
		"returnUrl":    thirdGameStructinit.ReturnUrl,
		"deviceType":   thirdGameStructinit.DeviceType,
		"deviceTypeId": thirdGameStructinit.DeviceTypeId,
		"language":     thirdGameStructinit.Language,
		"random":       thirdGameStructinit.Random,
		"signature":    thirdGameStructinit.Signature,
		"timestamp":    thirdGameStructinit.Timestamp,
	}

	// 获取请求头
	deskA := &common.DeskHeaderAstruct{}
	url_h5 := common.PLANT_H5
	desSlice := []interface{}{url_h5, url_h5, url_h5, token}
	headMap, _ := common.AssignSliceToStructMap(deskA, desSlice)

	// 发送请求
	resp, _, err := request.PostRequestCofig(thirdGameMap, base_url, api, headMap)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// fmt.Println("thirdgame响应结果", string(resp))

	var response userApi.Response
	err = json.Unmarshal([]byte(string(resp)), &response)
	if err != nil {
		fmt.Println("thirdgame响应结果反序列化失败", err)
		return ""
	}
	result := response.Data.(map[string]interface{})["url"]
	// fmt.Println(result)
	// 寻找token
	// 查找第一个匹配
	res := result.(string)
	// 用正则匹配 Token 的值
	re := regexp.MustCompile(`Token=([^&]+)`)
	matches := re.FindStringSubmatch(res)

	if len(matches) > 1 {
		// fmt.Println("Token:", matches[1])
		return matches[1]
	} else {
		fmt.Println("Token not found")
		return ""
	}
}
