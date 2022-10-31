// Package index test
package index

import (
	"ciphertext_retrieval/keygen"
	"fmt"
	"testing"
	"time"
)

// TestBuildIndex test Index construction input filepath, Cipher
func TestBuildIndex(t *testing.T) {
	cipher := keygen.GenCipher(5)
	start := time.Now()
	for i := 0; i < 100; i++ {
		BuildIndex("../testdata/test_2/", 1, 20, cipher)
	}
	//some func or operation
	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
	// for _, value := range I {
	// 	fmt.Println(value[0])
	// 	fmt.Println(value[1])
	// 	fmt.Println("============================================")
	// }
}

// // TestMatrix test Matrix
// func TestMatix(t *testing.T) {
// 	ciper := keygen.GenCipher(3)
// 	B1 := matrix.GenVecDense(9, 0)
// 	B2 := matrix.GenVecDense(9, 0)
// 	B1.SetVec(2, 1)
// 	B1.SetVec(4, 1)
// 	B1.SetVec(6, 1)
// 	B1.SetVec(8, 1)
// 	B2.SetVec(1, 1)
// 	B2.SetVec(3, 1)
// 	B2.SetVec(5, 1)
// 	B2.SetVec(7, 1)
// 	fmt.Println(ciper.M1)
// 	fmt.Println(ciper.M1.T())
// 	fa := mat.Formatted(ciper.M1.T(), mat.Prefix("    "), mat.FormatPython())
// 	fmt.Printf("layout syntax:\nM1T = %#v\n\n", fa)
// 	B1.MulVec(ciper.M1.T(), B1)
// 	B2.MulVec(ciper.M2.T(), B2)
// 	fa1 := mat.Formatted(ciper.M1, mat.Prefix("    "), mat.FormatPython())
// 	fmt.Printf("layout syntax:\nM1 = %#v\n\n", fa1)
// }
