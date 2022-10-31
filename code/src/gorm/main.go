package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Stu struct {
//	gorm.Model
	ID int64
	Name string `gorm:"default:'Jack'"`
	Age  int64
	hobby string
}

func main0()  {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err01 := gorm.Open("mysql", dsn)
	if err01 != nil {
		fmt.Println("链接失败")
		panic(err01)
	}
	defer db.Close()
	//db.SingularTable(true)
	db.AutoMigrate(&Stu{})
	u := Stu{Name:"Markle01",Age:89,hobby:"Rapdance02"}
	db.Create(&u)
	var user Stu
	db.Where("name=?","Markle01").Delete(&user)
	//db.First(&user)
	//fmt.Println("读出数据：",user)
	//user.Name = "Luse"
	//db.Delete(&user)
}
