package utils

import (
	"bufio"
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	_ "math/rand"
	"os"
	"reflect"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// Md5Info 计算 MD5 哈希值
func Md5Info(data string, uppercase bool) string {
	hash := md5.New()
	hash.Write([]byte(data))
	result := hex.EncodeToString(hash.Sum(nil))
	if uppercase {
		return strings.ToUpper(result)
	}
	return strings.ToLower(result)
}

// GetSignature 生成签名
func GetSignature(body map[string]interface{}, verifyPwd *string) string {
	// 过滤字段
	filteredObj := make(map[string]interface{})
	keys := make([]string, 0, len(body))
	for key := range body {
		keys = append(keys, key)
	}
	sort.Strings(keys) // 按键排序

	for _, key := range keys {
		value := body[key]
		// 检查 value 不为 nil 且不为空字符串，且 key 不在排除列表中，且 value 不是数组
		if value != nil && value != "" && key != "signature" && key != "timestamp" && key != "track" {
			// 确保 value 不是切片（相当于 Python 的 list）
			if _, ok := value.([]interface{}); !ok {
				filteredObj[key] = value
			}
		}
	}

	// 转换为 JSON 字符串
	jsonData, err := json.Marshal(filteredObj)
	if err != nil {
		return "" // 错误处理，可根据需求调整
	}

	encoder := string(jsonData)
	if verifyPwd != nil {
		encoder += *verifyPwd
	}

	// 计算 MD5
	return Md5Info(encoder, true)
}

func GetSignature2(body any, verifyPwd *string) string {
	// 过滤字段并转换为 map 以便排序
	filteredObj := make(map[string]interface{})
	excludeKeys := map[string]bool{"signature": true, "timestamp": true, "track": true}

	// 使用反射获取结构体字段
	val := reflect.ValueOf(body)
	typ := reflect.TypeOf(body)
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i).Interface()
		jsonTag := field.Tag.Get("json")
		// 提取 JSON 字段名（忽略 ",omitempty" 部分）
		key := jsonTag
		if idx := len(jsonTag); idx > 9 && jsonTag[idx-9:] == ",omitempty" {
			key = jsonTag[:idx-9]
		}

		// 过滤条件：非空值、不在排除列表中
		if !excludeKeys[key] && !isEmpty(value) {
			filteredObj[key] = value
		}
	}

	// 按键排序
	keys := make([]string, 0, len(filteredObj))
	for key := range filteredObj {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 转换为 JSON 字符串
	jsonData, err := json.Marshal(filteredObj)
	if err != nil {
		return "" // 错误处理
	}

	encoder := string(jsonData)
	if verifyPwd != nil {
		encoder += *verifyPwd
	}

	// 计算 MD5
	return Md5Info(encoder, true)
}

// isEmpty 检查值是否为空（nil、零值或空字符串）
func isEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Slice, reflect.Array, reflect.Map:
		return v.Len() == 0
	}
	return false
}

// 用来解析请求
func Unmarshal(strResbody string) (result map[string]interface{}) {
	//是一个字符串
	error := json.Unmarshal([]byte(strResbody), &result)
	if error != nil {
		log.Fatalf("解析响应失败~~:%v", error)
	}
	return

}

// 读取yaml
// ReadYAML 从指定路径读取 YAML 文件并解析到结构体
func ReadYAML(filePath string, result interface{}) error {
	// 读取文件内容
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("读取 YAML 文件失败: %w", err)
	}

	// 解析 YAML 数据到结构体
	if err := yaml.Unmarshal(data, result); err != nil {
		return fmt.Errorf("解析 YAML 数据失败: %w", err)
	}

	return err
}

