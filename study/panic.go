package study

import "fmt"

// WaitPanic 退出之前等待 panic
func WaitPanic() error {
	go func() {
		panic("wait panic")
	}()
	select {}
}

func NotWaitPanic() error {
	go func() {
		panic("not wait panic")
	}()
	return nil
}

func CatchPanic() error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("my error: %v", r)
		}
	}()

	go func() {
		panic("xxxxxxx")
	}()
	select {}
}
