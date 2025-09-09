package common

//基础结构体
type BaseStruct struct {
	Random    int64  `json:"randmo"`
	Language  string `json:"language"`
	Signature string `json:"signature"`
	Timestamp int64  `json:"timestamp"`
}

// 基础结构体带签名
type BaseSignatureStruct struct {
	Signature any `json:"signature"`
	BaseStruct
}

//基础结构体带 排序
type BaseOrderByStruct struct {
	PageNo   int8   `json:"pageNo"`
	PageSize int64  `json:"pageSize"`
	OrderBy  string `json:"orderBy"`
	BaseStruct
}
