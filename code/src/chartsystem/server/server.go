package main

import (
	"chartsystem/package/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

func serverProcessMes(conn net.Conn,mes message.Message)  {
	
}
func readPkg(conn net.Conn)(mes message.Message,err error)  {
	buf := make([]byte,1024*8)
	_,err0 := conn.Read(buf[:4])
	if err0 != nil{
		//err0 = errors.New("接收长度错误")
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	n,err11 := conn.Read(buf[:pkgLen])
	if n != int(pkgLen)||err11 != nil {

		err0 = errors.New("接收数据失败")
		return
	}
	err12 := json.Unmarshal(buf[:pkgLen],&mes)
	if err12 != nil {
		fmt.Println("反序列化失败err12",err12)
		return
	}
	return
}
func process(conn net.Conn)  {
	defer conn.Close()

	//fmt.Println("等待读取客户端发送的数据")
	for  {
      mes,err12 := readPkg(conn)
		if err12 == io.EOF {
			fmt.Println("客户端断链")
			return
		}else {
			fmt.Println("readPkg wrong",err12)
			return
		}
		fmt.Println("接收到数据 ：",mes)

	}


}

func main()  {
	listen,err01 :=net.Listen("tcp","0.0.0.0:8889")
	if err01 != nil{
		fmt.Println("there must be something wrong err01 ",err01)
		return
	}

    defer listen.Close()
	for {
		fmt.Println("等待读取客户端发送的数据")
		coon,err02 := listen.Accept()
		if err02 != nil {
				fmt.Println("客户端断链")
				return

		}

		go process(coon)


	}
}
