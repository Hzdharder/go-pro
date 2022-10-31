package blockchain

import (
	"BlookChainPro/package/block"
"BlookChainPro/package/makeblock"
)

type BlockChain struct {
	Blocks [] *block.Block //一个数组,每个元素都是指针,存储block区块的地址

}
//增加一个区块
func (blocks *BlockChain) AddBlock(data string)  {
prevBlock:=blocks.Blocks[len(blocks.Blocks)-1] //取出最后一个区块
newBlock:= makeblock.NewBlock(data,prevBlock.Hash) //创建一个区块
blocks.Blocks =append(blocks.Blocks,newBlock)  //区块链插入一个新的区块
}

//创建一个区块链
func NewBlockchain()*BlockChain  {
return  &BlockChain{[]*block.Block{makeblock.NewGenesisBlock()}}

}
