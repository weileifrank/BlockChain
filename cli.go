package main

import (
	"fmt"
	"os"
)

type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA     "add data to blockchain"
	printChain               "print all blockchain data" 
`

func (this *CLI) Run()  {
	args := os.Args
	if len(args)<2{
		fmt.Println(Usage)
		return
	}
	cmd:=args[1]
	switch cmd {
	case "addBlock":
		fmt.Println("添加区块逻辑")
		if len(args)==4&&args[2]=="--data"{
			data:=args[3]
			this.AddBlock(data)
		}else{
			fmt.Println("添加区块的参数有误,请重新输入")
			fmt.Println(Usage)
		}
	case "printChain":
		fmt.Println("打印区块逻辑")
		this.PrintBlockChain()
	default:
		fmt.Println("无效的命令,请重新输入")
		fmt.Println(Usage)
	}
}
