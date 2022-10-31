package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//type UserInfo struct {
//	ID uint
//	Name string
//	Gender string
//	Birthday time.Time
//	Hobby string
//}
//type User struct {
//gorm.Model
//	Name        string
//	Age           sql.NullInt64
//	Birthday     *time.Time
//	Email        string `gorm:"type:varchar(120);unique_index"`
//	Role string `gorm:"size:255"`
//	MemberNumber *string `gorm:"unique;not null"`
//Num int `gorm:"AUTO_INCREMENT"`
//Address string `gorm:"index:addr"`
//IgnoreMe int`gorm:"-"`
//}
//func main()  {
//	dsn := "root:123456@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err01 := gorm.Open("mysql", dsn)
//	if err01 != nil {
//		fmt.Println("链接失败")
//		panic(err01)
//	}
//	defer db.Close()
//	db.SingularTable(true)
//	db.AutoMigrate(&User{})
//	//u1 := UserInfo{1,"caixukun","man",time.Now(),"dancerap"}
//	//db.Create(&u1)
////	var u UserInfo
////	db.First(&u)//查询第一条记录
////	fmt.Println("记录为：",u)
////	db.Model(&u).Update("hobby","pee")//更新（修改）值
////db.Delete(&u)// s删除一条记录
//
//
//}
type User02 struct {
	ID int64
	Name string
	Age  int64
}

func main02()  {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
		db, err01 := gorm.Open("mysql", dsn)
		if err01 != nil {
			fmt.Println("链接失败")
			panic(err01)
		}
		defer db.Close()
		//db.SingularTable(true)
		db.AutoMigrate(&User02{})
		u := User02{Name:"蔡徐坤",Age:16}
		db.Create(&u)
}