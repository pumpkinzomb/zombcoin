package main

import (
	"fmt"

	"github.com/pumpkinzomb/zombcoin/blockchain"
)




func main(){
	myBlockChain := blockchain.GetBlockChain()
	
	myBlockChain.AddBlock("Second Block")
	myBlockChain.AddBlock("Third Block")
	myBlockChain.AddBlock("Fourth Block")
	myBlockChain.AddBlock("Fifth Block")

	for _, block := range myBlockChain.AllBlocks(){
		fmt.Printf("Hash: %s \n", block.Hash)
		fmt.Printf("Prev Hash: %s \n", block.PrevHash)
		fmt.Printf("Data: %s \n", block.Data)
	}

}