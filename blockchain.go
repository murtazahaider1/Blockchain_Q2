package blockchain

var Blockchain []Block

func init() {
    genesisBlock := Block{0, time.Now().String(), "Genesis Block", "", ""}
    genesisBlock.Hash = calculateHash(genesisBlock)
    Blockchain = append(Blockchain, genesisBlock)
}

func DisplayAllBlocks() {
    for _, block := range Blockchain {
        fmt.Printf("Index: %d\nTimestamp: %s\nData: %s\nPrevHash: %s\nHash: %s\n", block.Index, block.Timestamp, block.Data, block.PrevHash, block.Hash)
        fmt.Println("------------------------------------------------")
    }
}

func AddBlock(data string) {
    prevBlock := Blockchain[len(Blockchain)-1]
    newBlock := NewBlock(data, prevBlock.Hash)
    Blockchain = append(Blockchain, newBlock)
}

// ModifyBlock is not a standard blockchain operation due to immutability of blocks.
// However, for educational purposes, here's a simple way to modify a block's data.
// This function does not recompute hashes, so it will break the chain's integrity.
func ModifyBlock(index int, newData string) {
    if index < len(Blockchain) {
        Blockchain[index].Data = newData
    }
}
