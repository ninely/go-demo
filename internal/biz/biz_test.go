package biz

import "testing"

func TestMyBiz(t *testing.T) {
	var targetValue = 5
	GlobalValue = targetValue
	for i := 0; i < 10000; i++ {
		if GlobalValue != targetValue {
			t.Errorf("expected:%v got:%v", targetValue, GlobalValue)
		}
	}
}

type TranslateMock struct {
}

func (t *TranslateMock) Convert() {

}

func TestMyBiz1(t *testing.T) {
	type args struct {
		t Translate
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "A", args: args{t: &TranslateMock{}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyBiz(tt.args.t); got != tt.want {
				t.Errorf("MyBiz() = %v, want %v", got, tt.want)
			}
		})
	}
}
