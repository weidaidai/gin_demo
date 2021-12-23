package main

import "github.com/gin-gonic/gin"

//multipart/form-data格式用于文件上传
//
//gin文件上传与原生的net/http方法类似，不同在于gin把原生的request封装到c.Request中

func main()  {
	r:=gin.Default()
	r.POST("updload", func(c *gin.Context) {
		_,handers,err:=c.Request.FormFile("flie")
		if err != nil {
			return
		}
		//限制文件大小
		if handers.Size>1024*1024*2{
			return

		}
		//获取文件类型
		if handers.Header.Get("Content-Type")!="image/png"{
			return
		}

		c.SaveUploadedFile(handers,"./video/"+handers.Filename)
		c.String(200,handers.Filename)
	})
	r.Run()
}