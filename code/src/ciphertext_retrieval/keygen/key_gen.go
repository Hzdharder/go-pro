// Package keygen Generate key
package keygen

import (
	"ciphertext_retrieval/matrix"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

// Cipher include two matrix M1 M2 and SK (vector)
// M1 M2 The elements of the matrix(M1 M2) consist of 0 or 1
// SK  consist of 0 or 1
type Cipher struct {
	M1       *mat.Dense //matrix
	M2       *mat.Dense //matrix
	SK       *mat.Dense //vector
	BaseUnit int        //BaseUnit*BaseUnit is Dimension of vector
	R        float64    //random number
}

/**
 * @Author peng
 * @Description  Generate key
 * @Date 16:34 2022/6/28
 * @Param Dimension of vector
 * @return Cipher
 **/

// GenCipher based on baseUnit Generate key
func GenCipher(baseUnit int) *Cipher {
	k := baseUnit * baseUnit
	currentTime := time.Now().UnixNano()
	m1 := matrix.GenRandMatrix(k, k, 2, currentTime)
	m2 := matrix.GenRandMatrix(k, k, 2, currentTime+1)
	sk := matrix.GenSkVec(k)
	rand.Seed(currentTime)
	r := rand.Float64()
	return &Cipher{
		M1:       m1,
		M2:       m2,
		SK:       sk,
		BaseUnit: baseUnit,
		R:        r,
	}
}
