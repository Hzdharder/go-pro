package myRouter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test/handlers"
)

func BalanceRouter(v1 *gin.RouterGroup) {
	v1.GET("/available/balance", handlers.OnGetAvailableBalance)
	v1.GET("/foo%2fbar", func(context *gin.Context) {
		fmt.Printf("fullPath:%v, url:%v", context.FullPath(), context.Request.URL.Path)
	})
}
