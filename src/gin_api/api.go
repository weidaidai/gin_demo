package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
//可以通过Context的Param方法来获取API参数
func main()  {
	r:=gin.Default()
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name:=c.Param("name")
		action:=c.Param("action")
		//切片
		action=strings.Trim(action,"/")
		c.String(http.StatusOK,name+" is "+action)

	})
	r.Run()
}
//http://127.0.0.1:8080/user/开心/weidaidai
//输出 开心 is weidaidai