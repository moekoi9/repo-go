package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.创建路由engine
	//创建engine，r就是 *Engine的结构体
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	fmt.Println("http://127.0.0.1:8000")

	r.Run(":8000")
}
