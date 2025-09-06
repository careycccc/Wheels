package adminUser

import (
	"fmt"
	"project/common"
	"project/utils"
)

var (
	Token_addr_yaml_local = "./yaml/token.yaml"
)

// 获取yaml中的token
type ConfigToken struct {
	Token string
}

// 获取tonken
func GetToken() (string, error) {
	var common common.AdminUserName
	username := common.AdminUserInit().UserName
	pwd := common.AdminUserInit().Pwd
	// 获取token
	var config ConfigToken
	err := utils.ReadYAML(Token_addr_yaml_local, &config)
	if err != nil {
		return "", fmt.Errorf("读取失败%v", err)
	}
	// fmt.Printf("读取的内容%v", config.Token)
	n := 0
	for {
		if n > 3 || len(config.Token) > 0 {
			return config.Token, nil
		}
		if len(config.Token) == 0 {
			//读取的内容是空的，就发送登录请求
			err := Login(username, pwd)
			if err != nil {
				// fmt.Println(err)
				return "", err
			}
			return config.Token, nil
		}

		fmt.Printf("n的值==%v", n)
		n++
	}

}

// 对token的获取和请求头的封装
func GetHeaderUrl() (map[string]interface{}, string) {
	var baseurl common.CofingURL
	base_url := baseurl.ConfigUrlInit().ADMIN_URL
	// 获取token
	token, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return nil, ""
	}
	var head common.AdminHeaderAuthorizationConfig
	headMap := head.AdminHeaderAuthorizationFunc(token)
	return headMap, base_url
}
