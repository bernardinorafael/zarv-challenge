package main

import (
	"fmt"

	"github.com/bernardinorafael/zarv_challenge/internal/manhattan"
)

func main() {
	size := 10

	// Cria uma matriz de 10x10 (100 elementos)
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	matrix[0][0] = 1
	matrix[1][2] = 1

	result := manhattan.Distance(matrix)
	fmt.Println("Manhattan distance:", result)
}
