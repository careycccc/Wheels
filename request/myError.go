package request

/// 处理 各种错误

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

// 处理 try again late
func SomeFunction() error {
	return &MyError{"请求失败: Requests are too frequent, Please try again late"}
}
