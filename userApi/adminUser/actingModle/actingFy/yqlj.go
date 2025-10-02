package actingFy

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// User 定义用户结构
type User struct {
	InviteCode   string   // 用户的邀请码
	Subordinates []string // 直接下级的邀请码列表
}

// userDB 模拟数据库，存储用户数据
var userDB = make(map[string]*User)

// dbMutex 确保并发安全
var dbMutex sync.Mutex

// AA 将一个用户绑定到上级
// inviteCode: 当前用户的邀请码
// subordinates: 每层下级人数的切片，例如 []int{3, 4, 2}
// level: 当前层级索引（从0开始）
// newUsers: 收集所有新绑定的用户邀请码
func AA(inviteCode string, subordinates []int, level int, newUsers *[]string) error {
	// 输入验证
	if inviteCode == "" {
		return errors.New("邀请码不能为空")
	}
	if len(subordinates) == 0 {
		return errors.New("层级人数切片不能为空")
	}
	if level >= len(subordinates) {
		return errors.New("层级索引超出切片范围")
	}

	// 绑定新用户
	newUserCode := generateNewInviteCode()
	{
		// 关键部分：更新 userDB
		dbMutex.Lock()
		// 获取或创建当前用户
		currentUser, exists := userDB[inviteCode]
		if !exists {
			currentUser = &User{InviteCode: inviteCode, Subordinates: []string{}}
			userDB[inviteCode] = currentUser
		}

		// 当前层级的下级人数（用于验证）
		requiredSubordinates := subordinates[level]

		// 检查当前用户是否可以绑定更多下级
		if len(currentUser.Subordinates) >= requiredSubordinates {
			dbMutex.Unlock()
			return fmt.Errorf("用户 %s 下级人数已满 (%d/%d)", inviteCode, len(currentUser.Subordinates), requiredSubordinates)
		}

		// 将新用户绑定到当前用户
		currentUser.Subordinates = append(currentUser.Subordinates, newUserCode)

		// 创建新用户
		userDB[newUserCode] = &User{InviteCode: newUserCode, Subordinates: []string{}}
		// 记录新用户
		*newUsers = append(*newUsers, newUserCode)
		//fmt.Printf("绑定用户: %s (上级: %s, 层级: %d, 当前下级数: %d/%d)\n", newUserCode, inviteCode, level+1, len(currentUser.Subordinates), requiredSubordinates)
		dbMutex.Unlock()
	}

	return nil
}

// BB 模拟处理新用户的请求
func BB(inviteCode string) error {
	// fmt.Printf("处理用户 %s 的BB请求\n", inviteCode)
	return nil
}

// generateNewInviteCode 生成唯一的邀请码
func generateNewInviteCode() string {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	return fmt.Sprintf("CODE_%d", len(userDB)+1)
}

// RunAAWithBB 为所有层级绑定用户并调用 BB
func RunAAWithBB(inviteCode string, subordinates []int) error {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 按层级收集所有用户
	layers := make([][]string, len(subordinates))

	// 输入验证
	if len(subordinates) == 0 {
		return errors.New("层级人数切片不能为空")
	}

	// 绑定第一层用户
	for i := 0; i < subordinates[0]; i++ {
		//fmt.Printf("绑定第%d层用户 %d/%d\n", i+1, i+1, subordinates[0])
		newUsersForThisCall := []string{}
		if err := AA(inviteCode, subordinates, 0, &newUsersForThisCall); err != nil {
			return err
		}
		layers[0] = append(layers[0], newUsersForThisCall...)
	}
	// 验证第一层用户数
	if len(layers[0]) != subordinates[0] {
		return fmt.Errorf("层级 1 绑定用户数 %d 不等于预期 %d", len(layers[0]), subordinates[0])
	}
	// 打印第一层用户列表
	fmt.Printf("层级 1 用户列表: %v (总数: %d)\n", layers[0], len(layers[0]))

	// 绑定后续层级
	currentUsers := layers[0]
	for level := 1; level < len(subordinates); level++ {
		targetCount := subordinates[level]
		availableParents := len(currentUsers)
		if availableParents == 0 {
			return fmt.Errorf("层级 %d 没有可用父节点", level+1)
		}

		// 将 targetCount 个用户随机绑定到可用父节点
		newLevelUsers := []string{}
		for i := 0; i < targetCount; i++ {
			// 随机选择一个父节点
			parentIndex := rand.Intn(availableParents)
			parentCode := currentUsers[parentIndex]
			// fmt.Printf("为层级 %d 随机选择父节点 %s 绑定用户 %d/%d\n", level+1, parentCode, i+1, targetCount)
			newUsersForThisCall := []string{}
			if err := AA(parentCode, subordinates, level, &newUsersForThisCall); err != nil {
				return err
			}
			newLevelUsers = append(newLevelUsers, newUsersForThisCall...)
		}

		// 验证当前层级用户数
		if len(newLevelUsers) != targetCount {
			return fmt.Errorf("层级 %d 绑定用户数 %d 不等于预期 %d", level+1, len(newLevelUsers), targetCount)
		}

		// 打印当前层级用户列表
		fmt.Printf("层级 %d 用户列表: %v (总数: %d)\n", level+1, newLevelUsers, len(newLevelUsers))

		// 更新层级和当前用户
		layers[level] = newLevelUsers
		currentUsers = newLevelUsers
	}

	// 合并所有用户
	newUsers := []string{}
	for _, layer := range layers {
		newUsers = append(newUsers, layer...)
	}

	// 记录收集的用户
	// fmt.Printf("所有新绑定用户: %v (总数: %d)\n", newUsers, len(newUsers))

	// 为每个新用户调用 BB
	for _, userCode := range newUsers {
		if err := BB(userCode); err != nil {
			return fmt.Errorf("调用BB请求失败 for %s: %w", userCode, err)
		}
	}

	return nil
}

// main 展示使用示例
func RunInvite() {
	// 清空 userDB 以确保干净的状态
	userDB = make(map[string]*User)
	// 示例：任意层级
	subordinates := []int{5, 8, 5, 6, 9} // 可替换为任意切片，如 []int{3, 4, 2}
	err := RunAAWithBB("ROOT_CODE", subordinates)
	if err != nil {
		fmt.Println("操作失败:", err)
	} else {
		fmt.Println("操作成功")
	}
}
