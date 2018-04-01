package HanCoin

import (
  "fmt"
  "./HanCoin"
)

func main() {
	bc := HanCoin.NewBlockchain()
  
	bc.AddBlock("Send 1 BTC to Alice")
	bc.AddBlock("Send 2 more BTC to Alice")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}