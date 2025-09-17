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
	time.Sleep(time.Second * 2)
	// 点击免费的旋转次数
	ClickFreeWheelFunc(token)
	time.Sleep(time.Second * 1)
	// 已经点击了，可以点击分享按钮
	invital := ClickShareFunc(userName, token)
	if invital == "" {
		fmt.Println("邀请码为空", invital)
		return
	}
	fmt.Println("邀请码为", invital)
	RunTaskWhille(invital, moneny, token)
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
	// fmt.Println("当前的验证码", verifyCode)
	// 发送注册
	RegisterFunc(userAmount, verifyCode, yqCode)
	// 后台登录后进行充值
	time.Sleep(time.Second * 2)
	adminRun(userAmount, monenyCount)
	// 初始化随机数种子
	// rand.Seed(time.Now().UnixNano())
	// // 生成一个[0, 1]范围内的随机数
	// randomNumber := rand.Intn(2)
	// // 检查随机数以确定是否触发50%几率的事件
	// if randomNumber == 0 {
	// 	// 后台进行登录和人工充值
	// 	adminRun(userAmount, monenyCount)
	// }
}

// 并行邀请人
// 任务函数，接收三个公共数据及其对应的锁
// 任务函数，调用 RunWhille 并更新公共数据
func TaskWhille(id int, wg *sync.WaitGroup, yqCode *string, monenyCount *int64, lock *sync.Mutex) {
	defer wg.Done()
	userAmount, _ := utils.RandmoUserCount()
	// 使用单个锁保护对三个公共数据的联合操作
	lock.Lock()
	// 调用 RunWhille，传入当前公共数据值
	RunWhille(userAmount, *yqCode, *monenyCount)
	lock.Unlock()
}

// 运行并行的任务
// 只要填写邀请码，自动邀请下级，并且充值
func RunTaskWhille(yqCode string, monenyCount int64, token string) {
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
	// 充值结束后等待5s后，进行点击提现

	// 还差提现金额的获取
	time.Sleep(time.Second * 5)
	_, err := ClickWheelWithdrawFunc(499, token) // 点击转盘提现
	if err != nil {
		fmt.Println("点击转盘提现失败", err)
		return
	}

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
	var wg sync.WaitGroup
	// 通道收集结果
	payMoneyapi.ManualRecharge(userid, moneny, 0, &wg) // 用户充值
	// wg.Wait()

}

// 使用代理的总代注册 并且记录ip和账号 和随机的账号
// func GeneralRegiterFuncProxyRun() {
// 	//每次都需要获取新的ip地址
// 	//request.ProxyFunc()
// 	UserIpMap := make(map[string]string)
// 	// 随机一个账号
// 	userName, _ := utils.RandmoUserCount()
// 	// 使用这个账号进行注册
// 	token, ip := GeneralRegiterFuncProxy(userName)
// 	if token == "" {
// 		fmt.Println("前台注册没有成功,没有获取到token")
// 		return
// 	}
// 	// 把ip和账号放着一个map中
// 	UserIpMap[ip] = userName
// 	fmt.Println("当前的ip和账号", UserIpMap)
// }
