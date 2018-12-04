package main

import (
	"fmt"
	"os"
	"strconv"
)

type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA     "add data to blockchain"
	printChain               "print all blockchain data" 
	getBalance  --address ADDRESS "print balance for the address"
	send FROM TO AMOUNT MINER DATA "由FROM转AMOUNT给TO，由MINER挖矿，同时写入DATA"
	newWallet   "创建一个新的钱包(私钥公钥对)"
	listAddresses "列举所有的钱包地址"
`

func (this *CLI) Run() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(Usage)
		return
	}
	cmd := args[1]
	switch cmd {
	case "addBlock":
		fmt.Println("添加区块逻辑")
		if len(args) == 4 && args[2] == "--data" {
			data := args[3]
			this.AddBlock(data)
		} else {
			fmt.Println("添加区块的参数有误,请重新输入")
			fmt.Println(Usage)
		}
	case "printChain":
		fmt.Println("打印区块逻辑")
		this.PrintBlockChain()
	case "getBalance":
		fmt.Println("获取余额")
		if len(args) == 4 && args[2] == "--address" {
			address := args[3]
			this.GetBalance(address)
		} else {
			fmt.Println("获取参数有误,请重新输入")
			fmt.Println(Usage)
		}
	case "send":
		fmt.Printf("转账开始...\n")
		if len(args) != 7 {
			fmt.Printf("参数个数错误，请检查！\n")
			fmt.Printf(Usage)
			return
		}
		//./block send FROM TO AMOUNT MINER DATA "由FROM转AMOUNT给TO，由MINER挖矿，同时写入DATA"
		from := args[2]
		to := args[3]
		amount, _ := strconv.ParseFloat(args[4], 64) //知识点，请注意
		miner := args[5]
		data := args[6]
		this.Send(from, to, amount, miner, data)
	case "newWallet":
		fmt.Printf("创建新的钱包...\n")
		this.NewWallet()
	case "listAddresses":
		fmt.Println("列出所有地址")
		this.ListAddresses()
	default:
		fmt.Println("无效的命令,请重新输入")
		fmt.Println(Usage)
	}
}
