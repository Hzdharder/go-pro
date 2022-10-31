package proofofwook

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"BlookChainPro/package/utils"

	"BlookChainPro/package/block"
	"math"
	"math/big"
)

var(
	maxNonce=math.MaxInt64//最大的64位整数
)

const targetBits  =24//对比的位数

type ProofOfWork struct {
	block *block.Block
	target *big.Int//存储计算哈希值对比的特定整数
}

//创建一个工作量证明的挖矿对象
func NewProofOfWork(block *block.Block)*ProofOfWork  {
	target:=big.NewInt(1)//初始化目标整数
	target.Lsh(target,uint(256-targetBits))//数据转换
	pow:=&ProofOfWork{block,target}
	return pow

}
//准备数据进行挖矿计算
func (pow *ProofOfWork)prepareData(nonce int)  []byte{
data:=bytes.Join(
	[][]byte{pow.block.PrevBlockHash, //上一块哈希
		pow.block.Data,              //当前数据
		utils.IntTOHex(pow.block.Timestamp),
		utils.IntTOHex(int64(targetBits)),
		utils.IntTOHex(int64(nonce)),
	},[]byte{},
	)
return  data
}
//挖矿执行
func (pow *ProofOfWork)Run()(int,[]byte)  {
    var hashInt big.Int
	var  hash  [32]byte
    nonce:=0
    fmt.Printf("当前挖矿计算的区块数据%s",pow.block.Data)
    for nonce<maxNonce{
    	data:=pow.prepareData(nonce)//准备好的数据
    	hash=sha256.Sum256(data)//计算出hash
    	fmt.Printf("\r%x",hash)//打印显示hash
    	hashInt.SetBytes(hash[:])//获取要对比的数据
    	if hashInt.Cmp(pow.target)==-1{  //挖矿校验
    		break
		}else {
			nonce++
		}

		//return nonce,hash[:] //nonce为解题的答案，hash为当前的hash
	}
	fmt.Println("\n\n")
	return nonce,hash[:] //nonce为解题的答案，hash为当前的hash
}
//校验
func (pow *ProofOfWork)Validate() bool {
	var hashInt big.Int
	data:=pow.prepareData(pow.block.Nonce)
	hash:=sha256.Sum256(data)//计算出hash
	hashInt.SetBytes(hash[:])//获取要对比的数据
	isValid:=(hashInt.Cmp(pow.target)==-1)//校验数据
	return  isValid
}