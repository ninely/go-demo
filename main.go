package main

import (
	"context"
	"demo/data"
	"demo/pkg"
	"demo/service"
	"encoding/base64"
	"fmt"
	"os"
	"path"
	"strings"
)

func demo() {
	db := data.NewDB()
	dao, _ := data.NewData(db)
	repo := data.NewRepository(dao)
	tm := data.NewTransaction(dao)

	srv := service.NewMyDemo(repo, tm)
	_ = srv.DoSomeBusiness(context.Background())
}

func main() {
	//fmt.Println(getFileDat())
	a := "ABC"
	b := "BCDE"
	distance := pkg.ComputeDistance(a, b)
	fmt.Println(distance)
}

// Levenshtein abc abcd
// dp[i][j]
func Levenshtein(a, b string) int {

	return 0
}

func fileTest() {
	srcFileName := "../../root/run.sh"
	ext := path.Ext(path.Base(srcFileName))
	filePrefix := strings.TrimSuffix(path.Base(srcFileName), ext)
	fmt.Println(filePrefix)
}

func getFileDat() string {
	fileName := "_0.jpg"
	filePath := "/Users/melody/Downloads/pdf/" + fileName
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(data)
}

type MyDemo struct {
	Name string
}

func (d *MyDemo) ShowA() string {
	return fmt.Sprintf("ShowA:%v", d.Name)
}

func (d *MyDemo) ShowB() string {
	return fmt.Sprintf("ShowB:%v", d.Name)
}

func test() {
	handler := []func(*MyDemo) string{
		(*MyDemo).ShowA,
		(*MyDemo).ShowB,
	}

	d := &MyDemo{Name: "Test"}
	for _, v := range handler {
		result := v(d)
		fmt.Println(result)
	}

	var i interface{}
	i = nil
	value, ok := (i).([]*MyDemo)
	fmt.Println(value, ok)
}
