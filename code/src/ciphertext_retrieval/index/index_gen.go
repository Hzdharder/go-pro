// Package index build index
package index

import (
	"ciphertext_retrieval/dualcode"
	"ciphertext_retrieval/keygen"
	"ciphertext_retrieval/matrix"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// according to cipher Index each document
// filepath File path, numberOfFile Number of documents
// topN Number of keywords
// cipher Key generated in the previous step

/**
 * @Author peng
 * @Description //according to cipher Index each document
 * @Date 16:31 2022/6/28
 * @Param filepath File path, numberOfFile Number of documents
 * @return index
 **/

// BuildIndex BuildIndex for document
func BuildIndex(filePath string, numberOfFile, topN int, cipher *keygen.Cipher) [][]*mat.VecDense {

	// dualWordCodes[article][keywords]
	dualWordCodes := make([]*mat.VecDense, numberOfFile) // []*mat.VecDense为列向量
	// topKeys[article][keywords]
	topKeys := InitialKey(filePath, numberOfFile, topN) // 二维
	for _, value := range topKeys {
		fmt.Println(value)
	}
	for indexPage, topKey := range topKeys { //对每一篇文章的关键词，使用对偶函数
		dualWordCodes[indexPage], _ = dualcode.DualWordCodeSlice(topKey, cipher.BaseUnit)
	}

	// Extract private key
	sk := cipher.SK

	// Build B1 B2
	// Traverse SK and B
	// if s[i] = 1 then b[i] = b1[i] = b2[i]
	// if s[i]=0 then b1[i] = 1/2 * b[i] + r and b2[i] = 1/2 * b[i] - r
	I := make([][]*mat.VecDense, numberOfFile)
	for i := 0; i < numberOfFile; i++ {
		I1 := make([]*mat.VecDense, 2)
		B1 := matrix.GenVecDense(cipher.BaseUnit*cipher.BaseUnit, 0)
		B2 := matrix.GenVecDense(cipher.BaseUnit*cipher.BaseUnit, 0)
		for j := 0; j < cipher.BaseUnit*cipher.BaseUnit; j++ {
			if sk.At(j, 0) == 1 {
				B1.SetVec(j, dualWordCodes[i].AtVec(j))
				B2.SetVec(j, dualWordCodes[i].AtVec(j))
			} else {
				B1.SetVec(j, 0.5*float64(dualWordCodes[i].AtVec(j))+cipher.R)
				B2.SetVec(j, 0.5*float64(dualWordCodes[i].AtVec(j))-cipher.R)
			}
		}
		B1.MulVec(cipher.M1.T(), B1)
		B2.MulVec(cipher.M2.T(), B2)
		I1[0] = B1
		I1[1] = B2
		I[i] = I1
	}
	return I
}
