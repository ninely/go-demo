package main

import (
	"context"
	"crypto/rand"
	"demo/internal/biz"
	"demo/internal/conf"
	"demo/internal/data"
	"demo/internal/service"
	"demo/study"
	"encoding/base64"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"io"
	"log"
	"math/big"
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

type SubStruct struct {
	Name string
}

func (v *SubStruct) Run() {
	fmt.Println(v.Name)
}

type DStruct struct {
	Sub *SubStruct
}

func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}

const (
	// lowerLetters is the list of lowercase letters.
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// digits is the list of permitted digits.
	digits = "0123456789"

	length = 8
)

func main() {
	// ctx := context.Background()
	// rTest := NewRedisTest()
	//
	// used, usedHuman := rTest.InfoMemory(ctx)
	// fmt.Println("before: ", used, usedHuman)
	// rTest.InfoKeyspace(ctx)
	//
	// batch := 20 * 10000
	// rTest.WriteValue(ctx, 5120, batch)
	//
	// usedAfter, usedHumanAfter := rTest.InfoMemory(ctx)
	// usedEnd, _ := strconv.ParseInt(usedAfter, 10, 64)
	// usedStart, _ := strconv.ParseInt(used, 10, 64)
	// totalByte := usedEnd - usedStart
	// fmt.Println("after: ", usedAfter, usedHumanAfter)
	// rTest.InfoKeyspace(ctx)
	// fmt.Printf("total: %d total human: %0.2fM avg: %0.2f", totalByte, float64(totalByte)/1024/1024, float64(totalByte)/float64(batch))

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

	// res := study.AppendAhead()
	// fmt.Println(res)
	// url := "https://www.bilibili.com/"
	// var b []byte
	// ctx, cancel := chromedp.NewContext(
	// 	context.Background(),
	// 	// chromedp.WithDebugf(log.Printf),
	// )
	// defer cancel()
	// if err := chromedp.Run(ctx, fullScreenshot(url, 100, &b)); err != nil {
	// 	log.Fatal(err)
	// }
	// tmpPicPath := fmt.Sprintf("%s-%d.png", "render", time.Now().Unix())
	// err := ioutil.WriteFile(tmpPicPath, b, 0422)
	// if err != nil {
	// }

	// n := NewPasswdGeneratorImpl()
	// res, err := n.Generate(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res)

	study.SwitchError()
}

type passwdGeneratorImpl struct {
	reader io.Reader
}

func NewPasswdGeneratorImpl() *passwdGeneratorImpl {
	return &passwdGeneratorImpl{
		reader: rand.Reader,
	}
}

func (g *passwdGeneratorImpl) Generate(ctx context.Context) (string, error) {
	numDigits := length / 2
	numCharacters := length - numDigits
	var result string

	// Characters
	for i := 0; i < numCharacters; i++ {
		ch, err := randomElement(g.reader, lowerLetters)
		if err != nil {
			return "", err
		}

		result, err = randomInsert(g.reader, result, ch)
		if err != nil {
			return "", err
		}
	}

	// digits
	for i := 0; i < numDigits; i++ {
		d, err := randomElement(g.reader, digits)
		if err != nil {
			return "", err
		}
		result, err = randomInsert(g.reader, result, d)
		if err != nil {
			return "", err
		}
	}

	return result, nil
}

// randomInsert randomly inserts the given value into the given string.
func randomInsert(reader io.Reader, s, val string) (string, error) {
	if s == "" {
		return val, nil
	}

	n, err := rand.Int(reader, big.NewInt(int64(len(s)+1)))
	if err != nil {
		return "", err
	}
	i := n.Int64()
	return s[0:i] + val + s[i:], nil
}

// randomElement extracts a random element from the given string.
func randomElement(reader io.Reader, s string) (string, error) {
	n, err := rand.Int(reader, big.NewInt(int64(len(s))))
	if err != nil {
		return "", err
	}
	return string(s[n.Int64()]), nil
}

func RunDemo() {
	bc := &conf.Bootstrap{Data: &conf.Data{Database: &conf.Data_Database{
		Source: "root:abcd1234@tcp(127.0.0.1:3308)/fangbianmian?timeout=5s&readTimeout=5s&writeTimeout=5s&charset=utf8mb4&parseTime=true&loc=Local",
		Driver: "mysql",
	}}}

	db := data.NewDB(bc.Data)
	dao, _ := data.NewData(db)
	repo := data.NewRepository(dao)
	tm := data.NewTransaction(dao)

	srv := service.NewMyDemo(repo, tm)
	err := srv.DoSomeBusiness(context.Background())
	if err != nil {
		log.Printf("%+v", err)
	}
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
