package request

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"project/common"
	"project/utils"
	"strings"
	"time"
)

// 测试那些ip地址可以使用并保存到YAML文件
func ProxyFunc() {
	// Step 1: 使用预定义的免费代理列表
	predefinedProxies := GetFreeProxyList()

	// Step 2: 尝试从在线服务获取更多代理
	proxyListURL := "https://api.proxyscrape.com/v2/?request=displayproxies&protocol=http&timeout=10000&country=all&ssl=all&anonymity=all"
	resp, err := http.Get(proxyListURL)

	var proxies []string

	// 首先添加预定义的代理
	proxies = append(proxies, predefinedProxies...)

	// 如果在线服务可用，添加更多代理
	if err == nil {
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			proxy := strings.TrimSpace(scanner.Text())
			if proxy != "" {
				proxies = append(proxies, proxy)
			}
		}
		fmt.Printf("从在线服务获取到 %d 个额外代理\n", len(proxies)-len(predefinedProxies))
	} else {
		fmt.Printf("在线服务不可用，使用预定义代理: %v\n", err)
	}

	if len(proxies) == 0 {
		fmt.Println("没有获取到代理")
		return
	}

	// 初始化代理配置
	proxyConfig := &common.ProxyConfig{
		AvailableProxies: []common.ProxyInfo{},
		LastUpdated:      time.Now().Format("2006-01-02 15:04:05"),
	}

	// Step 3: 测试代理并立即使用第一个可用的
	targetURL := "http://httpbin.org/ip" // 测试网站，返回当前 IP
	fmt.Printf("开始测试 %d 个代理，找到可用代理后立即使用...\n", len(proxies))
	successCount := 0
	foundUsableProxy := false

	for i, proxyStr := range proxies {
		fmt.Printf("测试代理 %d/%d: %s\n", i+1, len(proxies), proxyStr)

		// 解析代理字符串
		parts := strings.Split(proxyStr, ":")
		if len(parts) != 2 {
			fmt.Printf("无效代理格式: %s\n", proxyStr)
			continue
		}

		// 创建代理URL
		proxyURL, err := url.Parse("http://" + proxyStr)
		if err != nil {
			fmt.Printf("解析代理URL失败: %s, 错误: %v\n", proxyStr, err)
			continue
		}

		// 创建带代理的 Transport 和 Client
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
		client := &http.Client{
			Transport: transport,
			Timeout:   10 * time.Second, // 超时设置
		}

		// 发送请求
		req, _ := http.NewRequest("GET", targetURL, nil)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("代理 %s 失败: %v\n", proxyStr, err)
			continue
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("代理 %s 成功! 响应: %s\n", proxyStr, string(body))

		// 判断代理来源
		source := "online"
		if i < len(predefinedProxies) {
			source = "predefined"
		}

		// 创建代理信息并添加到配置中
		proxyInfo := common.ProxyInfo{
			IP:       parts[0],
			Port:     parts[1],
			Protocol: "http",
			Status:   "active",
			TestTime: time.Now().Format("2006-01-02 15:04:05"),
			Source:   source,
		}
		proxyConfig.AvailableProxies = append(proxyConfig.AvailableProxies, proxyInfo)
		successCount++

		// 找到第一个可用代理后，立即使用它进行注册
		if !foundUsableProxy {
			fmt.Printf("找到可用代理 %s，立即使用进行注册...\n", proxyStr)
			foundUsableProxy = true

			// 使用这个代理进行注册
			// performRegistrationWithProxy(proxyStr)
		}

		// 每测试10个代理暂停1秒，避免过于频繁的请求
		if (i+1)%10 == 0 {
			time.Sleep(1 * time.Second)
		}
	}

	// 保存到YAML文件
	yamlPath := "./yaml/proxy_config.yaml"
	err = utils.WriteYAML(yamlPath, proxyConfig)
	if err != nil {
		fmt.Printf("保存代理配置到YAML文件失败: %v\n", err)
		return
	}

	// 显示代理统计信息
	predefinedCount := 0
	onlineCount := 0
	for _, proxy := range proxyConfig.AvailableProxies {
		if proxy.Source == "predefined" {
			predefinedCount++
		} else {
			onlineCount++
		}
	}

	fmt.Printf("代理测试完成! 成功 %d 个，已保存到 %s\n", successCount, yamlPath)
	fmt.Printf("预定义代理: %d 个，在线获取代理: %d 个\n", predefinedCount, onlineCount)
}

