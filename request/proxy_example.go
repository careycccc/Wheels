package request

import (
	"fmt"
)

func GetProxyIpFunc() {
	fmt.Println("=== 代理IP管理示例 ===")

	// 1. 测试代理并保存到YAML文件
	fmt.Println("1. 开始测试代理IP...")
	ProxyFunc()

	// 2. 读取YAML文件中的代理配置
	fmt.Println("\n2. 读取YAML文件中的代理配置...")
	config, err := LoadProxyConfig()
	if err != nil {
		fmt.Printf("读取配置失败: %v\n", err)
		return
	}

	fmt.Printf("最后更新时间: %s\n", config.LastUpdated)
	fmt.Printf("可用代理数量: %d\n", len(config.AvailableProxies))

	// 3. 显示前5个可用的代理
	fmt.Println("\n3. 前5个可用的代理:")
	for i, proxy := range config.AvailableProxies {
		if i >= 5 {
			break
		}
		fmt.Printf("  %d. %s:%s (%s) - 测试时间: %s\n",
			i+1, proxy.IP, proxy.Port, proxy.Protocol, proxy.TestTime)
	}

	// 4. 获取一个随机代理
	fmt.Println("\n4. 获取随机代理:")
	randomProxy, err := GetRandomProxy()
	if err != nil {
		fmt.Printf("获取随机代理失败: %v\n", err)
		return
	}

	fmt.Printf("随机代理: %s:%s (%s)\n",
		randomProxy.IP, randomProxy.Port, randomProxy.Protocol)
}
