package util

import (
	"fmt"
	"testing"
)

func TestNewBloomFilter(t *testing.T) {
	bf := NewBloomFilter(0.0001, 100) //失误率 和 插入元素个数
	fmt.Println(bf)
	bf.Set([]byte("peng"))
	bf.Set([]byte("network"))
	bf.Set([]byte("pen"))
	bf.Set([]byte("networj"))

	fmt.Println(bf.Check([]byte("pen")))
	fmt.Println(bf.Check([]byte("networj")))
	fmt.Println(bf.Check([]byte("networl")))
	fmt.Println(bf)

	floatSlice := make([]float64, 100)
	for i :=0; i < 100 ; i++ {
		floatSlice[i] = 0.1 * float64(i) + float64(i)
	}
	bf.SetFloatVector(floatSlice)

	for i :=0; i < 100 ; i++ {
		fmt.Println(bf.CheckFloat64(0.1 * float64(i) + float64(i)))
	}
	fmt.Println(bf.CheckFloat64(0.1 * float64(1) + float64(2)))
}