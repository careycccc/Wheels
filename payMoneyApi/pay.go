package payMoneyapi

import (
	"fmt"
	"project/request"
	"project/userApi/adminUser"
)

// 人工充值
type manualRecharge struct {
	ArtificialRechargeType int8   `json:"artificialRechargeType"`
	RechargeAmount         int64  `json:"rechargeAmount"` // 充值金额
	Remark                 string `json:"remark"`         // 备注
	AmountOfCode           int8   `json:"amountOfCode"`   // 打码量 null表示默认，数字表示倍数
	UserId                 int64  `json:"userId"`
	Random                 int64  `json:"random"`
	Language               string `json:"language"`
	Signature              string `json:"signature"`
	Timestamp              int64  `json:"timestamp"`
}

/*
*
userid 用户id
rechargeAmount 充值金额
amountOfCode 打码量
*/
func ManualRecharge(userid, rechargeAmount int64, amountOfCode int8) {
	api := "/api/ArtificialRechargeRecord/ArtificialRecharge"
	manualRechargeInit := manualRecharge{
		ArtificialRechargeType: 3,
		RechargeAmount:         rechargeAmount,
		Remark:                 "1",
		AmountOfCode:           amountOfCode,
		UserId:                 userid,
		Random:                 request.RandmoNie(),
		Language:               "en",
		Signature:              "",
		Timestamp:              request.GetNowTime(),
	}

	manualRechargedata := map[string]interface{}{
		"artificialRechargeType": manualRechargeInit.ArtificialRechargeType,
		"rechargeAmount":         manualRechargeInit.RechargeAmount,
		"remark":                 manualRechargeInit.Remark,
		"amountOfCode":           manualRechargeInit.AmountOfCode,
		"userId":                 manualRechargeInit.UserId,
		"random":                 manualRechargeInit.Random,
		"language":               manualRechargeInit.Language,
		"signature":              manualRechargeInit.Signature,
		"timestamp":              manualRechargeInit.Timestamp,
	}

	headMap, base_url := adminUser.GetHeaderUrl()
	resp, _, err := request.PostRequestCofig(manualRechargedata, base_url, api, headMap)
	if err != nil {
		fmt.Println("人工充值发送请求失败")
		return
	}
	fmt.Printf("充值结果%v", string(resp))
}
