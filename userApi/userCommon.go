package userApi

// 定义与 JSON 结构对应的结构体
type Response struct {
	Data          interface{} `json:"data"`          // 使用 interface{} 处理 null
	MsgParameters interface{} `json:"msgParameters"` // 使用 interface{} 处理 null
	Code          int         `json:"code"`
	Msg           string      `json:"msg"`
	MsgCode       int         `json:"msgCode"`
}