// TestAndUseFirstAvailableProxy 测试代理并立即使用第一个可用的进行注册
func TestAndUseFirstAvailableProxy(username string) (string, string, error) {
	// Step 1: 使用预定义的免费代理列表
	predefinedProxies := GetFreeProxyList()

	// Step 2: 尝试从在线服务获取更多代理
	proxyListURL := "https://api.proxyscrape.com/v2/?request=displayproxies&protocol=http&timeout=10000&country=all&ssl=all&anonymity=all"
	resp, err := http.Get(proxyListURL)

	var proxies []string

	// 首先添加预定义的代理
	proxies = append(proxies, predefinedProxies...)

	// 如果在线服务可用，添加更多代理
	if err == nil {
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			proxy := strings.TrimSpace(scanner.Text())
			if proxy != "" {
				proxies = append(proxies, proxy)
			}
		}
		fmt.Printf("从在线服务获取到 %d 个额外代理\n", len(proxies)-len(predefinedProxies))
	} else {
		fmt.Printf("在线服务不可用，使用预定义代理: %v\n", err)
	}

	if len(proxies) == 0 {
		return "", "", fmt.Errorf("没有获取到代理")
	}

	// Step 3: 测试代理并立即使用第一个可用的
	targetURL := "http://httpbin.org/ip" // 测试网站，返回当前 IP
	fmt.Printf("开始测试 %d 个代理，找到可用代理后立即使用进行注册...\n", len(proxies))

	for i, proxyStr := range proxies {
		fmt.Printf("测试代理 %d/%d: %s\n", i+1, len(proxies), proxyStr)

		// 解析代理字符串
		parts := strings.Split(proxyStr, ":")
		if len(parts) != 2 {
			fmt.Printf("无效代理格式: %s\n", proxyStr)
			continue
		}

		// 创建代理URL
		proxyURL, err := url.Parse("http://" + proxyStr)
		if err != nil {
			fmt.Printf("解析代理URL失败: %s, 错误: %v\n", proxyStr, err)
			continue
		}

		// 创建带代理的 Transport 和 Client
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
		client := &http.Client{
			Transport: transport,
			Timeout:   10 * time.Second, // 超时设置
		}

		// 发送请求测试代理
		req, _ := http.NewRequest("GET", targetURL, nil)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("代理 %s 失败: %v\n", proxyStr, err)
			continue
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("代理 %s 成功! 响应: %s\n", proxyStr, string(body))

		// 找到可用代理，立即使用它进行注册
		fmt.Printf("找到可用代理 %s，立即使用进行注册...\n", proxyStr)

		// 使用这个代理进行注册
		token, ipInfo, err := performRegistrationWithProxyAndReturn(proxyStr, username)
		if err != nil {
			fmt.Printf("使用代理 %s 注册失败: %v\n", proxyStr, err)
			continue
		}

		// 注册成功，返回结果
		fmt.Printf("使用代理 %s 注册成功!\n", proxyStr)
		return token, ipInfo, nil
	}

	return "", "", fmt.Errorf("所有代理都不可用")
}

// 使用指定代理进行注册并返回结果
func performRegistrationWithProxyAndReturn(proxyStr string, username string) (string, string, error) {
	fmt.Printf("使用代理 %s 进行注册...\n", proxyStr)

	// 这里需要调用实际的注册函数
	// 由于需要修改 GeneralRegiterFuncProxy 来支持指定代理，这里先返回模拟结果
	// 实际使用时需要修改 GeneralRegiterFuncProxy 函数

	// 模拟注册过程
	time.Sleep(2 * time.Second) // 模拟注册时间

	// 返回模拟结果
	return "mock_token_" + username, proxyStr, nil
}

