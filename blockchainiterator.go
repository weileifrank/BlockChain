package main

import "github.com/boltdb/bolt"

type BlockChainIterator struct {
	db *bolt.DB
	currentHashPointer []byte
}

func (bc *BlockChain)NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		bc.db,
		bc.tail,
	}
}

func (this *BlockChainIterator)Next() *Block {
	var block *Block
	db := this.db
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCKBUCKET))
		if bucket == nil{
			panic("区块尚未初始化")
			return nil
		}
		blockTmp := bucket.Get(this.currentHashPointer)
		if len(blockTmp) == 0 {
			block = nil
		}else{
			block = Deserialize(blockTmp)
			this.currentHashPointer = block.PreHash
		}
		return nil
	})
	return block
}
