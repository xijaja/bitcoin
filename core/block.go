package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// 定义一个区块
type Block struct {
	Timestamp     int64  // 区块创建的时间
	Data          []byte // 区块的数据内容
	PrevBlockHash []byte // 上一个区块的哈希值，用以组成链
	Hash          []byte // 当前区块的哈希值，用以校验区块数据有效性
}

// 创建一个新的区块
func NewBlock(date string, PrevBlockHash []byte) (block *Block) {
	block = &Block{
		Timestamp:     time.Now().Unix(), // 时间
		Data:          []byte(date),      // 内容
		PrevBlockHash: PrevBlockHash,     // 上个区块的哈希
	}
	block.setHash() // 添加当前区块的哈希
	return block
}

// 为区块添加哈希
func (b *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// 创世第一块，创世纪块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}) // 创世纪块非常特殊，他之前没有其他区块
}
