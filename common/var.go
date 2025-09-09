package common

import (
	"fmt"
	"reflect"
)

// 商户后台的账号和密码
type AdminUserName struct {
	UserName string // 账号
	Pwd      string // 密码
}

func (admin *AdminUserName) AdminUserInit() *AdminUserName {
	admin.UserName = "carey3003"
	admin.Pwd = "qwer1234"
	return admin
}

var cofingURL CofingURL

// url配置
type CofingURL struct {
	ADMIN_URL string // 后台地址
	H5_URL    string // 前台地址
	BET_URL   string // 投注地址
	Iss_URL   string // 获取期号的地址
}

func (config *CofingURL) ConfigUrlInit() *CofingURL {
	config.ADMIN_URL = "https://sit-tenantadmin-3003.mggametransit.com"
	config.H5_URL = "https://sit-webapi.mggametransit.com"
	config.BET_URL = "https://sit-lotteryh5.wmgametransit.com"
	config.Iss_URL = "https://h5.wmgametransit.com"
	return config
}

// 后台的请求头设置不带token
type AdminHeaderConfig struct {
	Domainurl string
	Referer   string
	Origin    string
}

// 结构体的自我初始化
func NewAdminHeaderConfig() *AdminHeaderConfig {
	return &AdminHeaderConfig{
		Domainurl: "Domainurl",
		Referer:   "Referer",
		Origin:    "Origin",
	}
}

func (header *AdminHeaderConfig) AdminHeaderConfigFunc() map[string]interface{} {
	header = NewAdminHeaderConfig()

	url := cofingURL.ConfigUrlInit().ADMIN_URL
	return map[string]interface{}{
		header.Domainurl: url,
		header.Referer:   url,
		header.Origin:    url,
	}
}

// 后台请求设置带toen
type AdminHeaderAuthorizationConfig2 struct {
	Domainurl     any
	Referer       any
	Origin        any
	Authorization any
}

// 后台请求头设置带token
type AdminHeaderAuthorizationConfig struct {
	Authorization string // token
	AdminHeaderConfig
}

func newAdminHeaderAuthorizationConfig() *AdminHeaderAuthorizationConfig {
	header := NewAdminHeaderConfig()
	return &AdminHeaderAuthorizationConfig{
		Authorization:     "Authorization",
		AdminHeaderConfig: *header,
	}
}

// token 把登录后的token
func (author *AdminHeaderAuthorizationConfig) AdminHeaderAuthorizationFunc(token string) map[string]interface{} {
	author = newAdminHeaderAuthorizationConfig()
	url := cofingURL.ADMIN_URL
	return map[string]interface{}{
		author.Authorization: "Bearer " + token,
		author.Domainurl:     url,
		author.Referer:       url,
		author.Origin:        url,
	}
}

// 前台的请求头设置
type DeskHeaderConfig struct {
	tenantId string
	AdminHeaderConfig
}

type DeskHeaderConfig2 struct {
	TenantId  any
	Domainurl any
	Referer   any
	Origin    any
}

func NewDeskHeaderConfig() *DeskHeaderConfig {
	header := NewAdminHeaderConfig()
	return &DeskHeaderConfig{
		tenantId:          "tenantId",
		AdminHeaderConfig: *header,
	}
}

// 前台登录的
func (desk *DeskHeaderConfig) DeskHeaderConfigFunc() map[string]interface{} {
	return map[string]interface{}{
		desk.tenantId:  "3003",
		desk.Domainurl: "https://sit-plath5-y1.mggametransit.com",
		desk.Referer:   "https://sit-plath5-y1.mggametransit.com",
		desk.Origin:    "https://sit-plath5-y1.mggametransit.com",
	}
}

// 下注的请求头设置
type BetHeaderConfig struct {
	Referer       string
	Origin        string
	Authorization string
	Sec           string
	SecCh         string
	SecUa         string
	SecFetch      string
	SecFetchMode  string
	SecFetchDest  string
}

func NewBetHeaderConfig() *BetHeaderConfig {
	return &BetHeaderConfig{
		Referer:       "Referer",
		Origin:        "Origin",
		Authorization: "Authorization",
		Sec:           "sec-ch-ua-platform",
		SecCh:         "sec-ch-ua",
		SecUa:         "sec-ch-ua-mobile",
		SecFetch:      "Sec-Fetch-Site",
		SecFetchMode:  "Sec-Fetch-Mode",
		SecFetchDest:  "Sec-Fetch-Dest",
	}
}

