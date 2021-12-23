package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("html/tem/*")
	r.LoadHTMLGlob("tem2/**/**/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.html", gin.H{"title": "我是测试", "address": "123456"})
	})
	r.Run()
}