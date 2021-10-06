package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func main()  {
	route:=gin.Default()
	route.GET("/pingo_test", Pull)
	err := route.Run(":9998")
	if err != nil {
		return 
	}
}

func Pull(c *gin.Context) {
	command :="./pingo_test.sh"
	cmd:=exec.Command("/bin/bash", "-c", command)
	fmt.Println(cmd)
	c.JSON(http.StatusOK, gin.H{"info":cmd})
}