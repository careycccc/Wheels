package adminUser

import (
	"fmt"
	"project/request"
	"project/utils"
)

type addUserInfoStruct struct {
	Account      string `json:"account"` // 添加的用户的账号91号码
	UserType     int8   `json:"userType"`
	PassWord     string `json:"password"`
	Remark       string `json:"remark"`
	RegisterType int8   `json:"registerType"`
}

// 添加用户
type AddUserStruct struct {
	AddUserList []addUserInfoStruct `json:"addUserList"`
	Random      int64               `json:"random"`
	Language    string              `json:"language"`
	Signature   string              `json:"signature"`
	Timestamp   int64               `json:"timestamp"`
}

// 发送添加用户的请求
func AddUserRequest(userAmount string) {
	//初始化这些添加用户的结构体
	api := "/api/Users/AddUsers"
	requesPayload := make(map[string]interface{})

	// 构建 addUserList 数组
	addUserList := []interface{}{
		map[string]interface{}{
			"account":      userAmount,
			"userType":     0,
			"password":     "qwer1234",
			"remark":       "",
			"registerType": 1,
		},
	}

	// 填充 payload
	requesPayload["addUserList"] = addUserList
	requesPayload["random"] = request.RandmoNie()
	requesPayload["language"] = "zh"
	requesPayload["signature"] = ""
	requesPayload["timestamp"] = request.GetNowTime()

	// 设置请求头，和获取token
	headMap, base_url := GetHeaderUrl()
	responBody, _, err := request.PostRequestCofig(requesPayload, base_url, api, headMap)
	if err != nil {
		fmt.Println("添加用户失败", err)
		return
	}
	result := utils.Unmarshal(string(responBody))
	fmt.Printf("添加用户的请求结果%v", result)
	if len(result["data"].([]interface{})) == 0 {
		fmt.Printf("添加成功%v", result["data"])
	} else {
		fmt.Printf("添加失败%v", result["data"])
	}

}
