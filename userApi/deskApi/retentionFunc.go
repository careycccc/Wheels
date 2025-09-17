package deskApi

import "fmt"

// 留存功能
// 留存第一天 需要注册和充值或投注
func RetentionFisterDay(userName string) {
	// 总代的方式进行注册
	token := GeneralRegiterFunc(userName)
	var result chan string = make(chan string)
	result <- token
	fmt.Println("注册成功，获取到的token", <-result)
	close(result)
}
