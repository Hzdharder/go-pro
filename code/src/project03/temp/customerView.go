package temp

import (
	"fmt"

)

type CustomerView struct {
	Loop bool
	Key string
    CustomerService *CustomerService

}



func (this *CustomerView) list() {
	customers := this.CustomerService.List()
	fmt.Println()
	fmt.Println()
	fmt.Println("******************************客户列表******************************")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for i :=0;i<len(customers) ;i++  {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("********************************************************************")
	fmt.Println()
	fmt.Println()
}
func (this *CustomerView)add()  {
	fmt.Println("**************************添加客户****************************")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	customer := NewCustomer2(name,gender,age,phone,email)
	if this.CustomerService.Add(customer){
		fmt.Println("**************************添加完成****************************")
		fmt.Println()
		fmt.Println()
	}else {
		fmt.Println("**************************添加失败****************************")
		fmt.Println()
		fmt.Println()
	}
}

func (this *CustomerView) delete()  {
	fmt.Println("********************删除客户***********************")
	fmt.Printf("请选择想要删除用户的编号（-1表示退出）: ")
	id := -1
	fmt.Scanln(&id)
	fmt.Println("****************************************************")
	if id == -1 {
		fmt.Println("你已经退出删除客户")
		fmt.Println()
		fmt.Println()
		return
	}
	fmt.Println("确认是否删除（Y/N）:")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "Y"||choice == "y" {
		if this.CustomerService.Delete(id) {
			fmt.Println("********************删除成功***********************")
			fmt.Println()
			fmt.Println()

		}else {
			fmt.Println("********************删除失败***********************")
			fmt.Println()
			fmt.Println()

		}
	}

}
func (this *CustomerView)Alter()  {
	fmt.Println("********************客户信息管理修改***********************")
	fmt.Printf("请选择要修改的用户编号:")
	id := 0
	fmt.Scanln(&id)
	fmt.Println()
	fmt.Printf("请选择要修改的内容:（姓名，性别，年龄，电话，邮箱）")
	choice := ""
	fmt.Scanln(&choice)
	switch choice {
	case "姓名":
		{
			fmt.Println("********************名字修改***********************")
			fmt.Printf("请输入你想改成的名字:")
			name0 := ""
			fmt.Scanln(&name0)
			if this.CustomerService.AlterName(id, name0) {
				fmt.Println("********************修改成功***********************")
				fmt.Println()
				fmt.Println()
			} else {
				fmt.Println("********************修改失败***********************")
				fmt.Println()
				fmt.Println()
			}
		}
	case "性别":
		{
			fmt.Println("********************性别修改***********************")
			fmt.Printf("请输入你想改成的性别:")
			gender0 := ""
			fmt.Scanln(&gender0)
			if this.CustomerService.AlterGender(id, gender0) {
				fmt.Println("********************修改成功***********************")
				fmt.Println()
				fmt.Println()
			} else {
				fmt.Println("********************修改失败***********************")
				fmt.Println()
				fmt.Println()
			}
		}
	case "年龄":
		{
			fmt.Println("********************年龄修改***********************")
			fmt.Printf("请输入你想改成的年龄:")
			age0 := 0
			fmt.Scanln(&age0)
			if this.CustomerService.AlterAge(id, age0) {
				fmt.Println("********************修改成功***********************")
				fmt.Println()
				fmt.Println()
			} else {
				fmt.Println("********************修改失败***********************")
				fmt.Println()
				fmt.Println()
			}
		}
	case "电话":
		{
			fmt.Println("********************电话修改***********************")
			fmt.Printf("请输入你想改成的电话:")
			phone0 := ""
			fmt.Scanln(&phone0)
			if this.CustomerService.AlterPhone(id, phone0) {
				fmt.Println("********************修改成功***********************")
				fmt.Println()
				fmt.Println()
			} else {
				fmt.Println("********************修改失败***********************")
				fmt.Println()
				fmt.Println()
			}
		}
	case "邮箱":
		{
			fmt.Println("********************邮箱修改***********************")
			fmt.Printf("请输入你想改成的邮箱:")
			email0 := ""
			fmt.Scanln(&email0)
			if this.CustomerService.AlterEmail(id, email0) {
				fmt.Println("********************修改成功***********************")
				fmt.Println()
				fmt.Println()
			} else {
				fmt.Println("********************修改失败***********************")
				fmt.Println()
				fmt.Println()
			}
		}
	default:
		fmt.Println("输入错误")
	}

	}


func (this *CustomerView)MainMenu(){
	for   {
		fmt.Println("********************客户信息管理软件***********************")
		fmt.Println("                     1 添加客户                            ")
		fmt.Println("                     2 修改客户                            ")
		fmt.Println("                     3 删除客户                            ")
		fmt.Println("                     4 客户列表                       ")
		fmt.Println("                     5 退   出                            ")
		fmt.Println("***********************************************************")
		fmt.Printf("请选择1--5功能:   ")

		fmt.Scanln(&this.Key)
		switch this.Key {
		case "1":
			this.add()
		case "2":
			this.Alter()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.Loop = false
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
		if !this.Loop {
			break
		}
	}
	fmt.Println("你退出了管理系统")

}

//func main()  {
//	customerView :=customerView{true," ",NewCustomerService()}
//
//	customerView.mainMenu()
//
//}