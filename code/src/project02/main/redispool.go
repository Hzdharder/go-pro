package main

import (
	"fmt"
	"redigo/redis"
)

var pool *redis.Pool

func init()  {
	pool = &redis.Pool{
		Dial: func() ( redis.Conn,  error) {
			return redis.Dial("tcp","localhost:6379")
		},
		DialContext:     nil,
		TestOnBorrow:    nil,
		MaxIdle:         8,
		MaxActive:       0,
		IdleTimeout:     100,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}

func main()  {
	conn := pool.Get()
	defer conn.Close()
	_,err :=conn.Do("hmset","user1","name","迪迦奥特曼","gender","男","age",56)
	if err != nil {
		fmt.Println("there must be something wrong",err)
		return
	}
	contain1,err1 := redis.Strings(conn.Do("hgetall","user1"))
	if err1 != nil {
		fmt.Println("there must be something wrong",err1)
	}else {
		fmt.Println("the contain is :",contain1)
	}
	_,err2 :=conn.Do("set","address","上海")
	if err2 != nil {
		fmt.Println("there must be something wrong",err)
		return
	}
	contain2,err3 := redis.String(conn.Do("get","address"))
	if err3 != nil {
		fmt.Println("there must be something wrong err3  ",err3)
	}else {
		fmt.Println("the next contain is ",contain2)
	}

}
