package index

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestDualWordCode(t *testing.T) {
	a := DualWordCode("peng",20)
	fmt.Println(a)
	fa := mat.Formatted(a, mat.Prefix("    "), mat.FormatPython())
	fmt.Printf("layout syntax:\na = %#v\n\n", fa)
}
