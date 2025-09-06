package adminUser

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"project/request"
)

type getUserApistruct struct {
	Account   string `json:"account"` // 用户账号
	PageNo    int8   `json:"pageNo"`  // 页码
	PageSize  int8   `json:"pageSize"`
	OrderBy   string `json:"orderBy"`
	Random    int64  `json:"random"`
	Language  string `json:"language"`
	Signature string `json:"signature"`
	Timestamp int64  `json:"timestamp"`
}

// 提取userid
type Useridstruct struct {
	Data struct {
		List []struct {
			UserId int64 `json:"userId"`
		} `json:"list"`
	} `json:"data"`
}

func GetUserApi(account string) int64 {
	api := "/api/Users/GetPageList"
	userapiInit := getUserApistruct{
		Account:   account,
		PageNo:    1,
		PageSize:  20,
		OrderBy:   "Desc",
		Random:    request.RandmoNie(),
		Language:  "zh",
		Signature: "",
		Timestamp: request.GetNowTime(),
	}
	userapiMap := map[string]interface{}{
		"account":   userapiInit.Account,
		"pageNo":    userapiInit.PageNo,
		"pageSize":  userapiInit.PageSize,
		"orderBy":   userapiInit.OrderBy,
		"random":    userapiInit.Random,
		"language":  userapiInit.Language,
		"signature": userapiInit.Signature,
		"timestamp": userapiInit.Timestamp,
	}

	// 设置请求头，和获取token
	headMap, base_url := GetHeaderUrl()
	resp, _, err := request.PostRequestCofig(userapiMap, base_url, api, headMap)
	if err != nil {
		fmt.Printf("获取用户的userid失败")
		return -1
	}
	usrid, err := returnUserId(string(resp))
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return usrid
}

// 解析userid
func returnUserId(jsonStr string) (int64, error) {
	// 定义结构体变量
	var response Useridstruct

	// 解析 JSON
	err := json.Unmarshal([]byte(jsonStr), &response)
	if err != nil {
		log.Fatalf("JSON 解析失败: %v", err)
	}

	// 提取 userId
	if len(response.Data.List) > 0 {
		userId := response.Data.List[0].UserId
		return userId, nil
	} else {
		err := errors.New("未找到 userId")
		return -1, err
	}
}
