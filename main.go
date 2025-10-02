package main

import (
	"fmt"
	"os"
	adminUser "project/userApi/adminUser"
	"project/userApi/adminUser/activeModle"
	// activeModle "project/userApi/adminUser/activeModle"
)

/*
后台启动 传入的账号是系统没有的账号
userAmount string 传入一个账号进行充值
moneny int64 充值金额
*
*/
// func adminRun(userAmount string, moneny int64) {
// 	var common common.AdminUserName
// 	username := common.AdminUserInit().UserName
// 	pwd := common.AdminUserInit().Pwd
// 	userApi.Login(username, pwd) // 商户后台登录
// 	// userApi.AddUserRequest(userAmount) // 添加用户
// 	userid := userApi.GetUserApi(userAmount) // 获取用户id
// 	if userid == -1 {
// 		return
// 	}
// 	// WaitGroup
// 	var wg sync.WaitGroup
// 	payMoneyapi.ManualRecharge(userid, moneny, 0, &wg) // 用户充值
// }

func main() {
	adminUser.InitConfig()
	// userAmount := "919162190451"           // 需要添加的用户账号
	// deskApi.RetentionFisterDay(userAmount) // 留存第一天 需要注册和充值或投注
	// deskApi.InvitationCarousel(userAmount, 120)

	// userAmount := utils.RandmoUserCount()
	// deskApi.RunWhille(userAmount, "IW_2437757", 100)
	// 只要填写邀请码，自动邀请下级，并且充值 但是不会点击4个礼物盒子
	// deskApi.RunTaskWhille("IW_Z8GYH5N", 195)
	// deskRun(userAmount)  // 前台登录并进行了投注
	// adminRun(userAmount, 778)  // 后台进行登录和人工充值
	// 多线程发送站内信
	// var wg sync.WaitGroup
	// wg.Add(1)
	// for i := 0; i <= 1; i++ {
	// 	go activeModle.SendOneZnx(&wg) // 发送站内信
	// }
	// wg.Wait()
	activeModle.SendOneZnx() // 发送站内信

	// actingFy.RunInvite()
	// actingFy.SendVerifiyCodeFunc(userAmount) // 发送验证码
	// time.Sleep(time.Second * 4)
	// verifyCode := actingFy.QueryTifyFunc2() //获取验证码
	// fmt.Println("当前的验证码", verifyCode)
	// // actingFy.GetInviteCodeFunc("919091997113") // 获取邀请码
	// deskApi.RegisterFunc(userAmount, verifyCode, "IW_YGN5QLN")

	// 总代注册
	// deskApi.GeneralRegiterFunc(userAmount)
	// 点击4个礼物盒
	// deskApi.InvitationCarousel(userAmount)

	// betApi.BetRun(userAmount)
	// memberlist.UpdataPasswordFunc(2437103, "qwer1234")  // 修改密码
	// betApi.RegesterBet(userAmount) // 前台进行注册，后台充值修改密码后，进行投注

	// actingFy.QueryTifyFunc("919111997678")
	// str, _ := deskApi.UserloginY1("919111997678", "qwer1234")
	// fmt.Println(str)
	// request.GetProxyIpFunc()
	// deskApi.GeneralRegiterFuncProxyRun()
	// betApi.RegesterBetRandmo()

	// 随机生成10个用户id，进行注册，充值，投注
	// idlist := utils.RandmoUserId(15)
	// //  打印几个示例ID
	// for i := 0; i < len(idlist); i++ {
	// 	fmt.Printf("Example ID: %s\n", idlist[i])
	// 	time.Sleep(time.Second * 1)
	// 	go AppendToYAML(idlist[i])
	// 	betApi.RegesterBet(idlist[i])
	// }

	// 随机10个总代进行邀请转盘的邀请
	// zd := 10
	// idlist := utils.RandmoUserId(zd)
	// for i := 0; i < len(idlist); i++ {
	// 	fmt.Printf("++++++++++++++++++++++++Example ID++++++++++++++++++++: %s\n", idlist[i])
	// 	time.Sleep(time.Second * 1)
	// 	AppendToYAML(idlist[i])
	// 	monery, _ := utils.GenerateRandomInt(5000, 15000)
	// 	deskApi.InvitationCarousel(idlist[i], monery)
	// }

	//

	// 从22.yaml文件中读取账号进行投注
	// 调用封装函数，传入 YAML 文件路径、betApi.BetRun 和并发量
	// err := common.ProcessAccounts("./22.yaml", betApi., 2)
	// if err != nil {
	// 	fmt.Printf("处理失败: %v", err)
	// 	return
	// }
	// time.Sleep(time.Second * 20)
	// lines, err := utils.ReadYAMLByLine("./44.yaml")
	// if err != nil {
	// 	println("读取YAML文件失败:", err.Error())
	// 	return
	// }

	// // 打印每一行
	// for _, line := range lines {
	// 	lin := strings.TrimSpace(line)
	// 	betApi.BetRun(lin)
	// }

	// 把22.yaml的人作为总代邀请下一级
	// if lines, err := utils.ReadYAMLByLine("./33.yaml"); err != nil {
	// 	println("读取YAML文件失败:", err.Error())
	// 	return
	// } else {
	// 	for _, line := range lines {
	// 		monery, _ := utils.GenerateRandomInt(5000, 15000)
	// 		// 总代邀请下一级 并且把下写入到33.yam中
	// 		deskApi.InvitationCarousel(line, monery)
	// 	}
	// }

	// 邀请转盘的自动分析结果
	// id := "IW250917033634169kRaIDOO3Wx" // 订单号
	// activeModle.WheelAutoResultFunc(id)
	// actingFy.GetUserTypeTotal(0, 2)

	// if result, err := deskApi.UserloginY1("919244104367", "qwer1234"); err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {

	// 	fmt.Println("登录结果", result)
	// }

}

// AppendToYAML 将数据追加写入到 1.yaml 文件，每行一个数据
func AppendToYAML(data ...interface{}) error {
	file, err := os.OpenFile("./1.yaml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, v := range data {
		// 将数据转换为字符串并写入，带换行符
		_, err := file.WriteString(fmt.Sprintf("%v\n", v))
		if err != nil {
			return err
		}
	}
	return nil
}
