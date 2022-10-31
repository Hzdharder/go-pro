package temp

import "fmt"

type Customers struct {
	//Acount string
	//Paw    string
	//Message string
	//Num     int
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

func NewCustomer(id int,name string,gender string,age int,phone string,email string)Customers  {
	return Customers{id,name,gender,age,phone,email}
}


func NewCustomer2(name string,gender string,age int,phone string,email string)Customers  {
	return Customers{Name:name,Gender:gender,Age:age,Phone:phone,Email:email}
}


func (this *Customers)GetInfo() string  {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t\t",this.Id,this.Name,this.Gender,this.Age,this.Phone,this.Email)
	return  info
}
//func (C Client)AddClient(acount string,paw string,message string)  {
//	var c []Client
//	c =make([]Client,0)
//	num := len(c) + 1
//	c[num] =Client{acount,paw,message,num}
//
//}
//func (C Client)AdjustClient()  {
//
//}