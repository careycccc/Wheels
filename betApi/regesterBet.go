package betApi

import (
	"fmt"
	payMoneyapi "project/payMoneyApi"
	"project/userApi/adminUser"
	memberlist "project/userApi/adminUser/MemberManagement/Memberlist"
	"project/userApi/deskApi"
	"project/utils"
	"sync"
)

// 负责组装注册投注
func RegesterBet(userName string) {
	// 进行前台注册
	token := deskApi.GeneralRegiterFunc(userName)
	if token == "" {
		fmt.Println("前台注册没有成功,没有获取到token")
		return
	}
	// 在后台查询这个人的userid
	userid := adminUser.GetUserApi(userName)
	// 后台修改密码
	if userid == 0 {
		fmt.Println("后台的userid没有获取到")
		return
	}
	// 充值金额
	manulMoneny, _ := utils.GenerateRandomInt(10000, 50000)
	var wg sync.WaitGroup
	// wg.Add(2)

	// 进行后台充值
	go payMoneyapi.ManualRecharge(userid, manulMoneny, 1, &wg)
	// 进行后台修改密码
	go memberlist.UpdataPasswordFunc(userid, "qwer1234", &wg)
	// 等待所有函数执行完成
	wg.Wait()
	// 进行前台登录 和 投注
	BetRun(userName)
}

func RegesterBetRandmo() {
	// 随机用户名
	userName, _ := utils.RandmoUserCount()
	// 进行前台注册
	RegesterBet(userName)
	fmt.Println("注册的账号", userName)
}
