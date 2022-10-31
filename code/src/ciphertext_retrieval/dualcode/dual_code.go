// Package dualcode dual function
package dualcode

import (
	"ciphertext_retrieval/matrix"
	"strings"

	"gonum.org/v1/gonum/mat"
)

// dual function
// input string
// output vector Composed of 0 and 1
// ==================================================
// Main processes
// generated vector is composed of 0
// traverse every character of a word
// adjacent characters are mapped to a position in vector
// ==================================================

// ASCIIA ASCII of a
const ASCIIA = 97

/**
 * @Author peng
 * @Description dual function
 * @Date 16:30 2022/6/28
 * @Param  string,len of vector m
 * @return vector
 **/

// DualWordCode dual function
func DualWordCode(str string, m int) (vector *mat.VecDense, index []int) {
	// ignore case
	str = strings.ToLower(str)
	strLen := len(str)
	vector = matrix.GenVecDense(m * m, 0)
	index = make([]int, strLen)

	//对每一个字符进行处理
	for i := 0; i < strLen; i++ {
		if i == strLen - 1 {
			c := int(str[i] - ASCIIA) * m + int(str[0] - ASCIIA)
			index[i] = c
			vector.SetVec(c, 1)
		} else {
			c := int(str[i] - ASCIIA) * m + int(str[i + 1] - ASCIIA)
			index[i] = c
			vector.SetVec(c, 1)
		}
	}
	return vector, index
}

// DualWordCodeSlice dual function process string slice
func DualWordCodeSlice(strs []string, m int) (vector *mat.VecDense, indexs [][]int) {
	vector = matrix.GenVecDense(m*m, 0)
	indexs = make([][]int, len(strs))
	for index, str := range strs {
		str = strings.ToLower(str)
		strLen := len(str)
		if indexs[index] == nil {
			indexs[index] = make([]int, strLen)
		}
		for i := 0; i < strLen; i++ {
			if i == strLen - 1 {
				c := int(str[i] - ASCIIA) * m + int(str[0] - ASCIIA)
				c = c % (m*m)
				indexs[index][i] = c
				vector.SetVec(c, 1)
			} else {
				c := int(str[i] - ASCIIA) * m + int(str[i + 1] - ASCIIA)
				c = c % (m * m)
				indexs[index][i] = c
				vector.SetVec(c, 1)
			}
		}
	}
	return vector, indexs
}
