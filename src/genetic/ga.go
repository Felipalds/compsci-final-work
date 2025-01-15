package genetic

import (
	"fmt"
	"math/rand"

	"github.com/Felipalds/compsci-final-work/src/helpers"
)

// starting the genetic!
//

const (
	populationSize = 100
	generations    = 500
	mutationRate   = 0.1
)

func generateInitialPopulation(size, numCities int) [][]int {
	population := make([][]int, size)
	for i := 0; i < size; i++ {
		route := rand.Perm(numCities) // TODO pesquisar o que é esse Perm
		population[i] = route
	}
	return population
}

func selection(population [][]int, matrix [][]float64) []int { // TODO estudar essa função
	best := population[rand.Intn(len(population))]
	for i := 0; i < 3; i++ {
		contender := population[rand.Intn(len(population))]
		if helpers.CalculateDistance(contender, matrix) < helpers.CalculateDistance(best, matrix) {
			best = contender
		}
	}
	return best
}

func crossover(parent1, parent2 []int) []int { // TODO estudar essa função, aqui talvez mudaremos a impl
	size := len(parent1)
	start, end := rand.Intn(size), rand.Intn(size) // TODO pesquisar o que é o Intn
	if start > end {
		start, end = end, start
	}
	child := make([]int, size)
	visited := make(map[int]bool)

	for i := start; i <= end; i++ {
		child[i] = parent1[i]
		visited[parent1[i]] = true
	}

	index := 0
	for _, gene := range parent2 {
		if !visited[gene] {
			for index >= start && index <= end {
				index++
			}
			child[index] = gene
			index++
		}
	}

	return child
}

func mutate(route []int) {
	if rand.Float64() < mutationRate {
		i, j := rand.Intn(len(route)), rand.Intn(len(route))
		route[i], route[j] = route[j], route[i]
	}
}

func run(matrix [][]float64, numCities int) ([]int, float64) {
	population := generateInitialPopulation(populationSize, numCities)
	for generation := 0; generation < generations; generation++ {
		// fmt.Println("Generation ", generation)
		newPopulation := make([][]int, 0, populationSize)
		for i := 0; i < populationSize; i++ {
			parent1 := selection(population, matrix)
			parent2 := selection(population, matrix)
			child := crossover(parent1, parent2)
			mutate(child)
			newPopulation = append(newPopulation, child)
		}
		population = newPopulation

		// bestCurrentRoute := helpers.FindBestRoute(population, matrix)
		// bestCurrentDistance := helpers.CalculateDistance(bestCurrentRoute, matrix)
		// fmt.Println("Best Pop:", bestCurrentRoute)
		// fmt.Println("Best Dist:", bestCurrentDistance)
	}

	bestRoute := helpers.FindBestRoute(population, matrix)
	bestDistance := helpers.CalculateDistance(bestRoute, matrix)

	return bestRoute, bestDistance
}

func Execute(matrix [][]float64, numCities int) {

	fmt.Println("Starting the Genetics . . .")
	bestRoute, bestDistance := run(matrix, int(numCities))
	fmt.Printf("Melhor rota: %v\n", bestRoute)
	fmt.Printf("Distância total: %.2f\n", bestDistance)
}
