package main

import (
	"context"
	"demo/internal/biz"
	"demo/internal/data"
	"demo/internal/service"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"reflect"
	"runtime"
	"strconv"
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

type SubStruct struct {
	Name string
}

func (v *SubStruct) Run() {
	fmt.Println(v.Name)
}

type DStruct struct {
	Sub *SubStruct
}

func main() {
	ctx := context.Background()
	rTest := NewRedisTest()

	used, usedHuman := rTest.InfoMemory(ctx)
	fmt.Println("before: ", used, usedHuman)
	rTest.InfoKeyspace(ctx)

	batch := 20 * 10000
	rTest.WriteValue(ctx, 20, batch)

	usedAfter, usedHumanAfter := rTest.InfoMemory(ctx)
	usedEnd, _ := strconv.ParseInt(usedAfter, 10, 64)
	usedStart, _ := strconv.ParseInt(used, 10, 64)
	totalByte := usedEnd - usedStart
	fmt.Println("after: ", usedAfter, usedHumanAfter)
	rTest.InfoKeyspace(ctx)
	fmt.Printf("total: %d total human: %0.2fM avg: %0.2f", totalByte, float64(totalByte)/1024/1024, float64(totalByte)/float64(batch))

	// d := DStruct{}
	// d.Sub.Run()
	// RunStudy(study.AppendNil)
	// RunStudy(study.RefAll)
	// RunStudy(study.ReadNil)
	// RunStudy(study.WriteNil)
	//
	// RunStudy(study.EmptyToByte)
	//
	// RunStudy(study.WriteNil)
	// RunStudy(study.FilePath)
	// RunStudy(study.LevCompare)
	// RunStudy(study.IntType)
	//
	// RunStudy(study.MarshalNil)
	// RunStudy(study.MarshalStruct)
	// RunStudy(study.MarshalPointStruct)
	// RunStudy(study.MarshalStructEmpty)
	// RunStudy(study.MarshalPointStructEmpty)

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
