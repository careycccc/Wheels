package main

import (
	"fmt"
	"project/betApi"
	"project/betApi/winGo"
	"project/common"
	payMoneyapi "project/payMoneyApi"
	userApi "project/userApi/adminUser"
	"project/userApi/adminUser/actingModle/actingFy"
	_ "project/userApi/adminUser/activeModle"
	"project/userApi/deskApi"
)

/*
后台启动 传入的账号是系统没有的账号
userAmount string 传入一个账号进行充值
moneny int64 充值金额
*
*/
func adminRun(userAmount string, moneny int64) {
	var common common.AdminUserName
	username := common.AdminUserInit().UserName
	pwd := common.AdminUserInit().Pwd
	userApi.Login(username, pwd) // 商户后台登录
	// userApi.AddUserRequest(userAmount) // 添加用户
	userid := userApi.GetUserApi(userAmount) // 获取用户id
	if userid == -1 {
		return
	}
	payMoneyapi.ManualRecharge(userid, moneny, 0) // 用户充值
}

/*
前台启动
userAmount string 传入登录的账号

*
*/
func deskRun(userAmount string) {
	gameCode := "WinGo_5M"                                     // 彩票类型
	betContent := "BigSmall_Big"                               // 投注盘口
	result, err := deskApi.UserloginY1(userAmount, "qwer1234") // 前台登录 返回token值，后面的请求都需要这个token
	if err != nil {
		fmt.Println(result)
		return
	}
	// if len(result) > 0 {
	// 	winGo.ThirdGameFunc(result, gameCode)
	// }
	BalanceToken := winGo.GetBalanceInfoFunc(result, gameCode)
	// 是否可以投注
	isBet, issNumber := betApi.IsBet("", gameCode)
	if isBet && result != "-1" {
		// 可以投注
		betApi.BetWingo(gameCode, 10, 2, betContent, issNumber, BalanceToken, userAmount)
	} else {
		fmt.Println("不可以投注")
		return
	}
}

func main() {
	// userAmount := "918281997445" // 需要添加的用户账号
	// // deskRun(userAmount)  // 前台登录并进行了投注
	// adminRun(userAmount, 778)
	// adminUser.SendOneZnx()
	// actingFy.RunInvite()
	actingFy.SendVerifiyCodeFunc("91919081997668")
}
