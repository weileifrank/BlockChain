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
	spentOutputs:=make(map[string][]int64)
	it := bc.NewIterator()
	for{
		block:= it.Next()
		if block == nil{
			break
		}
		for _,tx := range block.Transactions {
			fmt.Println("current txid:",tx.TXID)
		//OUTPUT:
			for _, output := range tx.TXOnputs {
				if output.PubKeyHash == address {
					UTXO = append(UTXO,output)
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
	return UTXO
}
