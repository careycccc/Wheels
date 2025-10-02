package actingFy

import (
	"encoding/json"
	"fmt"
	"project/common"
	"project/request"
	"project/userApi/adminUser"

	"github.com/bytedance/sonic"
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

// 用户查询 返回邀请码 和用户的id
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

// 查询所有的用户的数据
type GetAllUserInfoStruct struct {
	UserType int `json:"userType"`
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
	common.BaseStruct
}

// 每个用户在会员列表的的信息
type UserInfo struct {
	UserId           int64   `json:"userId"`
	Account          string  `json:"account"`
	NickName         string  `json:"nickName"`
	VipLevel         int     `json:"vipLevel"`
	Balance          float64 `json:"balance"`
	ParentId         int     `json:"parentId"`
	GeneralAgentId   int64   `json:"generalAgentId"`
	State            int     `json:"state"`
	IsFuncation      bool    `json:"isFuncation"`
	RegisterTime     int64   `json:"registerTime"`
	RegisterSource   int     `json:"registerSource"`
	LastLoginTime    int64   `json:"lastLoginTime"`
	Remark           string  `json:"remark"`
	UserType         int     `json:"userType"`
	InviteCode       string  `json:"inviteCode"`
	ChannelId        int     `json:"channelId"`
	RegisterIp       string  `json:"registerIp"`
	LastLoginIp      string  `json:"lastLoginIp"`
	IsBlackListIPTag bool    `json:"isBlackListIPTag"`
	PackageId        int     `json:"packageId"`
	PackageName      string  `json:"packageName"`
}

// Data结构体用于解析响应中的data字段
type Data struct {
	List       []UserInfo `json:"list"`
	PageNo     int        `json:"pageNo"`
	TotalPage  int        `json:"totalPage"`
	TotalCount int        `json:"totalCount"`
}

/*
查询所有的用户的数据
userType 0:正式账号 1:测试账号 2:游客
pageNumber 每页显示多少条数据
返回当前类型的用户列表和用户总数
*/
func GetAllUserInfoFunc(userType, pageNumber int) ([]UserInfo, int, error) {
	api := "/api/Users/GetPageList"
	randmo := request.RandmoNie()
	timestamp := request.GetNowTime()
	token := adminUser.GetToken()
	payloadStruct := &GetAllUserInfoStruct{}
	payloadList := []interface{}{userType, 1, pageNumber, randmo, "en", "", timestamp}
	headerStruct := &common.AdminHeaderAuthorizationConfig2{}
	headerList := []interface{}{common.ADMIN_SYSTEM_url, common.ADMIN_SYSTEM_url, common.ADMIN_SYSTEM_url, token}
	// 调用泛型函数
	respBody := request.PostGenericsFuncFlatten[GetAllUserInfoStruct, common.AdminHeaderAuthorizationConfig2](common.ADMIN_SYSTEM_url, api, payloadStruct, payloadList, headerStruct, headerList, common.StructToMap, common.AssignSliceToStructMap)

	userList, totalCount, err := ParseListFromResponse(respBody)
	if err != nil {
		return nil, -1, err
	}
	return userList, totalCount, nil

}

// 获取当前类型的用户的总的数量
// userType 0:正式账号 1:测试账号 2:游客
// 返回当前类型的用户列表和用户总数
func GetUserTypeTotal(userType int, pageNumber int) ([]UserInfo, int, error) {
	userList, totalCount, _ := GetAllUserInfoFunc(userType, pageNumber)
	fmt.Println("用户的总数量", len(userList), totalCount)
	// 获取的所有totalCount进行再次用户量列表
	userList, totalCount, _ = GetAllUserInfoFunc(userType, pageNumber)
	GetUserIdByUserList(userList)
	return userList, totalCount, nil
}

// 返回用户列表和总数
func ParseListFromResponse(postResp *request.PostReponsestruct) ([]UserInfo, int, error) {
	// 将 Data 转换为字节切片
	dataBytes, err := json.Marshal(postResp.Data)
	if err != nil {
		return nil, -1, fmt.Errorf("error marshaling Data: %v", err)
	}

	// 解析 Data 中的 list
	var data Data
	err = sonic.UnmarshalString(string([]byte(dataBytes)), &data)
	if err != nil {
		return nil, 0, fmt.Errorf("JSON 解析失败: %w", err)
	}
	return data.List, data.TotalCount, nil
}

// 根据用户的列表进行批量获取用户的userid
func GetUserIdByUserList(userList []UserInfo) []int64 {
	userIds := make([]int64, 0, len(userList))
	for _, user := range userList {
		userIds = append(userIds, user.UserId)
	}
	fmt.Println("用户的ID列表+++++", userIds)
	return userIds
}
