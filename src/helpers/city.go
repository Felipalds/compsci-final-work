package helpers

func CalculateDistance(route []int, matrix [][]float64) float64 {
	total := 0.0
	for i := 0; i < len(route)-1; i++ {
		total += matrix[route[i]][route[i+1]]
	}
	total += matrix[route[len(route)-1]][route[0]] // return to first city
	return total
}

func FindBestRoute(population [][]int, matrix [][]float64) []int {
	best := population[0]
	bestDistance := CalculateDistance(best, matrix)
	for _, route := range population {
		distance := CalculateDistance(route, matrix)
		if distance < bestDistance {
			best = route
			bestDistance = distance
		}
	}
	return best

}
