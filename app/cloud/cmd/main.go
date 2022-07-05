package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/k3s", getK3S)
	http.HandleFunc("/k3s-sha", getK3SSha)
	http.HandleFunc("/kilo", getKiloK3S)
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

func getK3SSha(w http.ResponseWriter, r *http.Request) {
	filePath := "../config/k3s-sha256sum-amd64.txt"
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

func getKiloK3S(w http.ResponseWriter, r *http.Request) {
	filePath := "../config/kilo-k3s.yaml"
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
