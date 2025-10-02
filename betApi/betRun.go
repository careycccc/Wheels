package betApi

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"project/betApi/winGo"
	"project/request"
	"project/userApi/deskApi"
	"sync"
	"time"
)

func marshalToken(jsonStr string) string {
	// 定义结构体来映射 JSON
	type Response struct {
		Token string `json:"token"`
	}

	// 解析 JSON
	var resp Response
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		fmt.Println("是这里解析出错了吗？？？")
		log.Fatalf("Error parsing JSON: %v", err)
		return ""
	}
	return resp.Token
}

// RandomNumber 为每个 goroutine 生成 [0, n) 范围内的随机整数
func RandomInt(n int) int {
	if n <= 0 {
		return 0 // 处理无效输入
	}
	// 使用 sync.Once 确保每个 goroutine 初始化一次随机源
	var src *rand.Rand
	var once sync.Once
	once.Do(func() {
		src = rand.New(rand.NewSource(time.Now().UnixNano())) // 为 goroutine 创建随机源
	})
	return src.Intn(n) // 生成 [0, n) 范围的随机数
}

// 投注
// 输入用户名
func BetRun(userName string) error {
	time.Sleep(100 * time.Millisecond)
	num := RandomInt(2)
	gameCodeList := []string{"WinGo_5M", "TrxWinGo_10M"}
	gameCode := gameCodeList[num]
	num1 := RandomInt(5)
	betContentList := []string{"Color_Green", "Color_Violet", "Color_Red", "BigSmall_Big", "BigSmall_Small"}
	betContent := betContentList[num1]
	num2 := RandomInt(4)
	amountList := []int{10, 20, 50, 100}
	amount := amountList[num2]
	num3 := RandomInt(4)
	betMultipleList := []int{10, 20, 50, 100}
	betMultiple := betMultipleList[num3]
	// fmt.Println(gameCode, betContent, amount, betMultiple)
	// result, err := deskApi.UserloginY1(userName, "qwer1234") // 前台登录 返回token值，后面的请求都需要这个token
	// if err != nil {
	// 	return
	// }
	//调用实例

	token, err := request.RetryOperationWithResult(request.Func2WithResult(deskApi.UserloginY1), userName, "qwer1234")
	if err != nil {
		//fmt.Printf("UserloginY1 failed: %v\n", err)
		return fmt.Errorf("UserloginY1 failed: %v\n", err)
	}
	//fmt.Println("投注重试的结果", result)
	// token := marshalToken(result.(string))
	// if token == "" {
	// 	fmt.Println("token没有获取到")
	// 	return fmt.Errorf("token没有获取到")
	// }
	// if len(result) > 0 {
	// 	winGo.ThirdGameFunc(result, gameCode)
	// }
	BalanceToken, balance := winGo.GetBalanceInfoFunc(token.(string), gameCode)
	if balance == 0.0 {
		//fmt.Println("------------------------余额为0,不可以投注------------------------")
		return fmt.Errorf("------------------------余额为0,不可以投注------------------------")
	}
	// 是否可以投注
	isBet, issNumber := IsBet("", gameCode)
	if isBet && len(token.(string)) > 0 {
		// 可以投注
		BetWingo(gameCode, amount, betMultiple, betContent, issNumber, BalanceToken, userName)
	} else {
		//fmt.Println("不可以投注")
		return fmt.Errorf("不可以投注")
	}
	return nil
}
