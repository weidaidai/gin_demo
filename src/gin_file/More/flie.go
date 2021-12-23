package main

import "github.com/gin-gonic/gin"

func main()  {
	r:=gin.Default()
	//限制大小
	r.MaxMultipartMemory=8<<20
	r.POST("/upload", func(c *gin.Context) {
		c.MultipartForm()
	})

}
