package study

func WaitPanic() error {
	go func() {
		panic("wait panic")
	}()
	select {}
	return nil
}

func NotWaitPanic() error {
	go func() {
		panic("not wait panic")
	}()
	return nil
}
