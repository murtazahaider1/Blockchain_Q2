package blockchain

import (
    "crypto/sha256"
    "encoding/hex"
    "strconv"
    "time"
)

type Block struct {
    Index     int
    Timestamp string
    Data      string
    PrevHash  string
    Hash      string
}

func calculateHash(block Block) string {
    record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}

func NewBlock(data string, prevHash string) Block {
    block := Block{0, time.Now().String(), data, prevHash, ""}
    block.Index = len(Blockchain) + 1
    block.Hash = calculateHash(block)
    return block
}
