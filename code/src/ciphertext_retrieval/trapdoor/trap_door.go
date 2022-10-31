// Package trapdoor is construct
package trapdoor

import (
	"ciphertext_retrieval/dualcode"
	"ciphertext_retrieval/keygen"
	"ciphertext_retrieval/matrix"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

// Build trap door

/**
 * @Author peng
 * @Description // Build trap door
 * @Date 16:37 2022/6/28
 * @Param  cipher Query string
 * @return slice of VecDense
 **/

// GenTrapdoor Generate trap door
func GenTrapdoor(cipher *keygen.Cipher, query []string) []*mat.VecDense {
	I := make([]*mat.VecDense, 2)

	dualWordCode, _ := dualcode.DualWordCodeSlice(query, cipher.BaseUnit)

	sk := cipher.SK
	rand.Seed(time.Now().UnixNano() + 3)
	r := rand.Float64()
	dimension := cipher.BaseUnit * cipher.BaseUnit
	B1 := matrix.GenVecDense(dimension, 0)
	B2 := matrix.GenVecDense(dimension, 0)
	for j := 0; j < dimension; j++ {
		if sk.At(j, 0) == 1 {
			B1.SetVec(j, dualWordCode.AtVec(j))
			B2.SetVec(j, dualWordCode.AtVec(j))
		} else {
			B1.SetVec(j, 0.5 * float64(dualWordCode.AtVec(j)) + r)
			B2.SetVec(j, 0.5 * float64(dualWordCode.AtVec(j)) - r)
		}
	}

	I1 := mat.NewVecDense(dimension, nil)
	I2 := mat.NewVecDense(dimension, nil)

	tempMatrixM1 := mat.NewDense(dimension, dimension, nil)
	tempMatrixM2 := mat.NewDense(dimension, dimension, nil)
	tempMatrixM1.Inverse(cipher.M1)
	tempMatrixM2.Inverse(cipher.M2)
	I1.MulVec(tempMatrixM1, B1)
	I2.MulVec(tempMatrixM2, B2)
	I[0] = I1
	I[1] = I2
	return I
}
