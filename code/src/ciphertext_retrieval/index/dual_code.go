package index

import (
	"ciphertext_retrieval/util"
	"gonum.org/v1/gonum/mat"
)

// 对偶编码函数
// 输入字符转成对应的0-1构成的二进制向量
// 输入参数 str字符串 m为输出二进制向量的长度
// ==================================================
// 主体逻辑
// 首先初始值为0一个维度的二进制向量
// 遍历字符串的每一个字符
// 将两个相邻的字符映射到二进制向量的某个位置
// ==================================================

// 1 字符串没有进行长度判断 理论中字符串长度应该小于m
// 2 中文字符处理问题 利用rune进行处理是否合理
// 3 参考了开源代码里面的处理方式最后一位是和空格一起进行的hash
// 4 hash是利用的最简单的对%m取模 冲突之后一直找下一个不为0的位置

func DualWordCode(str string, m int) *mat.VecDense {
	runeStr := []rune(str)
	strLen  := len(runeStr) //处理完中文字符长度
	vector  := util.GenVecDense(m,0) //建立一个全为0的向量

	//对每一个字符进行处理
	for i := 0; i < strLen; i++ {
		if i == strLen - 1 { //如果是最后一个字符，将最后一个字符加上空格
			c := int(runeStr[i]) + 32
			c  = c % m //映射到具体的位置 取模
			//fmt.Println("============")
			//fmt.Println(c)
			//查找插入的位置
			for {
				if vector.AtVec(c) == 0  { //找到位置直接插入
					vector.SetVec(c,1)
					break
				} else if c != m - 1 { 		//产生碰撞
					c += 1
				} else {              		//一直到了最后的位置
					c = 0
					vector.SetVec(c,1)
					break
				}
			}
		} else { //处理字符串中[0,m-1)的位置

			c := int(runeStr[i]) + int(runeStr[i+1])
			c = c % m
			//fmt.Println(c)
			for {
				if vector.AtVec(c) == 0  {
					vector.SetVec(c,1)
					break
				} else if c != m - 1 {
					c += 1
				} else {
					c = 0
					vector.SetVec(c,1)
				}
			}
		}
	}
	return vector
}