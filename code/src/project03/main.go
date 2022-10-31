package main

import "project03/temp"

func main(){
	customerView :=temp.CustomerView{Loop: true, Key: " ", CustomerService: temp.NewCustomerService(),}
	customerView.MainMenu()
}
