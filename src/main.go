package main

import (
	"fmt"
	"time"

	"github.com/Felipalds/compsci-final-work/src/annealing"
	"github.com/Felipalds/compsci-final-work/src/genetic"
	"github.com/Felipalds/compsci-final-work/src/helpers"
	"github.com/Felipalds/compsci-final-work/src/particle"
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

	matrix := helpers.ReadCsv("./data/formatted/berlin52.csv")

	start := time.Now()
	genetic.Execute(matrix, len(matrix))
	duration := time.Since(start)
	fmt.Println(duration)

	start = time.Now()
	particle.Execute(matrix, len(matrix))
	duration = time.Since(start)
	fmt.Println(duration)

	start = time.Now()
	annealing.Execute(matrix, len(matrix))
	duration = time.Since(start)
	fmt.Println(duration)
	fmt.Println("Completed")
}
