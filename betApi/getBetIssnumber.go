package betApi

import (
	"encoding/json"
	"fmt"
	"project/common"
	"project/request"
	"time"
)

type Response struct {
	Code int `json:"code"`
	Data struct {
		StartTime      int64  `json:"startTime"`
		EndTime        int64  `json:"endTime"`
		IssueNumber    string `json:"issueNumber"`
		IntervalMinute int64  `json:"intervalMinute"`
		GameCode       string `json:"gameCode"`
		Diif           int    `json:"diif"`
		Countdown      int    `json:"countdown"`
	} `json:"data"`
	Msg        string `json:"msg"`
	MsgCode    int    `json:"msgCode"`
	ServerTime int64  `json:"serverTime"`
}

// 获取当期的期号
func GetNowBetNumber(token, betType string) (map[string]interface{}, error) {
	api := "/webapi/kv/issue/" + betType
	// 获取路径地址
	var baseUrl common.CofingURL
	base_url := baseUrl.ConfigUrlInit().Iss_URL
	// 获取请求头
	var issNumber common.GetIssNunmberHeaderConfig
	headMap := issNumber.GetIssNunmberHeaderFunc(token, betType)
	k := make(map[string]interface{})
	respBody, _, err := request.GetRequest(base_url, api, headMap, k)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}
	// fmt.Printf("响应期号%v", string(resp))
	var response Response
	error := json.Unmarshal([]byte(string(respBody)), &response)
	if error != nil {
		fmt.Printf("响应期号解析失败%v", error)
		return nil, error
	}
	nowBetNumber := map[string]interface{}{
		"startTime":      response.Data.StartTime,      // 开始时间
		"endTime":        response.Data.EndTime,        // 结束时间
		"issueNumber":    response.Data.IssueNumber,    // 期号
		"intervalMinute": response.Data.IntervalMinute, // 间隔时间
	}
	// fmt.Println(nowBetNumber)
	return nowBetNumber, nil
}

// 判断是否可以下注,并且返回期号
func IsBet(token, betType string) (bool, string) {
	nowBetNumber, err := GetNowBetNumber(token, betType)
	if err != nil {
		fmt.Println("没有成功获取到期号")
		return false, "-1"
	}
	// startTime := nowBetNumber["startTime"].(int64)
	endTime := nowBetNumber["endTime"].(int64)
	issueNumber := nowBetNumber["issueNumber"]
	// intervalMinute := nowBetNumber["intervalMinute"].(int64)
	// fmt.Println("结束 - 开始", endTime-startTime)
	// 获取当前时间
	now := time.Now()

	// 获取时间戳（秒）
	secTimestamp := now.UnixMilli()
	// fmt.Printf("当前时间戳（秒）: %d\n", secTimestamp)
	// fmt.Println("结束时间", endTime)

	// fmt.Printf("间隔时间%v,%v\n", intervalMinute*1000, endTime-secTimestamp)
	// 结束时间 - 当前时间 >= 动画7s
	if endTime-secTimestamp >= 7000 {
		// 可以投注
		return true, issueNumber.(string)
	} else {
		// 不可以投注
		// 需要等待7s
		time.Sleep(time.Second * 7)
		return IsBet(token, betType)
	}
}
