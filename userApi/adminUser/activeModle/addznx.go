package activeModle

import (
	"encoding/json"
	"fmt"
	"log"
	"project/common"
	"project/request"
	"project/userApi/adminUser"
	"project/utils"
	"strconv"
	"sync"
)

// 发送站内信个数
const ZNX_NUMBER = 6

// Translation represents the translations array in the JSON
type Translation struct {
	Language  string `json:"language"`
	Content   string `json:"content"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
}

// FreeReward represents the freeReward object in rewardConfig
type FreeReward struct {
	RewardAmount         int    `json:"rewardAmount"`
	AmountCodingMultiple int    `json:"amountCodingMultiple"`
	CouponIds            string `json:"couponIds"`
}

// RechargeReward represents the rechargeReward object in rewardConfig
type RechargeReward struct {
	RechargeAmount       int    `json:"rechargeAmount"`
	RechargeCount        int    `json:"rechargeCount"`
	RewardAmount         int    `json:"rewardAmount"`
	AmountCodingMultiple int    `json:"amountCodingMultiple"`
	CouponIds            string `json:"couponIds"`
}

// RewardConfig represents the rewardConfig object in the JSON
type RewardConfig struct {
	FreeReward     FreeReward     `json:"freeReward"`
	RewardTypes    []int          `json:"rewardTypes"`
	RechargeReward RechargeReward `json:"rechargeReward"`
	ExpireType     int            `json:"expireType"`
}

// Message represents the entire JSON structure
type Message struct {
	BackstageDisplayName string        `json:"backstageDisplayName"`
	ValidType            int           `json:"validType"`
	Title                string        `json:"title"`
	JumpType             int           `json:"jumpType"`
	JumpPage             int           `json:"jumpPage"`
	JumpButtonText       string        `json:"jumpButtonText"`
	TargetType           int           `json:"targetType"`
	Translations         []Translation `json:"translations"`
	SendType             int           `json:"sendType"`
	IsHasReward          bool          `json:"isHasReward"`
	RewardConfig         RewardConfig  `json:"rewardConfig"`
	Random               int64         `json:"random"`
	Language             string        `json:"language"`
	Signature            string        `json:"signature"`
	Timestamp            int64         `json:"timestamp"`
}

/*
backstageDisplayName string,   站内信的名字
validType int,     默认值1
jumpType int,  跳转类型
jumpPage int,  跳转页面
jumpButtonText string  跳转的按钮文字
targetType int, 跳转目标 ，接收对象
content string,  站内信的内容
sendType int,  发送类型
*
*/
func CreateMessage(
	backstageDisplayName string,
	validType int,
	jumpType int,
	jumpPage int,
	jumpButtonText string,
	targetType int,
	content string,
	sendType int,
	random int64,
	timestamp int64,
	title string,
) map[string]interface{} {
	// Create the Message struct with provided parameters and default values from JSON
	message := Message{
		BackstageDisplayName: backstageDisplayName,
		ValidType:            validType,
		JumpType:             jumpType,
		JumpPage:             jumpPage,
		JumpButtonText:       jumpButtonText,
		TargetType:           targetType,
		SendType:             sendType,
		Random:               random,
		Timestamp:            timestamp,
		Language:             "en", // Default from JSON
		Signature:            "",   // Default from JSON
		IsHasReward:          true, // Default from JSON
		Translations: []Translation{
			{
				Language:  "hi",
				Content:   content, // Use provided content for hi
				Title:     title,
				Thumbnail: "3003/other/042436178-1139-IMG_202508293559_512x512.png",
			},
			{
				Language:  "en",
				Content:   content, // Use provided content for en
				Title:     title,
				Thumbnail: "3003/other/042440029-1140-IMG_202508294048_512x512.png",
			},
			{
				Language:  "zh",
				Content:   content, // Use provided content for zh
				Title:     title,
				Thumbnail: "3003/other/042443744-1141-IMG_202509011088_512x512.png",
			},
		},
		RewardConfig: RewardConfig{
			FreeReward: FreeReward{
				RewardAmount:         10,
				AmountCodingMultiple: 11,
				CouponIds:            "400007",
			},
			RewardTypes: []int{1, 2},
			RechargeReward: RechargeReward{
				RechargeAmount:       1000,
				RechargeCount:        11,
				RewardAmount:         100,
				AmountCodingMultiple: 11,
				CouponIds:            "400007",
			},
			ExpireType: 1,
		},
	}

	// Convert Message struct to map[string]interface{}
	result := map[string]interface{}{
		"backstageDisplayName": message.BackstageDisplayName,
		"validType":            message.ValidType,
		"jumpType":             message.JumpType,
		"jumpPage":             message.JumpPage,
		"jumpButtonText":       message.JumpButtonText,
		"targetType":           message.TargetType,
		"translations":         message.Translations,
		"sendType":             message.SendType,
		"isHasReward":          message.IsHasReward,
		"rewardConfig":         message.RewardConfig,
		"random":               message.Random,
		"language":             message.Language,
		"signature":            message.Signature,
		"timestamp":            message.Timestamp,
		"title":                message.Title,
	}

	return result
}

// 界面跳转的枚举值
type jumpTypeTotaIota int8

const (
	cz jumpTypeTotaIota = iota + 1
	tx
	lpm
	yhq
	cjdj
	xm
	vip
	jbs
	phb
	hdlb
	znx
	yqzp
	xbfy
	czzp
	me
	sy
	czhdlb
)

// 获取请求头的map
func GetHeaderMap() (map[string]interface{}, string) {
	// 请求头
	deskA := &common.AdminHeaderAuthorizationConfig2{}
	base_url := common.ADMIN_SYSTEM_url
	token := adminUser.GetToken()
	desSlice := []interface{}{base_url, base_url, base_url, token}
	headMap, err := common.AssignSliceToStructMap(deskA, desSlice)
	if err != nil {
		fmt.Println("headerMap获取失败", err)
		return nil, ""
	}
	return headMap, base_url
}

type QueryZnx struct {
	PageNo    any `json:"pageNo"`
	PgeSize   any `json:"pageSize"`
	OrderBy   any `json:"orderBy"`
	Random    any `json:"Random"`
	Language  any `json:"language"`
	Signature any `json:"signature"`
	Timestamp any `json:"timestamp"`
}

type Root struct {
	Data Data `json:"data"`
}

type Data struct {
	List []Notification `json:"list"`
}

type Notification struct {
	ID int64 `json:"id"`
}

// 返回一个id
func QueryZnxFunc() []int64 {
	randmo := request.RandmoNie()
	timestamp := request.GetNowTime()
	// 获取请求头
	headMap, base_url := GetHeaderMap()
	api := "/api/Inmail/GetPageList"
	query := &QueryZnx{}
	resultList := []interface{}{1, 20, "Desc", randmo, "en", "", timestamp}
	resultMap, err := common.AssignSliceToStructMap(query, resultList)
	if err != nil {
		fmt.Println("站内信查询的结构体初始化报错", err)
		return nil
	}

	resp, _, err := request.PostRequestCofig(resultMap, base_url, api, headMap)
	if err != nil {
		fmt.Println("站内信的请求失败", err)
	}
	// fmt.Println("站内信的查询响应内容", string(resp))
	// 解析 JSON
	var root Root
	err = json.Unmarshal([]byte(resp), &root)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// 提取所有 id
	var ids []int64
	for _, notification := range root.Data.List {
		ids = append(ids, notification.ID)
	}

	// 打印 id 列表
	return ids[:ZNX_NUMBER+10]
}

// 点击启用
type ClickZnxStruct struct {
	State     any `json:"state"`
	Id        any `json:"id"`
	Random    any `json:"random"`
	Language  any `json:"language"`
	Signature any `json:"signature"`
	Timestamp any `json:"timestamp"`
}

/*
传入要点击的id，才能启用
*
*/
func ClickZnxFunc(id int64) (string, error) {
	randmo := request.RandmoNie()
	timestamp := request.GetNowTime()
	api := "/api/Inmail/UpdateState"
	// 获取请求头和base_url
	headerMap, base_url := GetHeaderMap()
	// 初始化结构体
	clickZnx := &ClickZnxStruct{}
	resultList := []interface{}{1, id, randmo, "en", "", timestamp}
	resultMap, err := common.AssignSliceToStructMap(clickZnx, resultList)
	if err != nil {
		//fmt.Println("点击站内信的启用按钮请求结构体失败", err)
		return "", fmt.Errorf("点击站内信的启用按钮请求结构体失败%v\n", err)
	}

	resp, _, err := request.PostRequestCofig(resultMap, base_url, api, headerMap)
	if err != nil {
		// fmt.Println("点击站内信的启用按钮请求失败", err)
		return "", fmt.Errorf("点击站内信的启用按钮请求失败%v\n", err)
	}

	return string(resp), nil
}

// 需要提供跳转类型，和跳转文字
func SendZnx(jumpNumber int, jumpText string) {
	rand := request.RandmoNie()
	timestamp := request.GetNowTime()
	znxTitle := "自动化生成的站内信" + strconv.FormatInt(timestamp, 10)
	result := CreateMessage(znxTitle, 1, 1, jumpNumber, jumpText, 1, "这是内容", 1, rand, timestamp, znxTitle)
	headMap, base_url := GetHeaderMap()

	api := "/api/Inmail/Add"
	resp, _, err := request.PostRequestCofig(result, base_url, api, headMap)
	if err != nil {
		fmt.Println("站内信的请求失败", err)
	}
	fmt.Println("站内信响应内容", string(resp))
}

func SendOneZnx() {

	// 发送站内信
	// for i := 1; i <= 50; i++ {
	// 	SendZnx(i, "跳转"+strconv.Itoa(i))
	// 	// 获取id
	// 	id := QueryZnxFunc()
	// 	if id > 0 {
	// 		ClickZnxFunc(id)
	// 	}
	// }
	idlist := utils.RandmoUserId(ZNX_NUMBER)
	results := MoreSendZnx(idlist, 3)
	// 打印结果
	for _, result := range results {
		fmt.Println(result)
	}
}

func SendAllZnx() {
	// jumpTextList := []string{
	// 	"充值","提现","礼品码","优惠券","超级大奖","洗码","vip","锦标赛","排行榜","活动礼包","站内信","邀请转盘","新版返佣","充值转盘","我的","首页","充值活动礼包",
	// }
	// StructToMap()
}

/*
多并发的发送站内信
包含并发逻辑，保证 a -> b -> c 串行
inputs需要提供第一个函数的入参
concurrency 并发量
*
*/
func MoreSendZnx(inputs []string, concurrency int) []string {
	idChan := make(chan struct{}, 1)
	sem := make(chan struct{}, concurrency) // 信号量通道，控制并发量
	var wg sync.WaitGroup                   // 等待所有任务完成
	// results := make(chan struct {
	// 	index  int
	// 	result string
	// }, len(inputs)) // 存储结果，带索引以便排序

	// 处理每个输入
	for i, input := range inputs {
		wg.Add(1)
		go func(idx int, input string, dChan chan struct{}) {

			defer wg.Done()
			sem <- struct{}{}        // 占用一个并发槽
			defer func() { <-sem }() // 释放并发槽

			// 串行执行 a -> b -> c
			// 按顺序执行 a -> b -> c
			SendZnx(i, "跳转"+strconv.Itoa(i))
			// id := QueryZnxFunc()
			// if str, err := ClickZnxFunc(id); err != nil {
			// 	fmt.Println("站内信的启用失败", err)
			// 	return
			// } else {
			// 	// 将结果发送到通道
			// 	results <- struct {
			// 		index  int
			// 		result string
			// 	}{
			// 		index:  idx,
			// 		result: fmt.Sprintf("站内信发送结果 %v:\n", str),
			// 	}
			// }
			idChan <- struct{}{}

		}(i, input, idChan)
	}
	// 上面的步骤结束了，进行收集所有的站内信的id

	<-idChan
	znxIdList := QueryZnxFunc()
	for i := 0; i < len(znxIdList); i++ {
		if str, err := ClickZnxFunc(znxIdList[i]); err != nil {
			fmt.Println("站内信的启用失败", err)
			continue
		} else {
			fmt.Println("站内信启用成功", str)
		}
	}

	// 关闭结果通道
	go func() {
		wg.Wait()
		//close(results)
	}()

	// 收集结果并按索引排序
	finalResults := make([]string, len(inputs))
	// for result := range results {
	// 	finalResults[result.index] = result.result
	// }

	return finalResults
}
