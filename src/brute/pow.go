package brute

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
)

// this code will show you how Proof of Work... works LOL

type Block struct {
	Data      string
	Target    *big.Int
	Nonce     int
	Hash      string
	Timestamp time.Time
}

func calculateHash(data string, nonce int) string {
	input := fmt.Sprintf("%s%d", data, nonce)
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

// in golang, this is like a method
func (pow *Block) Mine() {
	for {
		pow.Hash = calculateHash(pow.Data, pow.Nonce)

		// Convert the hash to a big.Int to compare it with the target
		var hashInt big.Int
		hashInt.SetString(pow.Hash, 16)

		// Check if the hash is less than the target
		if hashInt.Cmp(pow.Target) == -1 { // hash < target
			break
		}
		pow.Nonce++
	}
	pow.Timestamp = time.Now()
}