// 写入yaml 以覆盖的方式
// WriteYAML 以覆盖方式将结构体写入 YAML 文件
func WriteYAML(filePath string, data interface{}) error {
	// 将结构体编码为 YAML 数据
	yamlData, err := yaml.Marshal(data)
	// fmt.Printf("文件路径%s,写入的文件内容%v", filePath, yamlData)
	if err != nil {
		return fmt.Errorf("编码 YAML 数据失败: %w", err)
	}

	// 以覆盖方式写入文件（os.Create 会创建或覆盖文件）
	if err := os.WriteFile(filePath, yamlData, 0644); err != nil {
		return fmt.Errorf("写入 YAML 文件失败: %w", err)
	}

	return nil
}

// 处理两层map[string]interface{}
func HandlerMap(strResbody string, str string) (string, error) {
	var result map[string]interface{}
	result = Unmarshal(strResbody)
	// fmt.Printf("解析结果%v", result)
	innerMap, ok := result["data"].(map[string]interface{})
	if !ok {
		// fmt.Println("data 不存在")
		err := errors.New("data 不存在")
		return "第一层data不存在", err
	}

	// 访问内层 map 的 token
	token, ok := innerMap[str].(string)
	if !ok {
		// fmt.Println("token 不是字符串或不存在")
		err := errors.New("token 不是字符串或不存在")
		return "token 不是字符串或不存在", err
	}
	return token, nil
}

// 生成随机浏览器指纹 ，设备号
func GenerateCryptoRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}
	// fmt.Println("本次指纹", string(bytes))
	return string(bytes)
}

// 随机生成用户
func RandmoUserCount() (string, error) {
	// 获取当前日期
	now := time.Now()
	month := now.Month()
	day := now.Day()

	// 格式化月和日
	var prefix string
	if month < 10 {
		prefix = fmt.Sprintf("%d%02d", month, day) // 月1位+日2位=3位
	} else {
		prefix = fmt.Sprintf("%02d%02d", month, day) // 月2位+日2位=4位
	}

	// 根据前缀长度决定随机数位数
	var randomLength int
	if len(prefix) == 3 {
		randomLength = 7
	} else {
		randomLength = 6
	}

	// 生成随机数
	max := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(randomLength)), nil) // 10^randomLength
	randNum, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}

	// 格式化随机数，补0到指定长度
	randStr := fmt.Sprintf("%0*d", randomLength, randNum)

	// 合并前缀和随机数
	return "91" + prefix + randStr, nil
}

// 随机生成一个数字
func RandmoNumber(number int) int64 {
	// 初始化随机数种子
	var b [8]byte // 一个 int64 需要 8 个字节
	_, err := rand.Read(b[:])
	if err != nil {
		panic(err) // 在实际应用中，应该优雅地处理错误
	}
	return int64(binary.BigEndian.Uint64(b[:]))
}

// 随机生成 min max的数据
func GenerateRandomInt(min, max int64) (int64, error) {
	if min > max {
		return 0, fmt.Errorf("min must be less than or equal to max")
	}

	// 生成一个大于等于0且小于max-min的随机数
	randomInt, err := rand.Int(rand.Reader, big.NewInt(max-min))
	if err != nil {
		return 0, err
	}

	// 将随机数加上最小值，得到最终的随机数范围在[min, max]之间
	randomInt.Add(randomInt, big.NewInt(min))
	return randomInt.Int64(), nil
}

// AppendToFile 将内容追加写入指定路径的文件
func AppendToYAML(filePath string, data ...interface{}) error {
	// 以追加模式打开文件，如果文件不存在则创建
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将可变参数转换为字符串并写入
	for _, item := range data {
		// 使用 fmt.Sprintf 将任意类型转换为字符串
		content := fmt.Sprintf("%v\n", item)
		if _, err := file.WriteString(content); err != nil {
			return err
		}
	}

	return nil
}

// ReadYAMLByLine 逐行读取YAML文件，返回字符串列表
func ReadYAMLByLine(filePath string) ([]string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 创建字符串切片存储每一行
	var lines []string
	scanner := bufio.NewScanner(file)

	// 逐行读取
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// 检查扫描过程中是否出现错误
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
