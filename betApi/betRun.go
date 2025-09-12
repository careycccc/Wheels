package betApi

import (
	"fmt"
	"project/betApi/winGo"
	"project/userApi/deskApi"
	"project/utils"
)

// 投注
// 输入用户名
func BetRun(userName string) {
	num := utils.RandmoNumber(2)
	gameCodeList := []string{"WinGo_5M", "TrxWinGo_10M"}
	gameCode := gameCodeList[num]
	num1 := utils.RandmoNumber(5)
	betContentList := []string{"Color_Green", "Color_Violet", "Color_Red", "BigSmall_Big", "BigSmall_Small"}
	betContent := betContentList[num1]
	num2 := utils.RandmoNumber(4)
	amountList := []int{10, 20, 50, 100}
	amount := amountList[num2]
	num3 := utils.RandmoNumber(4)
	betMultipleList := []int{1, 5, 10, 100}
	betMultiple := betMultipleList[num3]
	fmt.Println(gameCode, betContent, amount, betMultiple)
	result, err := deskApi.UserloginY1(userName, "qwer1234") // 前台登录 返回token值，后面的请求都需要这个token
	if err != nil {
		fmt.Println(result)
		return
	}

	// if len(result) > 0 {
	// 	winGo.ThirdGameFunc(result, gameCode)
	// }
	BalanceToken := winGo.GetBalanceInfoFunc(result, gameCode)
	// 是否可以投注
	isBet, issNumber := IsBet("", gameCode)
	if isBet && result != "-1" {
		// 可以投注
		BetWingo(gameCode, amount, betMultiple, betContent, issNumber, BalanceToken, userName)
	} else {
		fmt.Println("不可以投注")
		return
	}
}
