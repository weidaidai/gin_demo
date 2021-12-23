package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r:=gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"router")
	})
	r.POST("/xpost")
	r.PUT("/xxput")
	r.Run()
}