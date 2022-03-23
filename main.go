package main

import (
	"bitcoin/core"
	"fmt"
)

func main() {
	// 初始化区块链
	bc := core.NewBlockChain()

	// 向区块链中添加区块
	bc.AddBlock("张三借给我100块")
	bc.AddBlock("我需要还给张三101块")

	// 让我们看看这条链上的区块
	// fmt 占位符 %x 表示十六进制（其中：%x 输出的字母小写 %X 输出的字母大写）
	for _, block := range bc.Blocks {
		fmt.Printf("prev hash %x\n", block.PrevBlockHash)
		fmt.Printf("Data %s\n", block.Data)
		fmt.Printf("Hash %x\n", block.Hash)
		fmt.Println()
	}
}