func (bet *BetHeaderConfig) BetHeaderConfigFunc(token string) map[string]interface{} {
	bet = NewBetHeaderConfig()
	return map[string]interface{}{
		bet.Origin:        "https://h5.wmgametransit.com",
		bet.Referer:       "https://h5.wmgametransit.com",
		bet.Authorization: "Bearer " + token,
		bet.Sec:           "Windows",
		bet.SecCh:         `"Not;A=Brand";v="99", "Google Chrome";v="139", "Chromium";v="139"`,
		bet.SecUa:         "?0",
		bet.SecFetch:      "same-site",
		bet.SecFetchMode:  "cors",
		bet.SecFetchDest:  "empty",
	}
}

// 获取期号的请求头设置
type GetIssNunmberHeaderConfig struct {
	Referer string
}

func newGetIssNunmberHeaderConfig() *GetIssNunmberHeaderConfig {
	return &GetIssNunmberHeaderConfig{
		Referer: "Referer",
	}
}

// 前台获取下注的token
type BetTokenStruct struct {
	Referer       interface{}
	Origin        interface{}
	Authorization interface{}
}

// 前台登录后请求需要token的
type DeskHeaderAstruct struct {
	Referer       interface{}
	Origin        interface{}
	Domainurl     interface{}
	Authorization interface{}
}

// url地址
const (
	PLANT_H5         = "https://sit-plath5-y1.mggametransit.com"
	WMG_H5           = "https://h5.wmgametransit.com"
	LOTTERY_H5       = "https://sit-lotteryh5.wmgametransit.com"
	ADMIN_SYSTEM_url = "https://sit-tenantadmin-3003.mggametransit.com"
	REGISTER_url     = "https://sit-3003-register.mggametransit.com" // 注册地址
	SIT_WEB_API      = "https://sit-webapi.mggametransit.com"
)

// AssignSliceToStructMap 将切片的值一一对应赋值到结构体字段并返回 map[string]interface{}
// structObj结构体对象，sliceObj 切片对象
// 含有 Authorization
func AssignSliceToStructMap(structObj interface{}, sliceObj interface{}) (map[string]interface{}, error) {
	// 初始化结果 map
	result := make(map[string]interface{})

	// 检查结构体是否为指针
	structVal := reflect.ValueOf(structObj)
	if structVal.Kind() != reflect.Ptr || structVal.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("first parameter must be a pointer to a struct")
	}
	structVal = structVal.Elem()
	structType := structVal.Type()

	// 检查切片是否有效
	sliceVal := reflect.ValueOf(sliceObj)
	if sliceVal.Kind() != reflect.Slice {
		return nil, fmt.Errorf("second parameter must be a slice")
	}

	// 检查切片长度是否与结构体字段数量匹配
	numFields := structVal.NumField()
	if sliceVal.Len() < numFields {
		return nil, fmt.Errorf("slice length (%d) is less than struct field count (%d)", sliceVal.Len(), numFields)
	}

	// 将切片的值赋值给结构体字段
	for i := 0; i < numFields; i++ {
		field := structVal.Field(i)
		fieldName := structType.Field(i).Name
		sliceElement := sliceVal.Index(i)

		// 检查字段是否可设置
		if !field.CanSet() {
			return nil, fmt.Errorf("cannot set field %s", fieldName)
		}

		// 处理 Authorization 字段
		if fieldName == "Authorization" {
			// 尝试将切片元素转换为字符串
			var bearerValue string
			if sliceElement.Kind() == reflect.String {
				bearerValue = "Bearer " + sliceElement.String()
			} else {
				// 尝试将元素转换为字符串（支持常见类型）
				if sliceElement.CanInterface() {
					bearerValue = fmt.Sprintf("Bearer %v", sliceElement.Interface())
				} else {
					return nil, fmt.Errorf("slice element for Authorization must be convertible to string, got %v", sliceElement.Type())
				}
			}

			// 赋值给字段（任意类型支持）
			if field.Type().Kind() == reflect.Interface || field.Type() == reflect.TypeOf("") {
				field.Set(reflect.ValueOf(bearerValue))
				result[fieldName] = bearerValue
			} else {
				return nil, fmt.Errorf("Authorization field must be string or interface{} type, got %v", field.Type())
			}
		} else {
			// 其他字段的赋值
			if field.Type().Kind() == reflect.Interface || sliceElement.Type().AssignableTo(field.Type()) {
				field.Set(sliceElement)
				result[fieldName] = sliceElement.Interface()
			} else {
				return nil, fmt.Errorf("cannot assign slice element type %v to field %s of type %v",
					sliceElement.Type(), fieldName, field.Type())
			}
		}
	}

	return result, nil
}

