package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func maind()  {

	file := "src/go_code/project02/temp/cmc部署结果.txt"
	File,err := os.Open(file)
	defer  File.Close()
	if err != nil {
		fmt.Println("there must be some wrong here")
	}

	reader := bufio.NewReader(File)
	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
			fmt.Println(str)

	}

	/*
    file := "src/go_code/project02/temp/001"
    contain,err :=ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("there must be some wrong here")
	}else {
		fmt.Println(string(contain))
	}


	file := "src/go_code/project02/temp/001"
	File,err := os.OpenFile(file,os.O_RDWR|os.O_APPEND,0666)
	defer File.Close()
	if err != nil {
		fmt.Println("there must be some wrong here")
	}
	str := "这是新加得东西\n"
	write := bufio.NewWriter(File)
	write.WriteString(str)
	write.Flush()
	*/
}
