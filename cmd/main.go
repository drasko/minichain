package main

import (
  "fmt"
  mc "github.com/drasko/minichain"
)

func main() {
  bc := mc.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Drasko")
	bc.AddBlock("Send 2 more BTC to Drasko")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
