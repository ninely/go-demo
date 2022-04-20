package main

import (
	"context"
	"demo/internal/biz"
	"demo/internal/data"
	"demo/internal/service"
	"demo/study"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"reflect"
	"runtime"
)

func RunStudy(handler interface{}) {
	funcValue := reflect.ValueOf(handler)
	retList := funcValue.Call(nil)
	funcName := runtime.FuncForPC(funcValue.Pointer()).Name()

	var result string
	for _, v := range retList {
		result += fmt.Sprintf("%v", v.Interface())
	}

	fmt.Printf("%s: %s\n", funcName, result)
}

func main() {
	RunStudy(study.AppendNil)
	RunStudy(study.RefAll)
	RunStudy(study.ReadNil)
	RunStudy(study.WriteNil)

	RunStudy(study.EmptyToByte)

	RunStudy(study.WriteNil)
	RunStudy(study.FilePath)
	RunStudy(study.LevCompare)
	RunStudy(study.IntType)

	RunStudy(study.MarshalNil)
	RunStudy(study.MarshalStruct)
	RunStudy(study.MarshalPointStruct)
	RunStudy(study.MarshalStructEmpty)
	RunStudy(study.MarshalPointStructEmpty)

	// RunStudy(study.NotWaitPanic)
	// RunStudy(study.WaitPanic)

	// RunResume()

	// RunDemo()

	// RunErr()
}

func RunDemo() {
	db := data.NewDB()
	dao, _ := data.NewData(db)
	repo := data.NewRepository(dao)
	tm := data.NewTransaction(dao)

	srv := service.NewMyDemo(repo, tm)
	_ = srv.DoSomeBusiness(context.Background())
}

func RunErr() {
	err := errors.Wrap(biz.ErrDataNotFound, "test")
	fmt.Println(biz.IsDataNotFoundError(err))
}

func RunResume() {
	fmt.Println(getFileDat())
}

func getFileDat() string {
	fileName := "张昊杰.pdf"
	filePath := "/Users/melody/Downloads/pdf/" + fileName
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(data)
}
