package biz

import "testing"

func TestApp(t *testing.T) {
	var targetValue = 4
	GlobalValue = targetValue
	for i := 0; i < 10000; i++ {
		if GlobalValue != targetValue {
			t.Errorf("expected:%v got:%v", targetValue, GlobalValue)
		}
	}
}
