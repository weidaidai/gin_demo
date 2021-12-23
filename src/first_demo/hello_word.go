package main

import "github.com/gin-gonic/gin"

func main()  {
	//1.创
	r:=gin.Default()
	//2.绑定路由规则，执行函数，gin.context 封装request and response
	r.GET("/", func(context *gin.Context) {
		context.String(200,"hello world")
	})
	//3.监听端口
	r.Run(":8080")
}