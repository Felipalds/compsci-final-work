package particle

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/Felipalds/compsci-final-work/src/helpers"
)

const (
	numParticles    = 50
	maxIterations   = 1000
	inertiaWeight   = 0.5
	cognitiveFactor = 2.0
	socialFactor    = 2.0
)

func initializeParticles(numParticles, numCities int) ([][]int, [][]float64) {
	particles := make([][]int, numParticles)
	velocities := make([][]float64, numParticles)

	for i := 0; i < numParticles; i++ {
		particles[i] = rand.Perm(numCities) // TODO estudar o que é esse Perm
		velocities[i] = make([]float64, numCities)
	}
	return particles, velocities
}

func updateParticles(particles [][]int, velocities [][]float64, personalBests [][]int, globalBest []int, matrix [][]float64) {
	for i := range particles {
		for j := range velocities[i] {
			cognitive := cognitiveFactor * rand.Float64() * (float64(personalBests[i][j]) - float64(particles[i][j]))
			social := socialFactor * rand.Float64() * (float64(globalBest[j]) - float64(particles[i][j]))
			velocities[i][j] = inertiaWeight*velocities[i][j] + cognitive + social

		}
		rand.Shuffle(len(particles[i]), func(a, b int) {
			if rand.Float64() < math.Abs(velocities[i][a]) {
				particles[i][a], particles[i][b] = particles[i][b], particles[i][a]
			}
		})
	}
}

func run(matrix [][]float64, numCities int) ([]int, float64) {
	particles, velocities := initializeParticles(numParticles, numCities)
	personalBests := make([][]int, numParticles)
	copy(personalBests, particles)

	personalBestScores := make([]float64, numParticles)
	for i, particle := range particles {
		personalBestScores[i] = helpers.CalculateDistance(particle, matrix)
	}

	globalBest := personalBests[0]
	globalBestScore := personalBestScores[0]

	for i := 1; i < numParticles; i++ {
		if personalBestScores[i] < globalBestScore {
			globalBest = personalBests[i]
			globalBestScore = personalBestScores[i]
		}
	}

	for iteration := 0; iteration < maxIterations; iteration++ {
		updateParticles(particles, velocities, personalBests, globalBest, matrix)

		for i, particle := range particles {
			score := helpers.CalculateDistance(particle, matrix)
			if score < personalBestScores[i] {
				personalBests[i] = particle
				personalBestScores[i] = score
			}
			if score < globalBestScore {
				globalBest = particle
				globalBestScore = score
			}
		}
	}
	return globalBest, globalBestScore
}

func Execute(matrix [][]float64, numCities int) {

	fmt.Println("Starting the Particle Swarm . . .")

	globalBest, globalBestScore := run(matrix, int(numCities))
	fmt.Printf("Melhor rota: %v\n", globalBest)
	fmt.Printf("Distância total: %.2f\n", globalBestScore)
}
