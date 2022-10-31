package main

import (
	"fmt"
	"strings"
)

func Houzhui(houzhui string) func(string)string  {
	return func(name string) string {
		if !strings.HasSuffix(name,houzhui) {
			return name+houzhui
		}
		return name
	}

}
func main()  {
	version := Houzhui(".zip")
	Name1 :=version("weijin.zip")
	fmt.Println("转化文件为： ",Name1)
	Name2 :=version("indx")
	fmt.Println("转化文件为： ",Name2)


//var a  map[int]int
//a = make(map[int]int,5)
//c := make(map[int]int)
//k := map[int]int{
//6 : 5,
//}
}