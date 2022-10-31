package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	 rand.Seed(time.Now().Unix())
	 start := time.Now().Unix()
	for  {
		a := rand.Intn(100)
		fmt.Println(a)
		time.Sleep(time.Second*2)
		if a==56 {
			end := time.Now().Unix()
			need := end - start
			fmt.Println("the speciall num is got,use time: ",need,"s")
			break
		}
	}

}
