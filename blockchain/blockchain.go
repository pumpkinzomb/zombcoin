package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Hash string
	PrevHash string
	Data string
}

type blockChain struct {
	blocks []*block
}

var b *blockChain
var once sync.Once

func getLastHash() string{
	totalBlocks := len(b.blocks)
	if(totalBlocks == 0){
		return ""
	}
	return b.blocks[totalBlocks-1].Hash
}

func (b *block)calculateCurrentHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func createBlock(data string) *block{
	newBlock := block{"", getLastHash(), data}
	newBlock.calculateCurrentHash()
	return &newBlock
}

func GetBlockChain() *blockChain {
	if(b == nil){
		once.Do(func() {
			b = &blockChain{}
			b.AddBlock("This is New genesis block.")
		})
	}
	return b
}

func (b *blockChain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func (b *blockChain) AllBlocks() []*block {
	return b.blocks
}