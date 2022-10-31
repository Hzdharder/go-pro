package makeblock

import (
	"BlookChainPro/package/block"
	"BlookChainPro/package/proofofwook"
	"time"
)

func   NewBlock(data string,prevBlockHash []byte ) *block.Block  {
	//取得一个对象初始化之后的一个地址
	block:=&block.Block{time.Now().Unix(),[]byte(data),prevBlockHash,[]byte{},0}

	//block.SetHash()
	pow:= proofofwook.NewProofOfWork(block) //挖矿附加这个区块
	nonce,hash:=pow.Run()                   //开始挖矿
	block.Hash=hash[:]
	block.Nonce=nonce
	return  block
}
//创建创世区块
func   NewGenesisBlock() *block.Block {
	return  NewBlock("这是一个创世区块888888",[]byte{})
}