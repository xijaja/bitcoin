package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

// 目标（调整这个数字可以变更计算的难度，数字越小哈希的前置0越少）
const targetBits = 20

// 工作证明（对区块的计算要满足这个目标）
type ProofofWork struct {
	block  *Block   // 区块
	target *big.Int // 目标
}

// 通过传入区块获取其工作证明
func NewProofofWork(b *Block) *ProofofWork {
	target := big.NewInt(1)
	// 左移运算，target 的二进制位全部左移uint(256-20)位
	// 也理解为 target 乘以 2 的 uint(256-20) 次方
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofofWork{b, target}
	return pow
}

func (pow *ProofofWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,       // 前个区块的哈希
			pow.block.Data,                // 当前区块的数据
			IntToHex(pow.block.Timestamp), // 当前区块的时间
			IntToHex(int64(targetBits)),   // 目标值
			IntToHex(int64(nonce)),        // 工作量
		},
		[]byte{},
	)
	return data
}

// Run执行工作证明
func (pow *ProofofWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("开始挖掘，这个区块包含数据：\"%s\"\n", pow.block.Data)

	// 计算次数，在int64最大值内
	for nonce < math.MaxInt64 {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data) // 计算哈希值
		fmt.Printf("\r%x", hash)   // 在CLI的同一行中反复打印
		hashInt.SetBytes(hash[:])  // 将哈希值转换成一个整数

		if hashInt.Cmp(pow.target) == -1 {
			// 如果 hashInt 小于 目标值
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}
