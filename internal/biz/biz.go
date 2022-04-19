package biz

import "errors"

var (
	ErrDataNotFound = errors.New("data not found")
)

func IsDataNotFoundError(err error) bool {
	return errors.Is(err, ErrDataNotFound)
}

var GlobalValue int

type Translate interface {
	Convert()
}

func MyBiz(t Translate) int {
	return GlobalValue
}
