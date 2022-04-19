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
