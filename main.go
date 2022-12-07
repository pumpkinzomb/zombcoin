package main

import (
	"github.com/pumpkinzomb/zombcoin/explorer"
	"github.com/pumpkinzomb/zombcoin/rest"
)

func main(){
	go rest.Run(4000)
	explorer.Run(4001)	
}