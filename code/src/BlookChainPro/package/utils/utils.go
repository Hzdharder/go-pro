package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

//
func IntTOHex(num int64) []byte {
	buff:=new(bytes.Buffer)//开辟内存，存储字节
	err:=binary.Write(buff,binary.BigEndian,num)//num转化字节集写入
	if err!=nil{
		log.Panic(err)
	}
	 return buff.Bytes()//返回字节集

}
