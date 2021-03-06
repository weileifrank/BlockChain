package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

const reward = 50

//定义交易结构
type Transaction struct {
	TXID     []byte
	TXInputs []TXInput
	TXOnputs []TXOutput
}

//定义交易输入
type TXInput struct {
	//引用交易的ID
	TXid []byte
	//引用的output的索引值
	Index int64
	//解锁脚本 我们用地址来模拟
	Sig string
}

//定义交易输出
type TXOutput struct {
	//转账金额
	Value float64
	//锁定脚本
	PubKeyHash string
}

//设置交易id
func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		panic(err)
	}
	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}

//创建挖矿交易
func NewCoinbaseTX(address string, data string) *Transaction {
	input := TXInput{[]byte{}, -1, data}
	output := TXOutput{reward, address}
	tx := Transaction{[]byte{}, []TXInput{input}, []TXOutput{output}}
	tx.SetHash()
	return &tx;
}

//判断是否为挖矿交易
func (tx *Transaction) IsCoinbase() bool {
	//if len(tx.TXInputs) == 1 {
	//	input := tx.TXInputs[0]
	//	if !bytes.Equal(input.TXid, []byte{}) || input.Index != -1 {
	//		return false
	//	}
	//}
	if len(tx.TXInputs) == 1 && len(tx.TXInputs[0].TXid) == 0 && tx.TXInputs[0].Index == -1 {
		return true
	}
	return false
}

//创建普通交易
func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {
	utxos, resValue := bc.FindNeedUTXOs(from, amount)
	if resValue < amount {
		fmt.Println("余额不足,交易失败")
		return nil
	}
	var inputs []TXInput
	var outputs []TXOutput

	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input:=TXInput{[]byte(id),int64(i),from}
			inputs = append(inputs, input)
		}
	}
	output:=TXOutput{amount,to}
	outputs= append(outputs, output)
	if resValue > amount {
		outputs = append(outputs, TXOutput{resValue- amount,from})
	}
	tx:=Transaction{[]byte{},inputs,outputs}
	tx.SetHash()
	return &tx
}
