package main

import (
	"errors"
	"fmt"
	"google.golang.org/protobuf/proto"
	"grpc/service"

	//"pkg/mod/google.golang.org/protobuf@v1.28.1/proto"

	//"protobuf-master/proto"
)

func main()  {

	user := &service.User{

		Username:"Jack",
		Age:13,
	}
     marnum,err := proto.Marshal(user)
     if err != nil {
     	err = errors.New("序列化错误")
		 return
	 }
   newuser := &service.User{}
   err01 := proto.Unmarshal(marnum,newuser)
	if err01 != nil {
		fmt.Println("反序列化失败：",err01)
		return
	}
	fmt.Println("最后结果为：",newuser.String())
}
