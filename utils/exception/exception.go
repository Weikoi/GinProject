package exception

// 断言接口
func Assert(condition bool, text string) {
	if !condition {
		panic(text)
	}
}
