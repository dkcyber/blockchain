package main

import (
	"bytes"
	"fmt"
	"time"
)

/*实现命令行具体命令的操作*/

//命令行添加区块
func (cli *CLI)AddBlock(data string){
	cli.bc.AddBlock(data)
	fmt.Printf("添加区块成功!\n")
}

//命令行打印区块链数据
func (cli *CLI)PrintChain(){
	it:=cli.bc.NewIterator()
	for{
		block:=it.Next()
		fmt.Printf("++++++++++++++++++++++++++++++++\n")

		fmt.Printf("Version : %d\n", block.Version)
		fmt.Printf("PrevBlockHash : %x\n", block.PreBlockHash)
		fmt.Printf("MerKleRoot : %x\n", block.MerkleRoot)

		timeFormat := time.Unix(int64(block.TimpStamp), 0).Format("2006-01-02 15:04:05")
		fmt.Printf("TimeStamp : %s\n", timeFormat)

		fmt.Printf("Difficulity : %d\n", block.Difficulty)
		fmt.Printf("Nonce : %d\n", block.Nonce)

		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Printf("Data : %s\n", block.Data)

		pow := NewProofOfWork(block)
		fmt.Printf("IsValid: %v\n", pow.IsValid())

		if bytes.Equal(block.PreBlockHash,[]byte{}){
			fmt.Printf("区块链遍历结束!\n")
			break
		}
	}
}


