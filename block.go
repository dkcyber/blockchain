package main

import "time"

//区块结构体
type Block struct {
	Version uint64  //版本号
	PreBlockHash []byte //前一个区块的哈希
	MerkleRoot []byte //梅克尔根
	TimpStamp uint64 //区块创建时间
	Difficulty uint64 //挖矿难度值
	Nonce uint64 //随机数
	Data []byte //区块交易数据
	Hash []byte //本区块哈希，为了操作方便，实际不含此项
}

//创建区块
func NewBlock(data string,prevBlockHash []byte)*Block{
	//实例化一个区块，参数先填默认值，后续赋值
	block:=Block{
		Version:00,
		PreBlockHash:prevBlockHash,
		MerkleRoot:[]byte{},
		TimpStamp:uint64(time.Now().Unix()),
		Difficulty:Bits,
		Nonce:10,
		Data:[]byte(data),
		Hash:[]byte{},
	}

	pow:=NewProofOfWork(&block)
	nonce,hash:=pow.Run()
	block.Nonce=nonce
	block.Hash=hash

	return &block
}
