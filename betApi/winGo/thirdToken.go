package winGo

import (
	"fmt"
	"project/common"
	"project/request"
	"project/utils"
	"strings"
)

// 主要是针对那三个get请求获取token，为后面的投注的token
//  gameCode=WinGo_5M&language=en&random=131601285634&signature=C9A95C24297A0345B9DF3FC970BBB766&timestamp=1756999407

type GetGameInfoStruct struct {
	GameCode  string `json:"gameCode"`
	Language  string `json:"language"`
	Random    int64  `json:"random"`
	Signature string `json:"signature"`
	Timestamp int64  `json:"timestamp"`
}

// 请求GetGameInfo并且返回token
func ThridTokenFunc(tokenUser, gameCode string) string {
	// 初始化结构体并且赋值
	GetGameInfo := &GetGameInfoStruct{}
	values := []interface{}{gameCode, "en", request.RandmoNie(), "", request.GetNowTime()}
	paramsMap := common.InitStructToMap(GetGameInfo, values)
	// 获取签名
	verifyPwd := ""
	signatureStr := utils.GetSignature(paramsMap, &verifyPwd)
	paramsMap["signature"] = signatureStr

	api := "/api/Lottery/GetGameInfo"
	baseUrl := common.LOTTERY_H5
	// 获取请求头
	deskA := &common.BetTokenStruct{}
	url_h5 := common.WMG_H5
	token := ThirdGameFunc(tokenUser, gameCode)
	desSlice := []interface{}{url_h5, url_h5, token}
	headMap, _ := common.AssignSliceToStructMap(deskA, desSlice)
	_, resp, err := request.GetRequest(baseUrl, api, headMap, paramsMap)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	authorization := resp.Header.Get("Authorization")
	// 去掉前缀 "Bearer "
	cleanToken := strings.TrimPrefix(authorization, "Bearer ")
	// fmt.Println("响应头的token", cleanToken)
	return cleanToken
}

type GetBalanceInfoStruct struct {
	Language  string `json:"language"`
	Random    int64  `json:"random"`
	Signature string `json:"signature"`
	Timestamp int64  `json:"timestamp"`
}

// 请求GetBalance
func GetBalanceInfoFunc(tokenUser, gameCode string) string {
	// 初始化结构体并且赋值
	GetGameInfo := &GetBalanceInfoStruct{}
	values := []interface{}{"en", request.RandmoNie(), "", request.GetNowTime()}
	paramsMap := common.InitStructToMap(GetGameInfo, values)
	// 获取签名
	verifyPwd := ""
	signatureStr := utils.GetSignature(paramsMap, &verifyPwd)
	paramsMap["signature"] = signatureStr

	api := "/api/Lottery/GetBalance"
	baseUrl := common.LOTTERY_H5
	// 获取请求头
	deskA := &common.BetTokenStruct{}
	url_h5 := common.WMG_H5
	token := ThridTokenFunc(tokenUser, gameCode)
	desSlice := []interface{}{url_h5, url_h5, token}
	headMap, _ := common.AssignSliceToStructMap(deskA, desSlice)
	_, resp, err := request.GetRequest(baseUrl, api, headMap, paramsMap)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	authorization := resp.Header.Get("Authorization")
	// 去掉前缀 "Bearer "
	cleanToken := strings.TrimPrefix(authorization, "Bearer ")
	// fmt.Println("响应头的tokenBanlne-----", cleanToken)
	return cleanToken
}
