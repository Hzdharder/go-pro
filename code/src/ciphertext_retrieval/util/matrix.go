package util

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

// 根据行列生成随机矩阵
// rand.Seed(time.Now().UnixNano())保证每次都不相同
// 矩阵中每个元素都是int 范围是[0,n)

func  GenRandMatrix(column, row ,scope int) *mat.Dense {
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	randMatrix := mat.NewDense(column, row,nil)
	for i := 0; i < column; i++ {
		for j := 0; j < row; j++ {
			randMatrix.Set(i, j, float64(rander.Intn(scope) ))
		}
	}
	return randMatrix
}

// 根据行列和指定值生成随机矩阵
// 矩阵中每个元素都是int

func GenMatrix(column, row ,value int) *mat.Dense {
	randMatrix := mat.NewDense(column, row,nil)

	for i := 0; i < column; i++ {
		for j := 0; j < row; j++ {
			randMatrix.Set(i, j, float64(value))
		}
	}
	return randMatrix
}

// 生成正定矩阵

func GenSymmetricPositiveDefiniteMatrix(column, row  int) *mat.Dense {
	randMatrix := mat.NewDense(column, row,nil)
	for i := 0; i < column; i++ {
		for j := 0; j < row; j++ {
			if i == j {
				randMatrix.Set(i, j, 1.0)
			} else {
				randMatrix.Set(i , j,0.0)
			}
		}
	}
	return randMatrix
}

//生成向量 n 维度 value为向量的值

func GenVecDense(n ,value int) *mat.VecDense {
	randMatrix := mat.NewVecDense(n,nil)
	for i := 0; i < n ; i++ {
		randMatrix.SetVec(i, float64(value))
	}
	return randMatrix
}
