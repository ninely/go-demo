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
)

func main() {
	fmt.Println(study.AppendNil())
	fmt.Println(study.ReadNil())
	fmt.Println(study.EmptyToByte())
	fmt.Println(study.MarshalNil())
	fmt.Println(study.RefAll())
	fmt.Println(study.WriteNil())
	fmt.Println(study.MarshalStruct())
	fmt.Println(study.MarshalPointStruct())
	fmt.Println(study.MarshalStructEmpty())
	fmt.Println(study.MarshalPointStructEmpty())
	fmt.Println(study.FilePath())
	fmt.Println(study.LevCompare())
	fmt.Println(study.IntType())

	//fmt.Println(study.NotWaitPanic())
	//fmt.Println(study.WaitPanic())

	//RunResume()

	//RunDemo()

	//RunErr()
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
