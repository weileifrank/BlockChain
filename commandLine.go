package main

import "fmt"

func (cli *CLI) AddBlock(data string) {
	//cli.bc.AddBlock(data)
	fmt.Println("添加区块成功")
}
func (cli *CLI) PrintBlockChain() {
	iterator := cli.bc.NewIterator()
	for {
		block := iterator.Next()
		if block == nil {
			fmt.Println("遍历结束")
			break
		}
		fmt.Printf("版本号: %d\n", block.Version)
		fmt.Printf("前区块哈希值: %x\n", block.PreHash)
		fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
		fmt.Printf("时间戳: %d\n", block.TimeStamp)
		fmt.Printf("难度值(随便写的）: %d\n", block.Difficulty)
		fmt.Printf("随机数 : %d\n", block.Nonce)
		fmt.Printf("当前区块哈希值: %x\n", block.Hash)
		fmt.Printf("区块数据 :%s\n", block.Transactions[0].TXInputs[0].Sig)
		fmt.Println("============================")
	}
	fmt.Println("打印区块结束")
}

func (cli *CLI) GetBalance(address string) {
	utxo := cli.bc.FindUTXOs(address)
	fmt.Println(utxo)
}
