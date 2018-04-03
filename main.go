package main

import (
  "fmt"
  "./DCoin"
  "strconv"
)

func main() {
	bc := DCoin.NewBlockchain()
  
	bc.AddBlock("Send 1 Dcoin to Alice")
	bc.AddBlock("Send 2 Dcoin to Bob")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := DCoin.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		
		fmt.Println()
	}
}