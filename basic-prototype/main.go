package main

import (
	"blockchain-with-go/basic-prototype/prototype"
	"fmt"
)

func main() {
	bc := prototype.NewBlockchain()

	bc.AddBlock("I like donuts")
	bc.AddBlock("I like donutsca07ca")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
