package main

import (
	"github.com/gin-gonic/gin"
	"test/myRouter"
)

func main() {
	router := gin.New()

	v1 := router.Group("v1")
	{
		myRouter.BalanceRouter(v1)
	}
	router.Run(":8080")
}
