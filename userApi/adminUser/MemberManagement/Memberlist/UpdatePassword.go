package memberlist

import (
	"fmt"
	"project/common"
	"project/request"
	"project/userApi/adminUser"
	"sync"
)

// 后台的修改密码
type UpdataPasswordstruct struct {
	UserId   int64  `json:"userId"`
	Password string `json:"password"`
	common.BaseStruct
}

/*
后台的修改密码
userid 用户的id
password 要修改的密码
*
*/
func UpdataPasswordFunc(userid int64, password string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	api := "/api/Users/UpdatePassword"
	// 获取token
	token := adminUser.GetToken()
	headerStruct := &common.DeskHeaderAstruct{}
	header_url := common.ADMIN_SYSTEM_url
	header_list := []interface{}{header_url, header_url, header_url, token}
	headerMap, err := common.AssignSliceToStructMap(headerStruct, header_list)
	if err != nil {
		fmt.Println(err)
		return
	}
	randmo := request.RandmoNie()
	timestamp := request.GetNowTime()
	payloadstruct := &UpdataPasswordstruct{}
	payloadList := []interface{}{userid, password, randmo, "en", "", timestamp}
	payloadMap, err := common.StructToMap(payloadstruct, payloadList)
	if err != nil {
		fmt.Println(err)
		return
	}
	flatMap := common.FlattenMap(payloadMap)

	_, _, err = request.PostRequestCofig(flatMap, header_url, api, headerMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(respBoy))
	results <- "1"
}
