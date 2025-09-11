package adminUser

import (
	"fmt"
	"project/common"
	"project/utils"
	"sync"
	"time"
)

var (
	Token_addr_yaml_local = "./yaml/token.yaml"
)

// 获取yaml中的token
type ConfigToken struct {
	Token string
}

// once 确保 InitConfig 只执行一次
var once sync.Once

// InitConfig 是只执行一次的函数
// 手动调用后，后续调用会直接返回，不会再执行函数体
func InitConfig() {
	once.Do(func() {
		// 这里放你实际想执行一次的逻辑，例如加载配置文件、初始化数据库等
		// fmt.Println("正在初始化配置... 后台登录中..。。！")
		token, err := Login("carey3003", "qwer1234")
		if err != nil {
			fmt.Println(err)
			return
		}
		// 把token写入到ymal中
		config := Config{}
		config.Token = token
		errs := utils.WriteYAML(Token_addr_yaml_local, &config)
		if errs != nil {
			// fmt.Printf("token写入失败%v", errs)
			fmt.Println(errs)
			return
		}
		// fmt.Printf("token写入成功.......\n")
	})
}

func GetToken() string {
	var config ConfigToken
	err := utils.ReadYAML(Token_addr_yaml_local, &config)
	if err != nil {
		fmt.Println("读取失败", err)
		return ""
	}
	// fmt.Printf("读取的内容%v", config.Token)
	time.Sleep(time.Millisecond * 100)
	return config.Token
}

// 获取tonken
// func GetToken() (string, error) {
// 	var common common.AdminUserName
// 	username := common.AdminUserInit().UserName
// 	pwd := common.AdminUserInit().Pwd
// 	// 获取token
// 	var config ConfigToken
// 	err := utils.ReadYAML(Token_addr_yaml_local, &config)
// 	if err != nil {
// 		return "", fmt.Errorf("读取失败%v", err)
// 	}
// 	// fmt.Printf("读取的内容%v", config.Token)
// 	n := 0
// 	for {
// 		if n > 3 || len(config.Token) > 0 {
// 			return config.Token, nil
// 		}
// 		if len(config.Token) == 0 {
// 			//读取的内容是空的，就发送登录请求
// 			_, err := Login(username, pwd)
// 			if err != nil {
// 				// fmt.Println(err)
// 				return "", err
// 			}
// 			return config.Token, nil
// 		}

// 		fmt.Printf("n的值==%v", n)
// 		n++
// 	}

// }

// 对token的获取和请求头的封装
func GetHeaderUrl() (map[string]interface{}, string) {
	var baseurl common.CofingURL
	base_url := baseurl.ConfigUrlInit().ADMIN_URL
	var head common.AdminHeaderAuthorizationConfig
	token := GetToken()
	headMap := head.AdminHeaderAuthorizationFunc(token)
	return headMap, base_url
}
