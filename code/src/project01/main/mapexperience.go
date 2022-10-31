package main

import "fmt"

func main()  {
	var num map[string]string
	num = make(map[string]string,10)
	num["地方02"]="上海"
	num["地方03"]="北京"
	num["地方04"]="深圳"
	num["地方05"]="杭州"
	num["地方06"]="西安"
	fmt.Println("数据集为：",num)
}
