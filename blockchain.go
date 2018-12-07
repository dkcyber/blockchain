package main

import (
	"fmt"
	"github.com/bolt"
	"log"
	"os"
)

const (
	genesisInfo="The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"
	blockchainDB="blockchain.db"
	blockchainBucket="blockchainBucket"
	lastBlockHash="lastBlockHash"
)
//区块链结构体
type BlockChain struct {
	Db *bolt.DB //数据库句柄
	Tail []byte //最后一个区块哈希值
}

//创建区块链
func NewBlockChain()*BlockChain{
	//打开数据库
	db,err:=bolt.Open(blockchainDB,0600,nil)
	if err!=nil{
		log.Panic(err)
	}
	//<---不用关闭数据库，等调用者关闭（main.go）
	var tail []byte
	db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(blockchainBucket))
		if bucket==nil{
			//抽屉为空,证明还没有任何数据，此时创建创世块
			fmt.Printf("bucket不存在，准备创建!\n")
			bucket,err=tx.CreateBucket([]byte(blockchainBucket))
			if err!=nil{
				log.Panic(err)
			}
			//创建创世块
			block:=NewBlock(genesisInfo,[]byte{})
			//添加数据到数据库 区块哈希--区块数据 lastBlockHash--hash
			err=bucket.Put(block.Hash,block.Serialize())
			if err!=nil{
				log.Panic(err)
			}
			err=bucket.Put([]byte(lastBlockHash),block.Hash)
			if err!=nil{
				log.Panic(err)
			}
			//更新区块链Tail
			tail=block.Hash
		}else{
			//抽屉不为空，代表有数据，直接取出最后一个hash
			tail=bucket.Get([]byte(lastBlockHash))
		}

		return nil
	})

	return &BlockChain{db,tail}
}

//添加区块到区块链
func(bc *BlockChain)AddBlock(data string){

	bc.Db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(blockchainBucket))
		if bucket==nil{
			fmt.Printf("bucket不存在，请检查!\n")
			os.Exit(1)
		}
		//获取最后的一个哈希，创建新的结构体
		//lastHash:=bucket.Get([]byte(lastBlockHash))   <---记得直接可以在内存中获取，无需去数据库取
		block:=NewBlock(data,bc.Tail)
		//加入到数据库，并更新内存数据
		bucket.Put(block.Hash,block.Serialize())
		bucket.Put([]byte(lastBlockHash),block.Hash)
		bc.Tail=block.Hash
		return nil
	})
}

//迭代器结构体
type BlockChainIterator struct {
	Db *bolt.DB //数据库句柄
	Current []byte //当前游标
}

//创建迭代器
func(bc *BlockChain)NewIterator()*BlockChainIterator{
	return &BlockChainIterator{bc.Db,bc.Tail}
}

//用来迭代数据，返回block 并游标前移
func (it *BlockChainIterator)Next()*Block{
	var block Block
	it.Db.View(func(tx *bolt.Tx) error {
		//从数据库中，根据current,获取当前block的数据，并反序列
		bucket:=tx.Bucket([]byte(blockchainBucket))
		if bucket==nil{
			fmt.Printf("bucket不存在，请检查!\n")
			os.Exit(1)
		}
		blockBytes:=bucket.Get(it.Current)
		block=*Deserialize(blockBytes)
		//current前移
		it.Current=block.PreBlockHash
		return nil
	})
	return &block
}