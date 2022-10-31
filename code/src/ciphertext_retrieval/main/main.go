// Package main demo for search
package main

import (
	"ciphertext_retrieval/index"
	"ciphertext_retrieval/keygen"
	"ciphertext_retrieval/search"
	"ciphertext_retrieval/trapdoor"
	"fmt"
)

func main() {
	// Generate Cipher
	ciper := keygen.GenCipher(26)

	// Build Index
	// 文件地址，文件总数，关键词个数，密码
	I := index.BuildIndex("testdata/", 5, 10, ciper)
	fmt.Println(I)

	// Generate Trapdoor
	// query you, no
	strs := make([]string, 2)
	strs[0] = "you"
	strs[1] = "no"
	T := trapdoor.GenTrapdoor(ciper, strs)
	// query
	search.Search(I, T)
}
