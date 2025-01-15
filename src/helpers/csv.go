package helpers

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ReadCsv(filePath string) [][]float64 {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll() // TODO pesquisar o que é esse ReadAll
	if err != nil {
		panic(err)
	}

	matrix := make([][]float64, len(data))
	for i, row := range data {
		matrix[i] = make([]float64, len(row))
		for j, value := range row {
			matrix[i][j], _ = strconv.ParseFloat(value, 64)
		}
	}

	fmt.Println(matrix)

	return matrix


}
