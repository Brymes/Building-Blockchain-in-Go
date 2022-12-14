package main

import "log"

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		log.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		log.Printf("Data: %s\n", block.Data)
		log.Printf("Hash: %x\n", block.Hash)
		log.Println()
	}
}