/*
token
betType 投注的方式 wingo 30s  wingo1min  wingo 3min  wing 5min
*
*/
func (iss *GetIssNunmberHeaderConfig) GetIssNunmberHeaderFunc(token, betType string) map[string]interface{} {
	result := "https://h5.wmgametransit.com/WinGo/"
	iss = newGetIssNunmberHeaderConfig()
	if token == "" {
		//游客的方式
		result = result + betType
	} else {
		// token有值的情况
		r1 := "?Lang=en&Skin=Classic&SkinColor=Default&Token="
		r2 := "&RedirectUrl=https%3A%2F%2Fsit-plath5-y1.mggametransit.com%2Fgame%2Fcategory%3FcategoryCode%3DC202505280608510046&Beck=0"
		result = result + betType + r1 + token + r2
	}
	return map[string]interface{}{
		iss.Referer: result,
	}
}

// 初始化结构体，并且返回map
func InitStructToMap(strct interface{}, values []interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	v := reflect.ValueOf(strct).Elem() // 获取结构体值
	t := v.Type()

	for i := 0; i < v.NumField() && i < len(values); i++ {
		field := v.Field(i)

		// 处理字段可设置情况
		if field.CanSet() {
			val := reflect.ValueOf(values[i])

			// 类型不一致时尝试转换
			if val.Type().ConvertibleTo(field.Type()) {
				field.Set(val.Convert(field.Type()))
			}
		}

		// 优先用 JSON tag 作为 map key，否则用字段名
		tag := t.Field(i).Tag.Get("json")
		if tag == "" {
			tag = t.Field(i).Name
		}
		result[tag] = v.Field(i).Interface()
	}

	return result
}

// StructToMap 将结构体初始化并将切片值映射到 map   // 可以解决嵌套结构体
func StructToMap(structType interface{}, slice []interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// 获取结构体类型
	val := reflect.ValueOf(structType)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("structType must be a struct or struct pointer")
	}

	t := val.Type()
	sliceIndex := 0

	fmt.Printf("Processing struct with %d fields, slice length: %d\n", t.NumField(), len(slice))

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Name
		fieldType := field.Type

		// 检查字段是否可导出（PkgPath 非空表示不可导出）
		if field.PkgPath != "" {
			fmt.Printf("Skipping unexported field %s\n", fieldName)
			continue
		}

		// 处理嵌套结构体
		if fieldType.Kind() == reflect.Struct {
			if sliceIndex >= len(slice) {
				fmt.Printf("Slice exhausted at nested field %s, assigning nil\n", fieldName)
				result[fieldName] = nil
				continue
			}
			nestedMap, err := StructToMap(reflect.New(fieldType).Interface(), slice[sliceIndex:])
			if err != nil {
				return nil, fmt.Errorf("error in nested struct %s: %v", fieldName, err)
			}
			result[fieldName] = nestedMap
			continue
		}

		// 处理基本类型字段
		if sliceIndex >= len(slice) {
			fmt.Printf("Slice exhausted at field %s, assigning nil\n", fieldName)
			result[fieldName] = nil
			continue
		}

		sliceVal := reflect.ValueOf(slice[sliceIndex])
		if !sliceVal.IsValid() {
			fmt.Printf("Nil slice value at index %d for field %s\n", sliceIndex, fieldName)
			result[fieldName] = nil
			sliceIndex++
			continue
		}

		// 检查类型兼容性
		if sliceVal.Type().ConvertibleTo(fieldType) {
			result[fieldName] = sliceVal.Convert(fieldType).Interface()
			fmt.Printf("Assigned %v to field %s\n", sliceVal.Interface(), fieldName)
		} else {
			return nil, fmt.Errorf("cannot assign %v to field %s of type %v", sliceVal.Type(), fieldName, fieldType)
		}
		sliceIndex++
	}

	return result, nil
}

// FlattenMap 将嵌套的 map[string]interface{} 平铺为一层 map，忽略嵌套路径
func FlattenMap(nestedMap map[string]interface{}) map[string]interface{} {
	flatMap := make(map[string]interface{})

	for key, value := range nestedMap {
		// 如果值是嵌套 map，递归平铺
		if nested, ok := value.(map[string]interface{}); ok {
			// 将嵌套 map 的键值对直接合并到 flatMap
			for nestedKey, nestedValue := range FlattenMap(nested) {
				flatMap[nestedKey] = nestedValue // 后覆盖前
				fmt.Printf("Flattened key %s with value %v\n", nestedKey, nestedValue)
			}
		} else {
			// 直接赋值非 map 值
			flatMap[key] = value
			fmt.Printf("Flattened key %s with value %v\n", key, value)
		}
	}

	return flatMap
}
