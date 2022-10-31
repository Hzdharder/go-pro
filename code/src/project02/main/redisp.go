package main

import (
	"fmt"
	"redigo/redis"
)

func main()  {
	conn,err :=redis.Dial("tcp","127.0.0.1:6379")
	if err!=nil {
		fmt.Println("there must be wrong ",err)
		return
	}
	defer conn.Close()


	_,err01 :=conn.Do("set","name01","杨国强")
	if err01 != nil{
		fmt.Println("write database false")
	}
	name,err02 := redis.String(conn.Do("Get","name01"))
	if err02!=nil {
		fmt.Println("read database false")
	}
	fmt.Println("the name is :",name)



	_,err03 :=conn.Do("HSet","user1","name01","杨国强")
	if err03 != nil{
		fmt.Println("write database false")
	}

	_,err05 :=conn.Do("HSet","user1","sex","男")
	if err05 != nil{
		fmt.Println("write database false")
	}

	_,err06 :=conn.Do("HSet","user1","age",18)
	if err06 != nil{
		fmt.Println("write database false")
	}


	contain,err04 := redis.Strings(conn.Do("HGetAll","user1" ))
	if err04!=nil {
		fmt.Println("read database false")
	}
	fmt.Println("the result of contain :",contain)

}
