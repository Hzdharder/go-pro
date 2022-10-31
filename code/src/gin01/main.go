package main

import (

	"github.com/gin-gonic/gin"
	"net/http"

	"path"

)

func main()  {
	t := gin.Default()
	t.LoadHTMLFiles("./login.html","./index.html","./upload.html")


	t.GET("/myweb", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",nil)
	})
	//获取form表单提交的数据
	t.POST("/myweb", func(c *gin.Context) {
		Username := c.PostForm("username")
		Password := c.PostForm("password")
		c.HTML(http.StatusOK,"index.html",gin.H{
			"Username":Username,
			"Password":Password,
		})
	})






	t.GET("/myweb/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK,gin.H{
			"year":year,
			"month":month,
		})
	})








	t.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK,"upload.html",nil)
	})
	t.POST("/upload", func(c *gin.Context) {
		file,err01 := c.FormFile("f1")
		if err01 != nil {
			c.JSON(http.StatusOK,gin.H{"wrong err01":"文件获取失败"},)
		}
		filePath := path.Join("./",file.Filename)
		c.SaveUploadedFile(file,filePath)
		c.JSON(http.StatusOK,gin.H{"message":"上传成功"})

		})



	//重定向
	t.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,"https://cf.qq.com/cp/a20210707week/index.html")
	})
	t.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"
		t.HandleContext(c)
	})
	t.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"message":"转接成功"})
	})



   t.Run(":9090")

}
