package trapdoor

import (
	"ciphertext_retrieval/keygen"
	"fmt"
	"testing"
	"time"
)

func TestGenTrapdoor(t *testing.T) {
	strs := make([]string, 20)
	strs[0] = "and"
	strs[1] = "he"
	strs[2] = "as"
	strs[3] = "no"
	strs[4] = "you"
	strs[5] = "am"
	strs[6] = "it"
	strs[7] = "in"
	strs[8] = "that"
	strs[9] = "on"
	strs[10] = "the"
	strs[11] = "to"
	strs[12] = "black"
	strs[13] = "blockchain"
	strs[14] = "good"
	strs[15] = "IOT"
	strs[16] = "white"
	strs[17] = "red"
	strs[18] = "green"
	strs[19] = "rabbit"
	cipher := keygen.GenCipher(5)
	start := time.Now()
	for i := 0; i < 100; i++ {
		GenTrapdoor(cipher, strs)
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
	// fmt.Println(I[0])
	// fmt.Println(I[1])
}
