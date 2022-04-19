package study

func ReadNil() int {
	var m map[int]int
	return m[3]
}

func WriteNil() (err error) {
	defer func() {
		err = recover().(error)
	}()
	var m map[int]bool
	m[1] = false
	return
}
