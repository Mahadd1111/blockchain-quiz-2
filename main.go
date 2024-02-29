package main

import (
	"fmt"
)

type Block struct {
	ID   int
	Data string
}

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) NewBlock(id int, data string) {
	block := &Block{
		ID:   id,
		Data: data,
	}
	bc.Blocks = append(bc.Blocks, block)
}

func (bc *Blockchain) DisplayAllBlocks() {
	for _, block := range bc.Blocks {
		fmt.Printf("Block ID: %d, Data: %s\n", block.ID, block.Data)
	}
}

func (bc *Blockchain) ModifyBlock(id int, newData string) {
	for _, block := range bc.Blocks {
		if block.ID == id {
			block.Data = newData
			break
		}
	}
}

func main() {

	// Creating Blockchain
	blockchain := &Blockchain{}

	// Adding Blocks
	blockchain.NewBlock(1, "DataXYZ")
	blockchain.NewBlock(2, "DataASD")
	blockchain.NewBlock(3, "DataQWE")

	// Printing Initial Blocks
	fmt.Printf("\nInitial Blockchain: \n\n")
	blockchain.DisplayAllBlocks()

	// Modifying Block with ID 2
	blockchain.ModifyBlock(2, "DataHJK")

	// Reprinting Updated Block
	fmt.Printf("\n\nUpdated Blockchain: \n\n")
	blockchain.DisplayAllBlocks()

}
