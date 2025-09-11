package main

import (
	"fmt"
	"math/rand"
	"project/betApi"
	"project/betApi/winGo"
	"project/common"
	payMoneyapi "project/payMoneyApi"
	"project/userApi/adminUser"
	userApi "project/userApi/adminUser"
	"project/userApi/adminUser/actingModle/actingFy"
	_ "project/userApi/adminUser/activeModle"
	"project/userApi/deskApi"
	"project/utils"
	"sync"
	"time"
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

// 邀请转盘邀请下一级
/*
userAmount  邀请下级的账号
yqCode  邀请人的邀请码
monenyCount 邀请转盘的充值金额
**/
func RunWhille(userAmount string, yqCode string, monenyCount int64) {
	// 发送验证码
	actingFy.SendVerifiyCodeFunc(userAmount) // 发送验证码
	// 获取验证码
	time.Sleep(time.Second * 1)
	verifyCode := actingFy.QueryTifyFunc2() //获取验证码
	fmt.Println("当前的验证码", verifyCode)
	// 发送注册
	deskApi.RegisterFunc(userAmount, verifyCode, yqCode)
	// 后台登录后进行充值
	time.Sleep(time.Second * 1)

	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成一个[0, 1]范围内的随机数
	// randomNumber := rand.Intn(2)

	adminRun(userAmount, monenyCount)
	// 检查随机数以确定是否触发50%几率的事件
	// if randomNumber == 0 {
	// 	// 后台进行登录和人工充值
	// }
}

// 并行邀请人
// 任务函数，接收三个公共数据及其对应的锁
// 任务函数，调用 RunWhille 并更新公共数据
func TaskWhille(id int, wg *sync.WaitGroup, yqCode *string, monenyCount *int64, lock *sync.Mutex) {
	defer wg.Done()
	userAmount := utils.RandmoUserCount()
	// 使用单个锁保护对三个公共数据的联合操作
	lock.Lock()
	// 调用 RunWhille，传入当前公共数据值
	RunWhille(userAmount, *yqCode, *monenyCount)
	lock.Unlock()
}

// 运行并行的任务
func RunTaskWhille(yqCode string, monenyCount int64) {
	rand.Seed(time.Now().UnixNano()) // 初始化随机种子

	// 定义单个锁保护所有公共数据
	var lock sync.Mutex

	var wg sync.WaitGroup
	// 并发控制，使用通道限制最大并发数为
	semaphore := make(chan struct{}, 3)

	// 执行 几 次任务
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		semaphore <- struct{}{} // 获取一个令牌，控制并发数

		go func(id int) {
			defer func() { <-semaphore }() // 释放令牌
			// 将三个公共数据及其锁作为参数传入
			TaskWhille(id, &wg, &yqCode, &monenyCount, &lock)
		}(i)
	}

	// 等待所有任务完成
	wg.Wait()
}

func main() {
	adminUser.InitConfig()
	// userAmount := "919111997663" // 需要添加的用户账号
	// userAmount := utils.RandmoUserCount()
	// RunWhille(userAmount, "IW_6TXBN5N", 100)
	RunTaskWhille("IW_9SZ3X4N", 100)
	// deskRun(userAmount)  // 前台登录并进行了投注
	// adminRun(userAmount, 778)  // 后台进行登录和人工充值
	// adminUser.SendOneZnx() // 发送站内信
	// actingFy.RunInvite()
	// actingFy.SendVerifiyCodeFunc(userAmount) // 发送验证码
	// verifyCode := actingFy.QueryTifyFunc2() //获取验证码
	// fmt.Println("当前的验证码", verifyCode)
	// actingFy.GetInviteCodeFunc("919091997113") // 获取邀请码
	// deskApi.RegisterFunc(userAmount, "214537", "IW_YGN5QLN")

	// 总代注册
	// deskApi.GeneralRegiterFunc(userAmount)
	// 点击4个礼物盒
	// deskApi.InvitationCarousel(userAmount)
}
