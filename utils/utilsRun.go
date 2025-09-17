package utils

import (
	"fmt"
	"sync"
)

// 需要传入需要生成多少一个id的个数，并且返回id的列表
func RandmoUserId(generateCount int) []string {
	// 模拟高并发生成100万个ID
	var wg sync.WaitGroup
	generated := sync.Map{} // 存储已生成的ID，检查重复
	collisionCount := 0
	// generateCount := 1000000
	idList := make([]string, 0, generateCount)
	for i := 0; i < generateCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id, err := RandmoUserCount()
			if err != nil {
				fmt.Printf("Error generating ID: %v\n", err)
				return
			}
			// 检查重复
			if _, exists := generated.LoadOrStore(id, true); exists {
				fmt.Printf("Collision detected: %s\n", id)
				collisionCount++
			}
			// 生成的用户id，可以进行接下来的操作
			idList = append(idList, id)
		}()
	}

	wg.Wait()
	fmt.Printf("Generated %d IDs with %d collisions\n", generateCount, collisionCount)
	return idList
}
