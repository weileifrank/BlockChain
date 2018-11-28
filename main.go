package main

func main() {
	bc := NewBlockChain()
	cli := CLI{bc}
	cli.Run()

	//bc := NewBlockChain()
	//bc.AddBlock("frank")
	//bc.AddBlock("bupin")
	//
	//iterator := bc.NewIterator()
	//for{
	//	block := iterator.Next()
	//	if block == nil{
	//		fmt.Println("遍历结束")
	//		break
	//	}
	//	fmt.Printf("版本号: %d\n", block.Version)
	//	fmt.Printf("前区块哈希值: %x\n", block.PreHash)
	//	fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
	//	fmt.Printf("时间戳: %d\n", block.TimeStamp)
	//	fmt.Printf("难度值(随便写的）: %d\n", block.Difficulty)
	//	fmt.Printf("随机数 : %d\n", block.Nonce)
	//	fmt.Printf("当前区块哈希值: %x\n", block.Hash)
	//	fmt.Printf("区块数据 :%s\n", block.Data)
	//	fmt.Println("============================")
	//}
}
