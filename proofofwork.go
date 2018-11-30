package main

import (
	"bytes"
	"crypto/sha256"
	"math/big"
)

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		Block: block,
	}
	//我们指定的难度值，现在是一个string类型，需要进行转换
	targetStr := "0000f00000000000000000000000000000000000000000000000000000000000"
	tmpInt := big.Int{}
	tmpInt.SetString(targetStr, 16)
	pow.Target = &tmpInt
	return &pow
}

//计算hash的函数 -run
func (this *ProofOfWork) Run() ([]byte, uint64) {
	var hash [32]byte
	var nonce uint64
	for {
		tmp := [][]byte{
			Uint64ToByte(this.Block.Version),
			this.Block.PreHash,
			this.Block.MerkelRoot,
			Uint64ToByte(this.Block.TimeStamp),
			Uint64ToByte(this.Block.Difficulty),
			Uint64ToByte(nonce),
			//this.Block.Data,
		}
		blockInfo := bytes.Join(tmp, []byte{})
		hash = sha256.Sum256(blockInfo)
		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])
		if tmpInt.Cmp(this.Target) == -1 {
			return hash[:], nonce
		} else {
			nonce++
		}
	}
}
