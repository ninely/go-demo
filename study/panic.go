package study

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
