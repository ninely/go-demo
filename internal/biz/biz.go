package biz

var GlobalValue int

type Translate interface {
	Convert()
}

func MyBiz(t Translate) int {
	return GlobalValue
}
