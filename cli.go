package main

import (
	"fmt"
	"os"
)

const Usage  = `
	./blockchain addBlock "xxxxxx"   添加数据到区块链
	./blockchain printChain          打印区块链
`

//命令行结构体,主要是操作blockchain
type CLI struct {
	bc *BlockChain
}

//给CLI提供一个方法，进行命令解析，从而执行调度
func(cli *CLI)Run(){
	//获取命令行参数
	cmds:=os.Args

	if len(cmds)<2{
		fmt.Printf(Usage)
		os.Exit(1)
	}

	switch cmds[1] {
	case "addBlock":
		if len(cmds)!=3{
			fmt.Printf("无效命令，请按以下输入：\n")
			fmt.Printf(Usage)
			os.Exit(1)
		}
		fmt.Printf("添加区块命令被调用, 数据：%s\n", cmds[2])
		data:=cmds[2]
		cli.AddBlock(data)
	case "printChain":
		fmt.Printf("打印区块链命令被调用\n")
		cli.PrintChain()
	default:
		fmt.Printf("无效命令，请按以下输入：\n")
		fmt.Printf(Usage)
	}
	//添加区块的时候： bc.addBlock(data), data 通过os.Args拿回来
	//打印区块链时候：遍历区块链，不需要外部输入数据
}
