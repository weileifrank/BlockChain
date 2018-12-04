package main

import (
	"fmt"
	"time"
)

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
		timeFormat := time.Unix(int64(block.TimeStamp), 0).Format("2006-01-02 15:04:05")
		fmt.Printf("时间戳: %s\n", timeFormat)
		//fmt.Printf("时间戳: %d\n", block.TimeStamp)
		fmt.Printf("难度值(随便写的）: %d\n", block.Difficulty)
		fmt.Printf("随机数 : %d\n", block.Nonce)
		fmt.Printf("当前区块哈希值: %x\n", block.Hash)
		fmt.Printf("区块数据 :%s\n", block.Transactions[0].TXInputs[0].Sig)
		fmt.Println("============================")
	}
	fmt.Println("打印区块结束")
}

func (cli *CLI) GetBalance(address string) {
	utxos := cli.bc.FindUTXOs(address)
	total:=0.0
	for _, utxo := range utxos {
		total +=utxo.Value
	}
	fmt.Printf("\"%s\"的余额为：%f\n", address, total)
}

func (cli *CLI) Send(from, to string, amount float64, miner, data string) {
	coinbase := NewCoinbaseTX(miner, data)
	tx := NewTransaction(from, to, amount, cli.bc)
	if tx == nil{
		return
	}
	cli.bc.AddBlock([]*Transaction{coinbase,tx})
	fmt.Println("转账成功")
}
func (cli *CLI) NewWallet() {
	ws := NewWallets()
	address := ws.CreateWallet()
	fmt.Printf("地址：%s\n", address)
}

func (cli *CLI)ListAddresses()  {
	ws := NewWallets()
	addresses := ws.ListAllAddresses()
	for _, address := range addresses {
		fmt.Printf("地址:%s\n",address)
	}
}