package login

import (
	"chartsystem/package/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func Login(userId int,userPwd string)(err error)  {
     //fmt.Printf("your userId is %d,userPwd is %s",userId,userPwd)
     //return
     conn,err03 :=net.Dial("tcp","localhost:8889")
     if err03 != nil{
      fmt.Println(" connect false err03 ",err03)
      return
     }
    defer conn.Close()
     var mes message.Message
     mes.Type = message.LoginMesType
     var loginMes message.LoginMes
     loginMes.UserId = userId
     loginMes.UserPwd = userPwd

     data,err := json.Marshal(loginMes)
     if err != nil{
     	fmt.Println("marshal false",err)
	 }
	 mes.Data = string(data)
	data,err08 := json.Marshal(mes)
	if err08 != nil{
		fmt.Println("marshal false",err)
	}
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)
   n,err09 := conn.Write(buf[0:4])
   if n != 4{
   	fmt.Println("丢失数据长度")
   }
   if err09 != nil{
   	fmt.Println("传输数据错误err09",err09)
	}
	fmt.Println("传输字节数",len(data))
	_,err10 := conn.Write(data)
	if err10 != nil {
		fmt.Println("传输数据失败",err10)
		return
	}
       return
}







