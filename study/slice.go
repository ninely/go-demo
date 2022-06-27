package study

// AppendNil ... 运算符支持未初始化的 slice
func AppendNil() []int {
	var arr []int
	var brr []int
	brr = append(brr, arr...)
	return brr
}

// RefAll [:] 仍然指向相同底层数组
func RefAll() bool {
	var s1 = []int{1, 2}
	s2 := s1[:]
	s2[0] = 3
	return s1[0] == 1
}

// AppendAhead 头插法
func AppendAhead() []int {
	var size = 5
	var s = make([]int, 0, size)
	for i := 1; i <= size; i++ {
		s = append(s, 0)
		copy(s[1:], s[:])
		s[0] = i
	}
	return s
}
