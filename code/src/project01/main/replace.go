package main

import (
	"fmt"
	"strings"
)

func main()  {
	str1:= "go go go go to beijing"
	Nstr := strings.Replace(str1,"go","not go",-1)
	fmt.Println(Nstr)

	str2 := " nnj njn njn jn jn "
	Nstr2 := strings.TrimSpace(str2)
	fmt.Println(Nstr2)

	str3 := "A nameA"
	Nstr3 := strings.TrimLeft(str3,"A")
	fmt.Println(Nstr3)

}
