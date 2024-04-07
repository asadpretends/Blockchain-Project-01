package main

import (
	"fmt"
	"strconv"

	blockchain "github.com/asadpretends/Blockchain-Project-01/Blockchain"
)


func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block:%s\n", block.Data)
		fmt.Printf("Hash: %x\n",block.Hash )

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
		
	}
	
}