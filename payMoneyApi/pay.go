package payMoneyapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"project/request"
	"project/userApi/adminUser"
	"sync"
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

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

/*
*
userid 用户id
rechargeAmount 充值金额
amountOfCode 打码量
*/
func ManualRecharge(userid, rechargeAmount int64, amountOfCode int8, wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()
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
		err := errors.New("人工充值发送请求失败,原因" + error.Error(err))
		return err
	}
	var response Response
	error := json.Unmarshal([]byte(resp), &response)
	if error != nil {
		err := errors.New("人工充值反序劣化失败,原因" + error.Error())
		return err
	}
	if response.Code != 0 {
		err := errors.New("人工充值失败,原因" + response.Msg)
		return err
	}
	fmt.Printf("充值结果%v,充值金额%v\n", string(resp), rechargeAmount)
	return nil
}
