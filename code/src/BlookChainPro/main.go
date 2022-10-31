package BlookChainPro

import (
	"BlookChainPro/package/blockchain"
	"BlookChainPro/package/proofofwook"
	"fmt"
	"strconv"
)
func main(){
	num:=1
	fmt.Println("this is a simply test ")
	bc:= blockchain.NewBlockchain() //创建一个区块链
	bc.AddBlock("someone pay 10BTC to me")
	bc.AddBlock("someone pay 20BTC to me")
	bc.AddBlock("someone pay 30BTC to me")
	for _,block:=range  bc.Blocks {
		fmt.Printf("*******************************第%v个区块****************************\n",num)
		fmt.Printf("上一块哈希:%x\n",block.PrevBlockHash)
		fmt.Printf("数据:%s\n",block.Data)
		fmt.Printf("当前哈希:%x\n",block.Hash)
		pow:= proofofwook.NewProofOfWork(block) //校验工作量
		fmt.Printf("pow %s:此区块通过检验为合法区块\n",strconv.FormatBool(pow.Validate()))
		fmt.Printf("**********************************************************************")
		num++
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}
}
