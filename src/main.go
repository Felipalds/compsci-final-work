package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/Felipalds/compsci-final-work/src/brute"
)

func main() {
	fmt.Println("Start of the jorney...")

	// get the data
	// duas analises: 10x1 10x2 10x3 - para um tempo médio
	// dificuldade dinamica, qual dificuldade ele ficou mais
	// descartar a primeira (cache)
	// define the structs
	// create the algorithms
	// define the test cases
	// test and get statistics

	data := "Hello, world"

	target := big.NewInt(1)
	target.Lsh(target, 240) // This means 2^240, making it a relatively easy target

	pow := brute.Block{Data: data, Target: target}
	start := time.Now()

	pow.Mine()

	duration := time.Since(start)

	fmt.Println("Completed mining")
	fmt.Printf("Data: %s\nNonce: %d\nHash: %s\nTime Taken: %s\n", pow.Data, pow.Nonce, pow.Hash, duration)

}
