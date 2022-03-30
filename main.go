package main

import (
	"bitcoin/core"
	"fmt"
	"strconv"
)

func main() {
	// 初始化区块链
	bc := core.NewBlockChain()

	// 向区块链中添加区块
	bc.AddBlock("张三借给我100块")
	bc.AddBlock("我还给张三101块")
	fmt.Println("--------- --------- --------- --------- ---------")

	// 让我们看看这条链上的区块
	// fmt 占位符 %x 表示十六进制（其中：%x 输出的字母小写 %X 输出的字母大写）
	for _, block := range bc.Blocks {
		fmt.Printf("前块哈希：%x\n", block.PrevBlockHash)
		fmt.Printf("本块数据：%s\n", block.Data)
		fmt.Printf("本块哈希：%x\n", block.Hash)
		fmt.Printf("计算次数：%x\n", strconv.Itoa(block.Nonce)) // 工作量证明转string
		pow := core.NewProofofWork(block)
		fmt.Printf("是否有效：%v\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
