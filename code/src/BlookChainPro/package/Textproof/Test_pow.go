package Textproof

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

func mainx()  {
	start:=time.Now()
	for i:=0;i<1000000000;i++{
		data:=sha256.Sum256([]byte(strconv.Itoa(i)))
		fmt.Printf("%10d,%x\n",i,data)
		fmt.Printf("%s\n",string(data[len(data)-2:]))
		if string(data[len(data)-2:])=="00"{
			usedtime:=time.Since(start)
			fmt.Printf("你以获得区块奖励\n 挖矿时间耗费:%d 纳秒",usedtime)
			break
		}
	}

}