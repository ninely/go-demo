package study

import (
	"encoding/json"
)

type SubStruct struct {
	Name string `json:"name"`
}

type MainStructPoint struct {
	Data *SubStruct `json:"data"`
}

type MainStruct struct {
	Data SubStruct `json:"data"`
}

func MarshalNil() (string, error) {
	b, err := json.Marshal(nil)
	return string(b), err
}

// MarshalPointStruct
// return {"data":{"name":""}}
func MarshalPointStruct() string {
	var s SubStruct
	data := &s
	a := &MainStructPoint{data}
	b, err := json.Marshal(a)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// MarshalPointStructEmpty 对于内嵌的结构体指针，在初始化时未显式初始化内嵌字段，Marshal 结果字段不输出
// return {"data":null}
func MarshalPointStructEmpty() string {
	m := &MainStructPoint{}
	b, err := json.Marshal(m)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func MarshalStruct() string {
	var s SubStruct
	a := &MainStruct{s}
	b, err := json.Marshal(a)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func MarshalStructEmpty() string {
	m := &MainStruct{}
	b, err := json.Marshal(m)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
