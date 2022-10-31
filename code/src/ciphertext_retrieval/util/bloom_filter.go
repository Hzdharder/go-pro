package util

import (
	"encoding/binary"
	"math"
)

type BloomFilter struct {
	MLength   int64   // 过滤器长度
	MArr      []int64 // 过滤器数组
	NCount    int64   // 插入元素个数
	FalseRate float64 // 误报率
	KHash     int     // hash函数个数
}

//数学公式
// n:集合中元素的个数
// k:使用的hash函数的个数
// m: bit数组的宽度（bit数）
// 最优的哈希函数的个数k = ln2*(数组大小/元素个数) k = ln2 * m/n
// m = - n log p / (log2 )^2
// 确定的值 n 是确定的 m是确定的 误报率是确定的
// num_slices = int(math.ceil(math.log(1.0 / error_rate, 2)))
// bits_per_slice = int(math.ceil(
//     (capacity * abs(math.log(error_rate))) /
//     (num_slices * (math.log(2) ** 2))))
// num_slices = 4  # 哈希函数4个
// bits_per_slice = 32

// todo 1 是否符合论文预期 布隆过滤器容量是否有要求 现在控制了误差率和插入元素的值
//      2 数组可以优化成bit 节约内存

func NewBloomFilter(falseRate float64, size int64) *BloomFilter {
	m, k := getFilterParam(falseRate, size)

	return &BloomFilter{
		MLength:   128,
		MArr:      make([]int64, m),
		NCount:    size,
		FalseRate: falseRate,
		KHash:     k,
	}
}

func getFilterParam(falseRate float64, size int64) (int64, int) {

	m := -1 * math.Log(falseRate) * float64(size) / (math.Ln2 * math.Ln2)
	k := m * math.Ln2 / float64(size)
	return int64(m), int(k)
}

func (bf *BloomFilter) Set(data []byte) {
	for i := 0; i < bf.KHash; i++ {
		index := bf.hashFun(data, i)
		bf.MArr[index] = 1
	}
}

func (bf *BloomFilter) SetFloatVector(vector []float64) {
	for _, value := range vector {
		for i := 0; i < bf.KHash; i++ {
			index := bf.hashFun(Float64ToByte(value), i)
			bf.MArr[index] = 1
		}
	}
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}

func (bf *BloomFilter) CheckFloat64(date float64) bool {
	datebyte := Float64ToByte(date)
	return bf.Check(datebyte)
}
func (bf *BloomFilter) Check(data []byte) bool {
	for i := 0; i < bf.KHash; i++ {
		index := bf.hashFun(data, i)
		if bf.MArr[index] == 0 {
			return false
		}
	}
	return true
}
func (bf *BloomFilter) hashFun(data []byte, count int) int64 {
	hash := int64(5381)
	for i := 0; i < len(data); i++ {
		hash = hash * 33 + int64(data[i]) + int64(count)
	}

	res := hash % (bf.MLength - 1)
	return int64(math.Abs(float64(res)))
}
