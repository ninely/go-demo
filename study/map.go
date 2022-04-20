package study

// ReadNil 读未初始化 map 返回对应零值
// return 0
func ReadNil() int {
	var m map[int]int
	return m[3]
}

// WriteNil 写未初始化 map panic
// return assignment to entry in nil map
func WriteNil() (err error) {
	defer func() {
		err = recover().(error)
	}()
	var m map[int]bool
	m[1] = false
	return
}
