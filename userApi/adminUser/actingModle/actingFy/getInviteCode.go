package actingFy

import (
	"encoding/json"
	"fmt"
	"project/common"
	"project/request"
	"project/userApi/adminUser"
)

// 根据用户的userName获取邀请码
type GetInviteCodeStruct struct {
	Account string `json:"account"`
}

// 定义与 JSON 结构对应的 Go 结构体
type GetInviteCodeResponse struct {
	Data struct {
		List []struct {
			InviteCode string `json:"inviteCode"`
			UserId     int64  `json:"userId"`
		} `json:"list"`
	} `json:"data"`
}

// 用户查询 返回邀请码
func GetInviteCodeFunc(userName string) (string, int64) {
	api := "/api/Users/GetPageList"
	randmo := request.RandmoNie()
	timestamp := request.GetNowTime()
	// 结构体初始化
	query := &QueryTifyStruct{}
	queryList := []interface{}{userName, 1, 20, "Desc", randmo, "en", "", timestamp}
	queryMap, _ := common.StructToMap(query, queryList)
	// 请求头
	token := adminUser.GetToken()
	base_url := common.ADMIN_SYSTEM_url
	headerStruct := &common.AdminHeaderAuthorizationConfig2{}
	headerList := []interface{}{base_url, base_url, base_url, token}
	headerMap, _ := common.AssignSliceToStructMap(headerStruct, headerList)
	// 把嵌套的的map进行平铺
	fietMap := common.FlattenMap(queryMap)
	resp, _, err := request.PostRequestCofig(fietMap, base_url, api, headerMap)
	if err != nil {
		fmt.Println(err)
		return "", -1
	}
	// 解析 JSON
	var response GetInviteCodeResponse
	err = json.Unmarshal([]byte(string(resp)), &response)
	if err != nil {
		fmt.Printf("JSON 解析失败: %v", err)
	}

	// 提取 number 的值
	if len(response.Data.List) > 0 {
		inviteCode := response.Data.List[0].InviteCode
		userId := response.Data.List[0].UserId
		fmt.Println("inviteCode 的值:", inviteCode, userId)
		return inviteCode, userId
	} else {
		fmt.Println("List 为空，无法提取 number")
	}
	return "", -1
}
