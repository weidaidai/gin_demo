package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//URL参数可以通过DefaultQuery()或Query()方法获取

func main()  {
	r:=gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//指定默认值
		//8080/user
		name:=c.DefaultQuery("name","weidaidai")
		c.String(http.StatusOK,fmt.Sprintf("hello%s",name))
	})
	r.Run(":8090")
}

//http://127.0.0.1:8090/user
//hello weidaidai