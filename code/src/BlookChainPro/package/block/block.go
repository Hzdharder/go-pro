package block

//定义区块
type Block struct {
	Timestamp int64 //时间
	Data []byte   //交易数据
	PrevBlockHash []byte//上一个区块hash
	Hash []byte //当前区块数据hash
	Nonce  int//工作量证明
}
/*
//设定结构体对象的hash
func  (block *Block )SetHash(){
	//处理当前时间转化为十进制字符串，再转化为字符集
timestape:=[]byte(strconv.FormatInt(block.Timestamp,10))
//叠加要hash的数据
headers:=bytes.Join([][]byte{block.PrevBlockHash,block.Data,timestape},[]byte{})
//计算出hash地址
hash:=sha256.Sum256(headers)
block.Hash=hash[:] //设置hash
}

 */

//创建一个区块