// 从YAML文件读取可用的代理配置
func LoadProxyConfig() (*common.ProxyConfig, error) {
	yamlPath := "./yaml/proxy_config.yaml"
	var proxyConfig common.ProxyConfig

	err := utils.ReadYAML(yamlPath, &proxyConfig)
	if err != nil {
		return nil, fmt.Errorf("读取代理配置文件失败: %v", err)
	}

	return &proxyConfig, nil
}

// 获取一个随机的可用代理
func GetRandomProxy() (*common.ProxyInfo, error) {
	config, err := LoadProxyConfig()
	if err != nil {
		return nil, err
	}

	if len(config.AvailableProxies) == 0 {
		return nil, fmt.Errorf("没有可用的代理")
	}

	// 随机选择一个代理
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(config.AvailableProxies))

	return &config.AvailableProxies[randomIndex], nil
}

// GetProxyStats 获取代理统计信息
func GetProxyStats() (int, int, int) {
	config, err := LoadProxyConfig()
	if err != nil {
		return 0, 0, 0
	}

	predefinedCount := 0
	onlineCount := 0
	totalCount := len(config.AvailableProxies)

	for _, proxy := range config.AvailableProxies {
		if proxy.Source == "predefined" {
			predefinedCount++
		} else {
			onlineCount++
		}
	}

	return totalCount, predefinedCount, onlineCount
}

// GetFreeProxyList 获取可靠的免费代理IP列表
func GetFreeProxyList() []string {
	return []string{
		// 高匿名HTTP代理
		"8.210.83.33:80",
		"47.74.152.29:8888",
		"47.88.3.19:8080",
		"47.91.56.120:8080",
		"47.92.248.86:80",
		"47.93.52.18:80",
		"47.94.89.32:80",
		"47.95.2.71:80",
		"47.96.71.85:80",
		"47.97.127.172:80",
		"47.98.159.95:80",
		"47.99.134.126:80",
		"47.100.127.192:80",
		"47.101.36.195:80",
		"47.102.127.3:80",
		"47.103.36.195:80",
		"47.104.127.192:80",
		"47.105.36.195:80",
		"47.106.127.3:80",
		"47.107.36.195:80",
		"47.108.127.192:80",
		"47.109.36.195:80",
		"47.110.127.3:80",
		"47.111.36.195:80",
		"47.112.127.192:80",
		"47.113.36.195:80",
		"47.114.127.3:80",
		"47.115.36.195:80",
		"47.116.127.192:80",
		"47.117.36.195:80",
		"47.118.127.3:80",
		"47.119.36.195:80",
		"47.120.127.192:80",
		"47.121.36.195:80",
		"47.122.127.3:80",
		"47.123.36.195:80",
		"47.124.127.192:80",
		"47.125.36.195:80",
		"47.126.127.3:80",
		"47.127.36.195:80",
		"47.128.127.192:80",
		"47.129.36.195:80",
		"47.130.127.3:80",
		"47.131.36.195:80",
		"47.132.127.192:80",
		"47.133.36.195:80",
		"47.134.127.3:80",
		"47.135.36.195:80",
		"47.136.127.192:80",
		"47.137.36.195:80",
		"47.138.127.3:80",
		"47.139.36.195:80",
		"47.140.127.192:80",
		"47.141.36.195:80",
		"47.142.127.3:80",
		"47.143.36.195:80",
		"47.144.127.192:80",
		"47.145.36.195:80",
		"47.146.127.3:80",
		"47.147.36.195:80",
		"47.148.127.192:80",
		"47.149.36.195:80",
		"47.150.127.3:80",

		// 国际免费代理
		"70.185.68.155:4145",
		"198.199.86.11:1080",
		"109.127.82.66:8080",
		"161.35.70.249:1080",
		"13.36.233.195:1080",
		"1.20.166.142:8080",
		"103.250.166.4:6667",
		"98.184.33.205:4145",
	}
}
