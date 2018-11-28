package main

import (
	"github.com/boltdb/bolt"
)

type BlockChain struct {
	//blocks []*Block
	db *bolt.DB
	tail []byte
}
const BLOCKCHAINDB = "blockChain.db"
const BLOCKBUCKET = "blockBucket"
const LASTHASHKEY = "LastHashKey"

//创建区块链
func NewBlockChain() *BlockChain {
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
			genesisBlock:=GenesisBlock()
			bucket.Put(genesisBlock.Hash,genesisBlock.Serialize())
			bucket.Put([]byte(LASTHASHKEY),genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		}else{
			lastHash = bucket.Get([]byte(LASTHASHKEY))
		}

		return nil
	})
	return &BlockChain{db,lastHash}
}

//添加区块
func (this *BlockChain) AddBlock(data string) {
	db:=this.db
	lastHash:=this.tail
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCKBUCKET))
		if bucket == nil{
			panic("不能为空")
		}
		block := NewBlock(data, lastHash)
		bucket.Put(block.Hash,block.Serialize())
		bucket.Put([]byte(LASTHASHKEY),block.Hash)
		this.tail = block.Hash
		return nil;
	})


}

//定义创世区块
func GenesisBlock() *Block {
	return NewBlock("区块首页", []byte{})
}
