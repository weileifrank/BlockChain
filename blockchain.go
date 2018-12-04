package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

type BlockChain struct {
	//blocks []*Block
	db   *bolt.DB
	tail []byte
}

const BLOCKCHAINDB = "blockChain.db"
const BLOCKBUCKET = "blockBucket"
const LASTHASHKEY = "LastHashKey"

//创建区块链
func NewBlockChain(address string) *BlockChain {
	var lastHash []byte
	db, err := bolt.Open(BLOCKCHAINDB, 0600, nil)
	if err != nil {
		panic(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCKBUCKET))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(BLOCKBUCKET))
			if err != nil {
				panic(err)
			}
			genesisBlock := GenesisBlock(address)
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			bucket.Put([]byte(LASTHASHKEY), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte(LASTHASHKEY))
		}

		return nil
	})
	return &BlockChain{db, lastHash}
}

//添加区块
func (this *BlockChain) AddBlock(txs []*Transaction) {
	db := this.db
	lastHash := this.tail
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCKBUCKET))
		if bucket == nil {
			panic("不能为空")
		}
		block := NewBlock(txs, lastHash)
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte(LASTHASHKEY), block.Hash)
		this.tail = block.Hash
		return nil;
	})

}

//定义创世区块
func GenesisBlock(address string) *Block {
	coinbase := NewCoinbaseTX(address, "首页老牛逼了")
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func (bc *BlockChain) FindUTXOs(address string) []TXOutput {
	var UTXO []TXOutput
	txs := bc.FindUTXOTransactions(address)
	for _, tx := range txs {
		for _, output := range tx.TXOnputs {
			if address == output.PubKeyHash {
				UTXO = append(UTXO,output)
			}
		}
	}
	return UTXO
}

func (bc *BlockChain)FindNeedUTXOs(from string, amount float64)(map[string][]uint64,float64)  {
	var utxos = make(map[string][]uint64)
	var calc float64
	txs := bc.FindUTXOTransactions(from)
	for _, tx := range txs {
		for i, output := range tx.TXOnputs {
			if from == output.PubKeyHash {
				if calc < amount {
					utxos[string(tx.TXID)] = append(utxos[string(tx.TXID)], uint64(i))
					calc+=output.Value
					if calc >= amount {
						return utxos,calc
					}
				}else{
					fmt.Printf("不满足转账金额,当前总额：%f， 目标金额: %f\n", calc, amount)
				}
			}
		}
	}
	return utxos,calc
}

func (bc *BlockChain)FindUTXOTransactions(address string) []*Transaction {
	var txs []*Transaction
	spentOutputs:=make(map[string][]int64)
	it:=bc.NewIterator()
	for{
		block := it.Next()
		if block == nil{
			return txs
		}
		for _, tx := range block.Transactions {
			OUTPUT:
			for i, output := range tx.TXOnputs {
				if spentOutputs[string(tx.TXID)] != nil{
					for _, j := range spentOutputs[string(tx.TXID)] {
						if int64(i) == j{
							continue OUTPUT
						}
					}
				}
				if output.PubKeyHash == address {
					txs = append(txs, tx)
				}
			}
			if !tx.IsCoinbase() {
				for _, input := range tx.TXInputs {
					if input.Sig == address{
						spentOutputs[string(input.TXid)] = append(spentOutputs[string(input.TXid)],input.Index)
					}
				}
			}
		}
	}
	return txs
}



















