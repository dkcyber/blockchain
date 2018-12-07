package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const Bits=20
//工作量证明结构体
type ProofOfWork struct {
	Block *Block
	Target *big.Int //目标数值
}

//生成一个pow
func NewProofOfWork(block *Block)*ProofOfWork{
	//生成Target值
	target:=big.NewInt(1)
	target.Lsh(target,256-Bits)

	return &ProofOfWork{block,target}
}

//计算生成nonce和hash值
func (pow *ProofOfWork)Run()(uint64,[]byte){
	var nonce uint64
	var hash [32]byte

	for{
		fmt.Printf("%x\r", hash)
		//计算sha256
		hash=sha256.Sum256(pow.PrepareData(nonce))
		//转成big.Int
		var temp big.Int
		temp.SetBytes(hash[:])
		//算出的数据小于目标值则完成挖矿
		if temp.Cmp(pow.Target)==-1{
			fmt.Printf("挖矿成功！nonce: %d, 哈希值为: %x\n", nonce, hash)
			break
		}else{
			nonce++
		}
	}
	return nonce,hash[:]
}

//拼接block里面的数据（包括nonce），返回[]byte用来计算hash
func(pow *ProofOfWork)PrepareData(nonce uint64)[]byte{
	block:=pow.Block

	blockInfos:=[][]byte{
		Uint64ToBytes(block.Version),
		block.PreBlockHash,
		block.MerkleRoot,
		Uint64ToBytes(block.TimpStamp),
		Uint64ToBytes(block.Difficulty),
		Uint64ToBytes(nonce),
		block.Data,
	}

	data:=bytes.Join(blockInfos,[]byte{})

	return data
}