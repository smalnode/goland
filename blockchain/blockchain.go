package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// Block blockchain node
type Block struct {
	Index      int
	Timestamp  string
	Nonce      string
	Difficulty int
	Hash       string
	PrevHash   string

	// app data
	BPM int
}

// Blockchain global chain
var Blockchain []Block

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Nonce + string(block.Difficulty) + block.PrevHash + string(block.BPM)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, BPM int, dificulty int) (Block, error) {
	var newBlock Block

	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Timestamp = t.String()
	newBlock.Difficulty = dificulty
	for i := 0; ; i++ {
		newBlock.Nonce = fmt.Sprintf("%x", i)
		newBlock.Hash = calculateHash(newBlock)
		if !isHashValid(newBlock.Hash, newBlock.Difficulty) {
			<-time.After(time.Millisecond * 10)
		} else {
			break
		}
	}
	return newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if newBlock.Index != oldBlock.Index+1 {
		return false
	}

	if newBlock.PrevHash != oldBlock.Hash {
		return false
	}

	if newBlock.Hash != calculateHash(newBlock) {
		return false
	}

	return true
}

func isHashValid(hashed string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hashed, prefix)
}

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
