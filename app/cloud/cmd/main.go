package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/file/:name", GetFile)
	if err := r.Run(":80"); err != nil {
		fmt.Println(err)
	}
}

var fileDir = "./app/cloud/config/"

// kubesphere-installer.yaml
// k3s-install.sh

func GetFile(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "no file"})
		return
	}
	filePath := fileDir + name

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.Data(http.StatusOK, "", b)
}
