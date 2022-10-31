package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

const valueNum int= 5


type Todo struct {
	Id int
	Title string
	Status bool
}

func FindFromId(id string) map[string]string  {
	dsn := "root:123456@tcp(127.0.0.1:3306)/list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err01 := gorm.Open("mysql", dsn)
	if err01 != nil {
		fmt.Println("链接失败")
		panic(err01)
	}
	defer db.Close()
	//db.SingularTable(true)
	db.AutoMigrate(&Todo{})

	var user Todo
	var index int = 0
	var num int = 0
	var information [valueNum]string
	db.Where("Id=?",id).Take(&user)
	Message := user.Title+" "
	for i,value := range Message {
		if value == ' '||value == 0 {
			information[num] = Message[index:i]
			index = i
			num++
		}

	}
	informationmap := map[string]string{
		"Name":information[0],
		"Time":information[1],
		"CompanyId":information[2],
		"Operator":information[3],
	}
	return informationmap
}

func main()  {
	//var  informationOfblockchain [valueNum]string
	informationOfblockchain := FindFromId("6")
	fmt.Println("the result is ",informationOfblockchain)
	//dsn := "root:123456@tcp(127.0.0.1:3306)/list?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err01 := gorm.Open("mysql", dsn)
	//if err01 != nil {
	//	fmt.Println("链接失败")
	//	panic(err01)
	//}
	//defer db.Close()
	////db.SingularTable(true)
	//db.AutoMigrate(&Todo{})
	//
	//var user Todo
	//var index int = 0
	//var num int = 0
	//var information [valueNum]string
	//db.Where("Id=?","2").Take(&user)
	//Message := user.Title+" "
	//for i,value := range Message {
	//	if value == ' '||value == 0 {
	//		information[num] = Message[index:i]
	//		index = i
	//		num++
	//	}
	//
	//	}

//fmt.Println("the result:",information)

	//db.First(&user)
	//fmt.Println("读出数据：",user)
	//user.Name = "Luse"
	//db.Delete(&user)
}
