package main

import (
	"context"
	data2 "demo/internal/data"
	"demo/internal/service"
	"encoding/base64"
	"fmt"
	"os"
	"path"
	"strings"
)

func demo() {
	db := data2.NewDB()
	dao, _ := data2.NewData(db)
	repo := data2.NewRepository(dao)
	tm := data2.NewTransaction(dao)

	srv := service.NewMyDemo(repo, tm)
	_ = srv.DoSomeBusiness(context.Background())
}

func main() {
	//fileName := "abc.doc"
	//fmt.Println(path.Ext(fileName))
	//createTime := time.Now()
	//token := xid.NewWithTime(createTime).String()
	//fmt.Println(token)

	fmt.Println(getFileDat())

	//a := "ABC"
	//b := "ABCD"
	//distance := pkg.ComputeDistance(a, b)
	//fmt.Println(distance == Levenshtein(a, b))
}

func Levenshtein(a, b string) int {
	lenA := len(a)
	lenB := len(b)
	dp := make([][]int, lenA+1)
	for i := range dp {
		dp[i] = make([]int, lenB+1)
	}

	for i := 0; i <= lenA; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lenB; j++ {
		dp[0][j] = j
	}

	for i := 1; i < lenA; i++ {
		for j := 1; j < lenB; j++ {
			prev := min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			if a[i-1] != b[j-1] {
				prev += 1
			}
			dp[i][j] = prev
		}
	}

	return dp[lenA-1][lenB-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func fileTest() {
	srcFileName := "../../root/run.sh"
	ext := path.Ext(path.Base(srcFileName))
	filePrefix := strings.TrimSuffix(path.Base(srcFileName), ext)
	fmt.Println(filePrefix)
}

func getFileDat() string {
	fileName := "徐主峰 简历.pdf"
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
