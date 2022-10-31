package main

import (
	"fmt"
)

func Inchain(inchain chan  int)  {
	for i := 0;i <50 ;i ++  {
		inchain<- i
		fmt.Println("读入: ",i)
		//time.Sleep(time.Second)
	}
	close(inchain)
}
func JudgeOver(inchain chan int,BoolChain chan bool)  {
	for {
		v,ok := <-inchain
		fmt.Println("读出: ",v)
		//time.Sleep(time.Second)
		if !ok {
			break
		}
		}
		BoolChain<- true
		close(BoolChain)
		fmt.Println("现在已经读出了50个数据")
}

func mainc(){
	inchain := make(chan int, 50)
	boolchain := make(chan bool,1)
	go Inchain(inchain)
	go JudgeOver(inchain,boolchain)
	for{
		_,ok := <-boolchain
		if !ok {
			break
		}
	}
}


