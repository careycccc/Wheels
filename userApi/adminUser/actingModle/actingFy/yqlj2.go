package actingFy

// import (
// 	"errors"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"path/filepath"
// 	"sync"
// 	"time"

// 	"gopkg.in/yaml.v3"
// )

// // User 定义用户结构
// type User struct {
// 	InviteCode   string   // 用户的邀请码
// 	Subordinates []string // 直接下级的邀请码列表
// }

// // userDB 模拟数据库，存储用户数据
// var userDB = make(map[string]*User)

// // dbMutex 确保并发安全
// var dbMutex sync.Mutex

// // fileMutex 确保文件写入的并发安全
// var fileMutex sync.Mutex

// // Layer 定义 YAML 文件中的层级结构
// type Layer struct {
// 	Level int      `yaml:"level"` // 层级编号
// 	Users []string `yaml:"users"` // 该层级的用户列表
// }

// // Layers 定义 YAML 文件的整体结构
// type Layers struct {
// 	Layers []Layer `yaml:"layers"` // 所有层级
// }

// // AA 将一个用户绑定到上级
// // inviteCode: 当前用户的邀请码
// // subordinates: 每层下级人数的切片，例如 []int{5, 8, 5, 6, 7}
// // level: 当前层级索引（从0开始）
// // newUsers: 收集所有新绑定的用户邀请码
// func AA(inviteCode string, subordinates []int, level int, newUsers *[]string) error {
// 	// 输入验证
// 	if inviteCode == "" {
// 		return errors.New("邀请码不能为空")
// 	}
// 	if len(subordinates) == 0 {
// 		return errors.New("层级人数切片不能为空")
// 	}
// 	if level >= len(subordinates) {
// 		return errors.New("层级索引超出切片范围")
// 	}

// 	// 绑定新用户
// 	newUserCode := generateNewInviteCode()
// 	{
// 		// 关键部分：更新 userDB
// 		dbMutex.Lock()
// 		// 获取或创建当前用户
// 		currentUser, exists := userDB[inviteCode]
// 		if !exists {
// 			currentUser = &User{InviteCode: inviteCode, Subordinates: []string{}}
// 			userDB[inviteCode] = currentUser
// 		}

// 		// 当前层级的下级人数（用于验证）
// 		requiredSubordinates := subordinates[level]

// 		// 检查当前用户是否可以绑定更多下级
// 		if len(currentUser.Subordinates) >= requiredSubordinates {
// 			dbMutex.Unlock()
// 			return fmt.Errorf("用户 %s 下级人数已满 (%d/%d)", inviteCode, len(currentUser.Subordinates), requiredSubordinates)
// 		}

// 		// 将新用户绑定到当前用户
// 		currentUser.Subordinates = append(currentUser.Subordinates, newUserCode)

// 		// 创建新用户
// 		userDB[newUserCode] = &User{InviteCode: newUserCode, Subordinates: []string{}}
// 		// 记录新用户
// 		*newUsers = append(*newUsers, newUserCode)
// 		fmt.Printf("绑定用户: %s (上级: %s, 层级: %d, 当前下级数: %d/%d)\n", newUserCode, inviteCode, level+1, len(currentUser.Subordinates), requiredSubordinates)
// 		dbMutex.Unlock()
// 	}

// 	return nil
// }

// // BB 模拟处理新用户的请求
// func BB(inviteCode string) error {
// 	fmt.Printf("处理用户 %s 的BB请求\n", inviteCode)
// 	return nil
// }

// // generateNewInviteCode 生成唯一的邀请码
// func generateNewInviteCode() string {
// 	dbMutex.Lock()
// 	defer dbMutex.Unlock()
// 	return fmt.Sprintf("CODE_%d", len(userDB)+1)
// }

// // writeLayerToYAML 异步将层级用户写入 YAML 文件
// func writeLayerToYAML(layers []Layer, filename string, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	// 确保文件写入的并发安全
// 	fileMutex.Lock()
// 	defer fileMutex.Unlock()

// 	// 序列化为 YAML
// 	data, err := yaml.Marshal(&Layers{Layers: layers})
// 	if err != nil {
// 		fmt.Printf("序列化 YAML 失败: %v\n", err)
// 		return
// 	}

// 	// 确保目录存在
// 	dir := filepath.Dir(filename)
// 	if err := os.MkdirAll(dir, 0755); err != nil {
// 		fmt.Printf("创建目录 %s 失败: %v\n", dir, err)
// 		return
// 	}

// 	// 以覆盖方式写入文件
// 	if err := os.WriteFile(filename, data, 0644); err != nil {
// 		fmt.Printf("写入文件 %s 失败: %v\n", filename, err)
// 		return
// 	}
// 	fmt.Printf("成功写入层级用户到 %s\n", filename)
// }

