package activeModle

import (
	"fmt"
	"project/common"
	"project/request"
	"project/userApi/adminUser"
)

// 邀请转盘的自动分析结果

type WheelAutoResult struct {
	OrderNo  string `json:"orderNo"`  // 订单号
	PageNo   int    `json:"pageNo"`   // 当前页码
	PageSize int    `json:"pageSize"` // 每页数量
	common.BaseStruct
}

/*
*
传入订单号
*/
func WheelAutoResultFunc(orderNo string) {
	base_url := common.SIT_WEB_API
	api := "/api/InvitedWheel/GetPageListWithdrawRecordDetail"
	headerStruct := &common.DeskHeaderAstruct{}
	headerUrl := common.ADMIN_SYSTEM_url
	token := adminUser.GetToken()
	headerList := []interface{}{headerUrl, headerUrl, headerUrl, token}
	// 组装请求体
	payloadStruct := &WheelAutoResult{}
	random := request.RandmoNie()
	timeStamp := request.GetNowTime()
	payloadList := []interface{}{orderNo, 1, 20, random, "en", "", timeStamp}
	resp := request.PostGenericsFuncFlatten[WheelAutoResult, common.DeskHeaderAstruct](base_url, api, payloadStruct, payloadList, headerStruct, headerList, common.StructToMap, common.AssignSliceToStructMap)
	fmt.Println("邀请转盘的自动分析结果", resp)
}
