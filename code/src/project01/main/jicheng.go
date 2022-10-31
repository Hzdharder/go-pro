package main

import "fmt"

type Student struct {
	Name string
	Num  string
}

type Colleage struct {
	Student
	Score int
}

type Child struct {
	Student
	Score int
}


func main()  {
	Col := new(Colleage)
	Col.Name = "Tom"
	Col.Num = "2022010539"
	Col.Score = 98
	fmt.Println("the first student message is :",*Col)
	Chi := Child{Student{"Jack","20225600"},88}
	fmt.Println("the scoend student message is :",Chi)
    Col2 := Colleage{}
    Col2.Name ="xiao"
    Col2.Num = "2022154554"
    Col2.Score = 21
	fmt.Println("the third student message is :",Col2)
    var Chi2 Child
    Chi2.Name = "xiaomi"
	fmt.Println("the forth student message is :",Chi2)
}
