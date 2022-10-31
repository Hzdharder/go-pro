package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"

	//"net/http"
)
var (
	DB *gorm.DB
)

type Todo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func initMySQL()(err error)  {
	dsn := "root:123456@tcp(127.0.0.1:3306)/list?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("链接失败")

		return
	}
return DB.DB().Ping()
}
func main()  {
	err :=initMySQL()
	if err != nil{
		panic(err)
	}
	defer DB.Close()
	DB.AutoMigrate(&Todo{})

    r := gin.Default()
    r.Static("/static","static")
    r.LoadHTMLFiles("templates/index.html" )
    r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

    V1Group := r.Group("v1")
    {
    	//添加待办事务
		V1Group.POST("/todo", func(c *gin.Context) {
			//点击提交请求会发送来这个模块
			//1从请求中拿出数据
			var todo Todo

			c.BindJSON(&todo)



			//2把数据存入数据库

			if err = DB.Create(&todo).Error;err != nil{
				c.JSON(http.StatusOK,gin.H{"error" : err.Error()})
			}else {
				c.JSON(http.StatusOK,todo)
			}
			//3反应请求
		})
    	//查看所有待办事务
		V1Group.GET("todo", func(c *gin.Context) {
			var  todoList  []Todo
			if err = DB.Find(&todoList).Error;err != nil{
				c.JSON(http.StatusOK,gin.H{"error" : err.Error()})
			}else {
				c.JSON(http.StatusOK,todoList)
			}
		})
    	//查看某个代办事务
		V1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//修改一个代办事务
		V1Group.PUT("todo/:id", func(c *gin.Context) {
            id,err := c.Params.Get("id")
            if !err{
				c.JSON(http.StatusOK,gin.H{"error" : "无效id"})
				return
			}
			var  todo Todo
            if err01 := DB.Where("id=?",id).First(&todo).Error;err01 != nil{
				c.JSON(http.StatusOK,gin.H{"error" : err01.Error()})
				return
			}
			c.BindJSON(&todo)
            if err02 := DB.Save(&todo).Error;err02 != nil{
				c.JSON(http.StatusOK,gin.H{"error" : err02.Error()})
				return
			}else {
				c.JSON(http.StatusOK,todo)
			}
		})
		//删除一个代办事务
		V1Group.DELETE("todo/:id", func(c *gin.Context) {
			id,err := c.Params.Get("id")
			if !err{
				c.JSON(http.StatusOK,gin.H{"error" : "无效id"})
				return
			}
			if err01 := DB.Where("id=?",id).Delete(Todo{}).Error;err01 != nil{
				c.JSON(http.StatusOK,gin.H{"error" : err01.Error()})
				return
			}else {
				c.JSON(http.StatusOK,gin.H{"message":"删除完成"})
			}
		})
	}
	r.Run()
}



