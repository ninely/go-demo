package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", getK3S)
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println(err)
		return
	}
}

func getK3S(w http.ResponseWriter, r *http.Request) {
	filePath := "../config/k3s-install.sh"
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println(err)
		return
	}
}
