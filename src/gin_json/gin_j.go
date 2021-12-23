package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//定义接收数据的结构体
type Login struct {
	User string `form:"username"json:"user"uri:"user"xml:"user"binding:"required"`
	Password string `form:"password"json:"password"uri:"password" xml:"password"binding:"required"`
}

func main()  {
r:=gin.Default()
//json 绑定
r.POST("login", func(c *gin.Context) {
	var json Login
	if err:=c.ShouldBindJSON(&json);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if json.User!="root"||json.Password!="admin"{
c.JSON(http.StatusBadRequest,gin.H{"staus":"304"})
		return
	}
	c.JSON(200,gin.H{"staus":200})
})

r.Run()

}
// 请求  ---- curl http://127.0.0.1:8080/login -H `content-type:application/json` -d "{\"user\":\"root\",\"password\":\"admin\"}" -X POST
//{"状态码":200}