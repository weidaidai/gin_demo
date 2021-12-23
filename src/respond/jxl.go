package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{"mes": "json"})
	})
	r.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{"mes": "xml"})

	})
	r.GET("/struct", func(c *gin.Context) {

		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200, msg)

		//要注意大小写，首字母要大写
		type User struct {
			Id   int64
			Name string
			Age  uint32
		}
		u := User{Id: 1, Name: "小明", Age: 15}
		c.JSON(200, u)
//{"Name":"root","Message":"message","Number":123}{"Id":1,"Name":"小明","Age":15} 输出内容
	})

	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(200, gin.H{"mes": "yaml"})

	})
	r.GET("/all", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "string"
		//protobuf
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})
r.Run()
}
