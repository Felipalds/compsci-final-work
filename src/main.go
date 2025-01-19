package main

import (
	"fmt"
	"github.com/Felipalds/compsci-final-work/src/annealing"
	"github.com/Felipalds/compsci-final-work/src/particle"
	"sync"
	"time"

	"github.com/Felipalds/compsci-final-work/src/genetic"
	"github.com/Felipalds/compsci-final-work/src/helpers"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	fmt.Println("Start of the jorney...")

	matrix := helpers.ReadCsv("./data/formatted/dantzig.csv")

	go func() {
		var start time.Time
		var duration time.Duration
		start = time.Now()
		defer wg.Done()
		genetic.Execute(matrix, len(matrix))
		duration = time.Since(start)
		fmt.Println(duration)
	}()

	go func() {
		var start time.Time
		var duration time.Duration
		defer wg.Done()
		start = time.Now()
		particle.Execute(matrix, len(matrix))
		duration = time.Since(start)
		fmt.Println(duration)
	}()

	go func() {
		var start time.Time
		var duration time.Duration
		defer wg.Done()
		start = time.Now()
		annealing.Execute(matrix, len(matrix))
		duration = time.Since(start)
		fmt.Println(duration)
	}()

	wg.Wait()
	fmt.Println("Completed")

}
