package study

type MyType int8

const demoType MyType = iota + 1

func IntType() bool {
	value := 10000000
	return demoType == MyType(value)
}
