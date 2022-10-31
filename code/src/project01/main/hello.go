package main

import (
	"fmt"
)

func main()  {

var lever int = 20
for i:=1;i<=lever;i++{
	for k:=1;k<=lever-i;k++  {
		fmt.Printf(" ")
	}
	for j:=1;j<=2*i-1;j++ {
		if j==1 || j==2*i-1 || i==lever {
			fmt.Printf("*")
		}else {
			fmt.Printf(" ")
		}

	}
	fmt.Printf("\n")
}
}