// // RunAAWithBB 为所有层级绑定用户并调用 BB
// func RunAAWithBB(inviteCode string, subordinates []int) error {
// 	// 初始化随机数生成器
// 	rand.Seed(time.Now().UnixNano())

// 	// 按层级收集所有用户
// 	layers := make([][]string, len(subordinates))
// 	yamlLayers := make([]Layer, len(subordinates)) // 用于 YAML 输出
// 	var wg sync.WaitGroup                          // 等待异步写入完成

// 	// 输入验证
// 	if len(subordinates) == 0 {
// 		return errors.New("层级人数切片不能为空")
// 	}

// 	// 绑定第一层用户
// 	for i := 0; i < subordinates[0]; i++ {
// 		fmt.Printf("绑定第一层用户 %d/%d\n", i+1, subordinates[0])
// 		newUsersForThisCall := []string{}
// 		if err := AA(inviteCode, subordinates, 0, &newUsersForThisCall); err != nil {
// 			return err
// 		}
// 		layers[0] = append(layers[0], newUsersForThisCall...)
// 	}
// 	// 验证第一层用户数
// 	if len(layers[0]) != subordinates[0] {
// 		return fmt.Errorf("层级 1 绑定用户数 %d 不等于预期 %d", len(layers[0]), subordinates[0])
// 	}
// 	// 更新 YAML 层级并异步写入
// 	yamlLayers[0] = Layer{Level: 1, Users: layers[0]}
// 	fmt.Printf("层级 1 用户列表: %v (总数: %d)\n", layers[0], len(layers[0]))
// 	wg.Add(1)
// 	go writeLayerToYAML(yamlLayers[:1], "./userList/userList.yaml", &wg)

// 	// 绑定后续层级
// 	currentUsers := layers[0]
// 	for level := 1; level < len(subordinates); level++ {
// 		targetCount := subordinates[level]
// 		availableParents := len(currentUsers)
// 		if availableParents == 0 {
// 			return fmt.Errorf("层级 %d 没有可用父节点", level+1)
// 		}

// 		// 将 targetCount 个用户随机绑定到可用父节点
// 		newLevelUsers := []string{}
// 		for i := 0; i < targetCount; i++ {
// 			// 随机选择一个父节点
// 			parentIndex := rand.Intn(availableParents)
// 			parentCode := currentUsers[parentIndex]
// 			fmt.Printf("为层级 %d 随机选择父节点 %s 绑定用户 %d/%d\n", level+1, parentCode, i+1, targetCount)
// 			newUsersForThisCall := []string{}
// 			if err := AA(parentCode, subordinates, level, &newUsersForThisCall); err != nil {
// 				return err
// 			}
// 			newLevelUsers = append(newLevelUsers, newUsersForThisCall...)
// 		}

// 		// 验证当前层级用户数
// 		if len(newLevelUsers) != targetCount {
// 			return fmt.Errorf("层级 %d 绑定用户数 %d 不等于预期 %d", level+1, len(newLevelUsers), targetCount)
// 		}

// 		// 更新层级和 YAML 数据
// 		layers[level] = newLevelUsers
// 		yamlLayers[level] = Layer{Level: level + 1, Users: newLevelUsers}
// 		fmt.Printf("层级 %d 用户列表: %v (总数: %d)\n", level+1, newLevelUsers, len(newLevelUsers))

// 		// 异步写入当前所有层级到 YAML 文件
// 		wg.Add(1)
// 		go writeLayerToYAML(yamlLayers[:level+1], "./userList/userList.yaml", &wg)

// 		// 更新当前用户
// 		currentUsers = newLevelUsers
// 	}

// 	// 合并所有用户
// 	newUsers := []string{}
// 	for _, layer := range layers {
// 		newUsers = append(newUsers, layer...)
// 	}

// 	// 记录收集的用户
// 	fmt.Printf("所有新绑定用户: %v (总数: %d)\n", newUsers, len(newUsers))

// 	// 为每个新用户调用 BB
// 	for _, userCode := range newUsers {
// 		if err := BB(userCode); err != nil {
// 			return fmt.Errorf("调用BB请求失败 for %s: %w", userCode, err)
// 		}
// 	}

// 	// 等待所有异步写入完成
// 	wg.Wait()
// 	return nil
// }

// // main 展示使用示例
// func RunYqlj2() {
// 	// 清空 userDB 以确保干净的状态
// 	userDB = make(map[string]*User)
// 	// 示例：任意层级
// 	subordinates := []int{5, 8, 5, 6, 7} // 可替换为任意切片，如 []int{3, 4, 2}
// 	err := RunAAWithBB("ROOT_CODE", subordinates)
// 	if err != nil {
// 		fmt.Println("操作失败:", err)
// 	} else {
// 		fmt.Println("操作成功")
// 	}
// }
