package betApi

import (
	"encoding/json"
	"fmt"
	"project/common"
	"project/request"
)

// BetRequest 定义请求体的结构体
type BetRequest struct {
	GameCode    string `json:"gameCode"`
	IssueNumber string `json:"issueNumber"`
	Amount      int    `json:"amount"`
	BetMultiple int    `json:"betMultiple"`
	BetContent  string `json:"betContent"`
	Language    string `json:"language"`
	Random      int64  `json:"random"`
	Signature   string `json:"signature"`
	Timestamp   int64  `json:"timestamp"`
}

type ResponseStruct struct {
	Code        int
	Msg         string
	MsgCode     int
	ServiceTime int64
}

/*
*
gameCode  彩票投注种类
amount 投注金额 = 单个金额 * 倍率
betMultiple 投注倍率
betContent 投注盘口
issueNumber 期号
token token对象
*/
func BetWingo(gameCode string, amount, betMultiple int, betContent, issueNumber, token, username string) {
	var apiArg string
	if gameCode == "TrxWinGo_10M" {
		apiArg = "TrxWinGoBet"
	} else {
		apiArg = "WinGoBet"
	}
	// 请求体地址
	api := "/api/Lottery/" + apiArg
	// url := "https://sit-lotteryh5.wmgametransit.com"
	url := common.LOTTERY_H5
	// 参数化
	bet := &BetRequest{}
	betResultList := []interface{}{gameCode, issueNumber, amount, betMultiple, betContent, "en", request.RandmoNie(), "", request.GetNowTime()}
	resultMap, _ := common.InitStructToMap(bet, betResultList)
	// 获取请求头
	deskA := &common.BetTokenStruct{}
	url_h5 := common.WMG_H5
	desSlice := []interface{}{url_h5, url_h5, token}
	headMap, _ := common.AssignSliceToStructMap(deskA, desSlice)
	respBody, _, err := request.PostRequestCofig(resultMap, url, api, headMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	var res ResponseStruct
	err = json.Unmarshal([]byte(string(respBody)), &res)
	if err != nil {
		fmt.Println("投注的反序列失败", err)
		return
	}
	code := res.Code
	msgcode := res.MsgCode
	if code == 0 && msgcode == 0 {
		fmt.Printf("%v在%v投注了%v成功\n", username, gameCode, amount*betMultiple)
	} else {
		fmt.Println(res)
	}
}
