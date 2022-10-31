// Package search search
package search

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)


/**
 * @Author peng
 * @Description //TODO
 * @Date 16:36 2022/6/28
 * @Param index trapdoor
 * @return Similarity of each article
 **/


func Search(indexs [][]*mat.VecDense, trap []*mat.VecDense) {
	result := make([]float64, len(indexs))
	for index, value := range indexs {
		I0 := value[0]
		I1 := value[1]
		result[index] = mat.Dot(I0, trap[0]) + mat.Dot(I1, trap[1])
	}
	//sort.Slice(result, func(i, j int) bool {
	//	if result[i] > result[j] {
	//		return true
	//	}
	//	return false
	//})
	fmt.Println(result)
}
