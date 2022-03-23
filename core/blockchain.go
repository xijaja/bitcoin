package core

// 将区块串成链
type BlockChain struct {
	Blocks []*Block
}

// 添加区块到链中
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]   // 获取前一个区块
	newBlock := NewBlock(data, prevBlock.Hash) // 创建一个新的区块
	bc.Blocks = append(bc.Blocks, newBlock)    // 将区块添加入链
}

// 创建一个新的创世纪块，一个新的区块链就此诞生
func NewBlockChain() *BlockChain {
	return &BlockChain{
		Blocks: []*Block{
			NewGenesisBlock(), // 区块链中的第一块
			// ... 这是之后的其他块
		},
	}
}
