package main

import (
    "fmt"
    "github.com/murtazahaider1/Blockchain_Q2/Dev"
)


func main() {
    blockchain.AddBlock("First Block after Genesis")
    blockchain.AddBlock("Second Block after Genesis")
    blockchain.DisplayAllBlocks()

    blockchain.ModifyBlock(1, "Modified Data")
    fmt.Println("After modification:")
    blockchain.DisplayAllBlocks()
}
