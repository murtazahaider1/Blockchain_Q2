package blockchain

import (
    "testing"
)

func TestAddBlock(t *testing.T) {
    blockchain := InitializeBlockchain() 
    blockchain.AddBlock("Test Block")
    if len(blockchain.Blocks) != 2 { 
        t.Errorf("Expected Blockchain length of 2, got %d", len(blockchain.Blocks))
    }
}

