package deskApi

// // 单例模式
// import (
// 	"fmt"
// 	"sync"
// )

// // 前台一个账号登录进去后，可以进行其他的操作，无需重复登录
// // 需要传入一个账号
// func DeskCooperationRun(userName string) {
// 	// 模拟多个 goroutine 并行调用，代表不同操作
// 	var wg sync.WaitGroup
// 	operations := []func(){
// 		// 模拟多次调用登录（实际只执行一次）
// 		func() { fmt.Println("操作1: 获取 token:", GetToken(userName)) },
// 		// 点击4个礼物盒
// 		func() { fmt.Println("操作2:", ClickShareFunc(userName)) },
// 		// 点击分享按钮
// 		func() { ClickShareFunc(userName) },
// 		// func() { fmt.Println("操作3:", UpdateAvatar("https://example.com/avatar.png")) },
// 		// 再次查看详情
// 		// func() { fmt.Println("操作4:", tGetUserProfile()) },
// 		// 再次获取 token
// 		// func() { fmt.Println("操作5: 获取 token:", tokenpkg.GetToken()) },
// 	}

// 	// 并行执行操作
// 	for i, op := range operations {
// 		wg.Add(1)
// 		go func(id int, fn func()) {
// 			defer wg.Done()
// 			// fmt.Printf("Goroutine %d 开始...\n", id)
// 			fn()
// 		}(i, op)
// 	}

// 	wg.Wait()
// 	fmt.Println("所有操作完成，登录只执行一次。")
// }
