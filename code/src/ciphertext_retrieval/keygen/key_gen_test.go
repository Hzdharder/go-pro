package keygen

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestCipher_GenKey(t *testing.T) {
	a := GenCipher(4)
	fa := mat.Formatted(a.M1, mat.Prefix("    "), mat.FormatPython())
	fmt.Printf("layout syntax:\nM1 = %#v\n\n", fa)
	fb := mat.Formatted(a.M2, mat.Prefix("    "), mat.FormatPython())
	fmt.Printf("layout syntax:\nM2 = %#v\n\n", fb)
	fc := mat.Formatted(a.SK, mat.Prefix("    "), mat.FormatPython())
	fmt.Printf("layout syntax:\nSK = %#v\n\n", fc)
}
