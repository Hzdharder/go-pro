package index

import (
	"fmt"
	"testing"
)

func TestInitialKey(t *testing.T) {
	frequentWordSet := InitialKey("../testdata/", 5, 10)
	for index, value := range frequentWordSet {
		fmt.Print(index)
		fmt.Println(value)
	}
}
