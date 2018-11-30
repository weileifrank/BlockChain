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
	getBalance  --address ADDRESS "print balance for the address"
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
	case "getBalance":
		fmt.Println("获取余额")
		if len(args)==4&&args[2]=="--address"{
			address:=args[3]
			this.GetBalance(address)
		}else{
			fmt.Println("获取参数有误,请重新输入")
			fmt.Println(Usage)
		}
	default:
		fmt.Println("无效的命令,请重新输入")
		fmt.Println(Usage)
	}
}
