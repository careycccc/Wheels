package deskApi

// import (
// 	"fmt"
// 	"sync"
// )

// // 全局变量存储 token
// var token string
// var once sync.Once

// // initOnce 执行一次登录，设置全局 token
// func initOnce(userName string) {
// 	token = GeneralRegiterFunc(userName)
// 	fmt.Println("登录完成，获取 token:", token)
// }

// // GetToken 获取 token，确保只登录一次
// func GetToken(userName string) string {
// 	once.Do(func() {
// 		initOnce(userName) // 在闭包中调用 initOnce，传递 userName
// 	})
// 	return token
// }
