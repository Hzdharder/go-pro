package main

import  (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"message":"Hello welcome here",
	})
}
func main()  {
	r := gin.Default()
	r.GET("/myweb01",sayHello)
	r.GET("myweb02", func(c *gin.Context) {
		data := map[string]interface{}{
			"姓名":"小王子",
			"性别":"男",
			"年龄":12,
		}
		c.JSON(http.StatusOK,data)
	})
	r.Run(":9090")
}
