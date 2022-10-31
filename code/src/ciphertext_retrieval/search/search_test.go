// Package search test
package search

import (
	"ciphertext_retrieval/index"
	"ciphertext_retrieval/keygen"
	"ciphertext_retrieval/trapdoor"
	"fmt"
	"testing"
	"time"
)

func TestSearch(t *testing.T) {
	key := keygen.GenCipher(5)
	I := index.BuildIndex("../testdata/search_test/", 32, 10, key)
	strs := make([]string, 2)
	strs[0] = "you"
	strs[1] = "no"
	T := trapdoor.GenTrapdoor(key, strs)
	start := time.Now()
	for i := 0; i < 100; i++ {
		Search(I, T)
	}
	//some func or operation
	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
