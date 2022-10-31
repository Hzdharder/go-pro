package util

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestGenRandMatrix(t *testing.T) {
	a := GenRandMatrix(5,6,9)
 	fmt.Println(a)
	fa := mat.Formatted(a, mat.Prefix("    "), mat.FormatPython())
	fmt.Printf("layout syntax:\na = %#v\n\n", fa)
}

func TestGenMatrix(t *testing.T) {
	a := GenMatrix(5,4,1)
	fmt.Println(a)
	fa := mat.Formatted(a, mat.Prefix("    "), mat.FormatPython())
	fmt.Printf("layout syntax:\na = %#v\n\n", fa)
}

func TestGenSymmetricPositiveDefiniteMatrix(t *testing.T)  {
	a := GenSymmetricPositiveDefiniteMatrix(5,5)
	fmt.Println(a)
	fa := mat.Formatted(a, mat.Prefix("    "), mat.FormatPython())
	fmt.Printf("layout syntax:\na = %#v\n\n", fa)
}

func TestGenVecDense(t *testing.T) {
	a := GenVecDense(5,1)
	fmt.Println(a)
}
