package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Block struct {
	Hash string `json:"hash"`
	PrevHash string `json:"prev_hash,omitempty"`
	Data string `json:"data"`
	Height int `json:"height"`
}

type blockChain struct {
	Blocks []*Block
}

var (
	ErrNotFound = errors.New("block is not found.")
)

var (
	b *blockChain
 	once sync.Once
)

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
	newBlock := Block{"", getLastHash(), data, len(GetBlockChain().Blocks) + 1}
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

func (b *blockChain) GetBlock(height int) (*Block, error) {
	if(height > len(b.Blocks)){
		return nil, ErrNotFound
	}
	return b.Blocks[height - 1], nil
}