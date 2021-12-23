package main

import "github.com/gin-gonic/gin"

func main()  {
	r:=gin.Default()
	r.GET("/",hander)
	if err:=r.Run();
	 err != nil {

	}
	//r.Run()

}
func hander(c *gin.Context)  {
	c.JSON(200,gin.H{
		"json":"val",


	})

}