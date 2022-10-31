package main

import (
	"chartsystem/package/login"
	"fmt"
	"os"
)
var userId int
var userPwd string
func main()  {
	loop := true
    key := 0
	for loop  {
       fmt.Println("********************Welcome to the multi chart room*********************")
       fmt.Println("***********************1 login the chart room***************************")
       fmt.Println("***********************2 register the chart room************************")
       fmt.Println("***********************3 quite the system*******************************")
       fmt.Println("***********************Please select your choose************************")

       fmt.Scanf("%d\n",&key)
		switch key {
		case 1 :
			fmt.Println("1")
			loop =false
		case 2 :
			fmt.Println("2")
			loop =false
		case 3 :
			fmt.Println("you have exit system")
			os.Exit(0)
		default:

			fmt.Println("there must be something wrong please putin again")

		}
	}
	if key == 1 {
		fmt.Println("Please input your userId")
		fmt.Scanf("%d\n",&userId)
		fmt.Println("Please input your userPwd")
		fmt.Scanf("%s\n",&userPwd)
		err := login.Login(userId,userPwd)
		if err != nil{
			fmt.Println("login false :err ",err)
		}else{
			fmt.Println("login success")
		}
	}else if key == 2{
       fmt.Println("please register ")
	}

}
