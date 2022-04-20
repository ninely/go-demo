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
