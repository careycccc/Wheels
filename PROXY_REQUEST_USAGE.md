# PostRequestCofigProxy 使用说明

## 功能概述

`PostRequestCofigProxy` 是一个带代理功能的POST请求函数，它会自动从YAML文件中读取可用的代理IP，如果当前代理失效会自动尝试下一个可用的代理。

## 函数签名

```go
func PostRequestCofigProxy(payload map[string]interface{}, base_url, api string, args ...map[string]interface{}) ([]byte, *http.Response, error)
```

## 参数说明

- `payload`: 请求参数，类型为 `map[string]interface{}`
- `base_url`: 基础URL地址
- `api`: API接口路径
- `args`: 可选的请求头设置，类型为 `map[string]interface{}`

## 返回值

- `[]byte`: 响应体内容
- `*http.Response`: HTTP响应对象
- `error`: 错误信息

## 使用示例

### 基本使用

```go
package main

import (
    "fmt"
    "project/request"
    "project/common"
)

func main() {
    // 准备请求参数
    payload := map[string]interface{}{
        "userName":    "testuser",
        "password":    "testpass",
        "loginType":   "Mobile",
        "deviceId":    "",
        "browserId":   "testbrowser123",
        "packageName": "",
        "language":    "en",
        "random":      123456789,
        "signature":   "",
        "timestamp":   1640995200,
    }
    
    // 设置请求头
    headers := map[string]interface{}{
        "Origin":    "https://sit-plath5-y1.mggametransit.com",
        "Referer":   "https://sit-plath5-y1.mggametransit.com",
        "Domainurl": "https://sit-plath5-y1.mggametransit.com",
    }
    
    // 发送带代理的POST请求
    baseURL := common.PLANT_H5
    api := "/api/Home/Login"
    
    respBody, resp, err := request.PostRequestCofigProxy(payload, baseURL, api, headers)
    
    if err != nil {
        fmt.Printf("请求失败: %v\n", err)
        return
    }
    
    fmt.Printf("请求成功! 状态码: %d\n", resp.StatusCode)
    fmt.Printf("响应内容: %s\n", string(respBody))
}
```

### 替换原有的PostRequestCofig调用

将原来的：
```go
resp, _, err := request.PostRequestCofig(payload, baseURL, api, headers)
```

替换为：
```go
resp, _, err := request.PostRequestCofigProxy(payload, baseURL, api, headers)
```

## 工作原理

1. **加载代理配置**: 从 `yaml/proxy_config.yaml` 文件中读取可用的代理列表
2. **尝试代理请求**: 按顺序尝试每个可用的代理
3. **自动重试**: 如果当前代理失败，自动尝试下一个代理
4. **降级处理**: 如果所有代理都失败，自动降级为直连请求
5. **错误处理**: 提供详细的错误信息和重试状态

## 代理管理

### 更新代理列表

```go
// 重新测试代理并更新YAML文件
request.ProxyFunc()
```

### 读取代理配置

```go
// 读取当前代理配置
config, err := request.LoadProxyConfig()
if err != nil {
    fmt.Printf("读取配置失败: %v\n", err)
    return
}

fmt.Printf("可用代理数量: %d\n", len(config.AvailableProxies))
```

### 获取随机代理

```go
// 获取一个随机的可用代理
randomProxy, err := request.GetRandomProxy()
if err != nil {
    fmt.Printf("获取随机代理失败: %v\n", err)
    return
}

fmt.Printf("随机代理: %s:%s\n", randomProxy.IP, randomProxy.Port)
```

## 特性

- ✅ **自动代理切换**: 代理失效时自动尝试下一个
- ✅ **降级处理**: 所有代理失败时自动使用直连
- ✅ **详细日志**: 显示代理尝试过程和结果
- ✅ **超时控制**: 每个代理请求设置30秒超时
- ✅ **错误处理**: 完善的错误处理和重试机制
- ✅ **签名支持**: 保持原有的签名功能

## 注意事项

1. 确保 `yaml/proxy_config.yaml` 文件存在且包含可用的代理
2. 代理的可用性可能会变化，建议定期更新代理列表
3. 如果所有代理都不可用，函数会自动降级为直连请求
4. 函数会保持与原有 `PostRequestCofig` 相同的接口和功能



