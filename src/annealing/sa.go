package annealing

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/Felipalds/compsci-final-work/src/helpers"
)

const (
	initialTemperature  = 1000.0
	coolingRate         = 0.995
	stoppingTemperature = 1e-6
)

func generateNeighbor(route []int) []int {
	neighbor := make([]int, len(route))
	copy(neighbor, route)
	i, j := rand.Intn(len(route)), rand.Intn(len(route))
	neighbor[i], neighbor[j] = neighbor[j], neighbor[i]
	return neighbor
}

func run(matrix [][]float64, numCities int) ([]int, float64) {
	rand.Seed(time.Now().UnixNano())
	currentSolution := rand.Perm(numCities)
	currentDistance := helpers.CalculateDistance(currentSolution, matrix)

	temperature := initialTemperature
	bestSolution := make([]int, len(currentSolution))
	copy(bestSolution, currentSolution)
	bestDistance := currentDistance

	for temperature > stoppingTemperature {
		neighbor := generateNeighbor(currentSolution)
		neighborDistance := helpers.CalculateDistance(neighbor, matrix)

		if neighborDistance < currentDistance || math.Exp((currentDistance-neighborDistance)/temperature) > rand.Float64() {
			currentSolution = neighbor
			currentDistance = neighborDistance
		}

		if currentDistance < bestDistance {
			copy(bestSolution, currentSolution)
			bestDistance = currentDistance
		}

		temperature *= coolingRate
	}
	return bestSolution, bestDistance
}

func Execute(matrix [][]float64, numCities int) {

	fmt.Println("Starting the Annealing . . .")
	bestRoute, bestDistance := run(matrix, int(numCities))
	fmt.Printf("Melhor rota: %v\n", bestRoute)
	fmt.Printf("Distância total: %.2f\n", bestDistance)
}
