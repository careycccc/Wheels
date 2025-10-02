package deskApi

import (
	"encoding/json"
	"fmt"
	"project/common"
	"project/request"
	"project/utils"
	"reflect"
	"strings"
)

// // 注册接口

// type RegisterStruct struct {
// 	UserName            string `json:"userName"`
// 	VerifyCode          string `json:"verifyCode"`
// 	InviteCode          string `json:"inviteCode"`
// 	RegisterFingerprint string `json:"registerFingerprint"`
// 	Random              int64  `json:"randmo"`
// 	Language            string `json:"language"`
// 	Signature           string `json:"signature"`
// 	Timestamp           int64  `json:"timestamp"`
// 	TrackStruct
// }

// type TrackStruct struct {
// 	IsTrusted bool  `json:"isTrusted"`
// 	Vts       int64 `json:"_vts"`
// }

type ResponseResiter struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

// 嵌套结构体 Track
type Track struct {
	IsTrusted bool  `json:"isTrusted"`
	Vts       int64 `json:"_vts"`
}

// 主结构体
type RegisterStruct struct {
	UserName            string `json:"userName"`
	VerifyCode          string `json:"verifyCode"`
	InviteCode          string `json:"inviteCode"`
	RegisterFingerprint string `json:"registerFingerprint"`
	Track               Track  `json:"track"`
	Language            string `json:"language"`
	Random              int64  `json:"random"`
	Signature           string `json:"signature"`
	Timestamp           int64  `json:"timestamp"`
}

/*
注册接口
userName  用户名
verifyCode 验证码
inviteCode 邀请码
return token  返回token
*
*/
func RegisterFunc(userName, verifyCode, inviteCode string) string {
	fmt.Println("+++++++++", userName, verifyCode, inviteCode)
	api := "/api/Home/MobileAutoLogin"
	base_url := common.SIT_WEB_API
	random := request.RandmoNie()
	timestamp := request.GetNowTime()
	generate := utils.GenerateCryptoRandomString(32)
	payloadStruct := &RegisterStruct{}
	RegisterList := []interface{}{userName, verifyCode, inviteCode, generate, Track{IsTrusted: true, Vts: timestamp}, "en", random, "", timestamp}
	registerMap, err := common.InitStructToMap(payloadStruct, RegisterList)
	if err != nil {
		fmt.Println("注册信息的payload的map报错", err)
		return ""
	}
	// filedMap := common.FlattenMap(registerMap)
	register_url := common.REGISTER_url
	registreList := []interface{}{"3003", register_url, register_url, register_url}
	// 初始化请求头
	headerconfig := &common.DeskHeaderConfig2{}
	headerMap, err := common.InitStructToMap(headerconfig, registreList)
	if err != nil {
		fmt.Println("注册信息的payload的map报错", err)
		return ""
	}
	//fmt.Println("+++++++registerMap", registerMap)
	respBoy, _, err := request.PostRequestCofig(registerMap, base_url, api, headerMap)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 返回token
	//fmt.Println("注册返回的结果", string(respBoy))
	var response ResponseResiter
	err = json.Unmarshal(respBoy, &response)
	if err != nil {
		fmt.Printf("解析响应失败: %v\n", err)
		return ""
	}
	return response.Data.Token

}

func InitStructToMap(strct interface{}, values []interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// 检查输入是否为结构体指针
	v := reflect.ValueOf(strct)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("first parameter must be a pointer to a struct")
	}
	v = v.Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// 跳过不可导出字段
		if fieldType.PkgPath != "" {
			continue
		}

		// 获取 JSON 标签中的字段名
		tag := fieldType.Tag.Get("json")
		if tag == "" {
			tag = fieldType.Name
		} else {
			if commaIdx := strings.Index(tag, ","); commaIdx != -1 {
				tag = tag[:commaIdx]
			}
		}

		// 如果切片长度不足，使用字段的当前值
		if i >= len(values) || !reflect.ValueOf(values[i]).IsValid() {
			result[tag] = v.Field(i).Interface()
			continue
		}

		// 处理嵌套结构体
		if fieldType.Type.Kind() == reflect.Struct {
			nestedMap, err := InitStructToMap(reflect.New(fieldType.Type).Interface(), values[i:])
			if err != nil {
				return nil, fmt.Errorf("error in nested struct %s: %v", fieldType.Name, err)
			}
			result[tag] = nestedMap
			if field.CanSet() && i < len(values) {
				val := reflect.ValueOf(values[i])
				if val.Type().AssignableTo(fieldType.Type) {
					field.Set(val)
				}
			}
			continue
		}

		// 处理字段赋值
		if field.CanSet() {
			val := reflect.ValueOf(values[i])
			if val.Type().ConvertibleTo(field.Type()) {
				field.Set(val.Convert(field.Type()))
			} else {
				return nil, fmt.Errorf("cannot assign slice element type %v to field %s of type %v", val.Type(), fieldType.Name, field.Type())
			}
		}

		// 将字段值添加到结果 map
		result[tag] = v.Field(i).Interface()
	}

	return result, nil
}
