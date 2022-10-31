package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func main00()  {
	dsn := "root:123456@tcp(127.0.0.1:3306)/list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("链接失败")
	}
	defer db.Close()
	db.AutoMigrate(&Todo{})
	db.Where("Id=?","3").Take(&Todo{})
	fmt.Println(Todo{})
}
