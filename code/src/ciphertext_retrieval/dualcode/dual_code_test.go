// Package dualcode test package dual
package dualcode

import (
	"fmt"
	"testing"
)

// TestDualWordCode Test word coding
func TestDualWordCode(t *testing.T) {
	a, index := DualWordCode("network", 26)
	fmt.Println(a)
	fmt.Println(index)
}

// TestDualWordCodeSlice Test word slice coding
func TestDualWordCodeSlice(t *testing.T) {
	a, index := DualWordCode("network", 26)
	fmt.Println(a)
	fmt.Println(index)
	strs := make([]string, 2)
	strs[0] = "network"
	strs[1] = "netword"
	vector, indexs := DualWordCodeSlice(strs, 26)
	fmt.Println(vector)
	fmt.Println(indexs)

}