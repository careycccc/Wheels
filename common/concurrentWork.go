package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

// Executor 定义执行函数类型
type Executor func(string) error

// ProcessAccounts 封装的处理函数
func ProcessAccounts(filePath string, executor Executor, concurrencyLimit int) error {
	// 创建 channel 和 WaitGroup
	accountChan := make(chan string, concurrencyLimit)
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, concurrencyLimit)

	// 记录开始时间
	start := time.Now()

	// 启动 goroutine 读取文件并发送到 channel
	go func() {
		err := readTextToChannel(filePath, accountChan)
		if err != nil {
			log.Printf("读取文件失败: %v", err)
		}
	}()

	// 启动 goroutine 处理 channel 中的请求
	go processRequests(accountChan, &wg, semaphore, executor)

	// 等待所有请求完成
	wg.Wait()
	fmt.Printf("所有请求完成，耗时: %v\n", time.Since(start))
	return nil
}

// ReadYAMLByLine 逐行读取文件并返回字符串切片
func ReadYAMLByLine(filePath string) ([]string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %v", err)
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
		return nil, fmt.Errorf("扫描文件失败: %v", err)
	}

	return lines, nil
}

// readTextToChannel 从纯文本文件逐行读取并发送到 channel
func readTextToChannel(filename string, accountChan chan<- string) error {
	defer close(accountChan) // 确保 channel 在函数退出时关闭

	// 使用 ReadYAMLByLine 读取文件
	lines, err := ReadYAMLByLine(filename)
	if err != nil {
		return err
	}

	// 将账号发送到 channel
	for _, account := range lines {
		account = strings.TrimSpace(account)
		if len(account) != 12 {
			log.Printf("账号 %s 不是12位，跳过", account)
			continue
		}
		accountChan <- account
	}

	return nil
}

// processRequests 处理 channel 中的账号并执行请求
func processRequests(accountChan <-chan string, wg *sync.WaitGroup, semaphore chan struct{}, executor Executor) {
	for account := range accountChan {
		semaphore <- struct{}{} // 获取信号量
		wg.Add(1)
		go func(acc string) {
			log.Printf("账号 %s 请求成功", acc)
			defer wg.Done()
			defer func() { <-semaphore }() // 释放信号量
			// 执行传入的函数
			err := executor(acc)
			if err != nil {
				log.Printf("账号 %s 请求失败: %v", acc, err)
				return
			}
			// log.Printf("账号 %s 请求成功", acc)
		}(account)
	}
}
