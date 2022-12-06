package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Hash string
	PrevHash string
	Data string
}

type blockChain struct {
	Blocks []*Block
}

var b *blockChain
var once sync.Once

func getLastHash() string{
	totalBlocks := len(b.Blocks)
	if(totalBlocks == 0){
		return ""
	}
	return b.Blocks[totalBlocks-1].Hash
}

func (b *Block)calculateCurrentHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func createBlock(data string) *Block{
	newBlock := Block{"", getLastHash(), data}
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
	b.Blocks = append(b.Blocks, createBlock(data))
}

func (b *blockChain) AllBlocks() []*Block {
	return b.Blocks
}