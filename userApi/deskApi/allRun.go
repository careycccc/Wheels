package deskApi

import (
	"fmt"
	"math/rand"
	payMoneyapi "project/payMoneyApi"
	"project/userApi/adminUser"
	"project/userApi/adminUser/actingModle/actingFy"
	"project/utils"
	"sync"
	"time"
)

// 这里负责组装前台的一些功能

/*
总代注册，点击了邀请转盘，并且点击分享，下级人数进行自动邀请
需要传入总代的一个账号，和下级充值的金额
*
*/
func InvitationCarousel(userName string, moneny int64) {
	token := GeneralRegiterFunc(userName)
	// 总代进行注册登录，并且进行点击了邀请转盘的4个礼物盒
	ClickWheelFunc(userName, token)
	// fmt.Println("点击的礼物盒子状态有问题的状态", bool)
	// if !bool {
	// 	// 没有点击
	// 	fmt.Println("点击的礼物盒子状态有问题")
	// 	return
	// }
	// 已经点击了，可以点击分享按钮
	invital := ClickShareFunc(userName, token)
	if invital == "" {
		fmt.Println("邀请码为空", invital)
		return
	}
	fmt.Println("邀请码为", invital)
	RunTaskWhille(invital, moneny)
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
	time.Sleep(time.Second * 2)
	verifyCode := actingFy.QueryTifyFunc2() //获取验证码
	fmt.Println("当前的验证码", verifyCode)
	// 发送注册
	RegisterFunc(userAmount, verifyCode, yqCode)
	// 后台登录后进行充值
	time.Sleep(time.Second * 2)
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	// 生成一个[0, 1]范围内的随机数
	randomNumber := rand.Intn(2)
	// 检查随机数以确定是否触发50%几率的事件
	if randomNumber == 0 {
		// 后台进行登录和人工充值
		adminRun(userAmount, monenyCount)
	}
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

/*
后台启动 传入的账号是系统没有的账号
userAmount string 传入一个账号进行充值
moneny int64 充值金额
*
*/
func adminRun(userAmount string, moneny int64) {
	// var common common.AdminUserName
	// username := common.AdminUserInit().UserName
	// pwd := common.AdminUserInit().Pwd
	// userApi.Login(username, pwd) // 商户后台登录
	// userApi.AddUserRequest(userAmount) // 添加用户
	userid := adminUser.GetUserApi(userAmount) // 获取用户id
	if userid == -1 {
		return
	}
	results := make(chan string, 2)
	var wg sync.WaitGroup
	// 通道收集结果
	payMoneyapi.ManualRecharge(userid, moneny, 0, &wg, results) // 用户充值
}
