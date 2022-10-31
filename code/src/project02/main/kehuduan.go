package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main()  {
	//用TCP链接服务器
	conn,err := net.Dial("tcp","192.168.229.1:8888")
	if err != nil {
		fmt.Println("链接失败")
		return
	}
	fmt.Println("链接成功请输入数据：")
	//向终端读取信息并放入缓存
	reader := bufio.NewReader(os.Stdin)
	//从缓存中取出信息
	line,errs :=reader.ReadString('\n')
	if errs!=nil {
		fmt.Println("读取失败")

	}
	//把取出的信息发送给服务器
	n,err0 :=conn.Write([]byte(line))
	if err0 != nil{
		fmt.Printf("发送到服务器失败%v",err0)
	}
	fmt.Printf("服务器已经收到%v字节信息",n)
}
