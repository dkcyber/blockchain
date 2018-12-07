package main

func main(){
	//主函数，负责运行命令行启动
	bc:=NewBlockChain()
	defer bc.Db.Close()
	cli:=CLI{bc}
	cli.Run() //根据输入的命令，自动分配操作
}
