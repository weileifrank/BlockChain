package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	//版本号
	Version uint64
	//2前区块hash
	PreHash []byte
	//3 梅克尔跟
	MerkelRoot []byte
	//时间戳
	TimeStamp uint64
	//难度值
	Difficulty uint64
	//随机数
	Nonce uint64
	//当前hash
	Hash []byte
	//数据
	Data []byte
}

//创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PreHash:    prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		Data:       []byte(data),
	}
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

//序列化
func (this *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&this)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}
//反序列化
func Deserialize(data []byte) *Block {
	var block *Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}
	return block
}

//uint64 to byte
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

// 给区块赋值
//func (this *Block) SetHash() {
//	tmp := [][]byte{
//		Uint64ToByte(this.Version),
//		this.PreHash,
//		this.MerkelRoot,
//		Uint64ToByte(this.TimeStamp),
//		Uint64ToByte(this.Difficulty),
//		Uint64ToByte(this.Nonce),
//		this.Data,
//	}
//	blockInfo := bytes.Join(tmp, []byte{})
//	hash := sha256.Sum256(blockInfo)
//	this.Hash = hash[:]
//}
