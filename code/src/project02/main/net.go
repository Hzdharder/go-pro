package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn)  {
	defer conn.Close()

	for {
		buf := make([]byte,1024)
		fmt.Printf("服务器正在等待客户端%v发送信息",conn.RemoteAddr().String())
		n,err := conn.Read(buf)
		if err!= nil{
			fmt.Print("客户端出错：%V",err)
			return
		}
		fmt.Printf("收到客户端内容：%v",string(buf[:n]))
	}

	}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("监听失败")
		return
	}
	defer listen.Close()
	for   {
		fmt.Println("等待客户链接")
		//拿到conn链接通道
		conn,err :=listen.Accept()
		if err!=nil {
			fmt.Println("链接失败")
		}else {
			fmt.Printf("客户%v链接成功",conn.RemoteAddr().String())
		}
		go process(conn)
	}

}
