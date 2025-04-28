package main

import (
	"fmt"

	"github.com/bernardinorafael/zarv_challenge/internal/manhattan"
)

func main() {
	size := 10

	// Create a 10x10 matrix (100 elements)
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	// Define the coordinates (0, 0) and (1, 2)
	// To change the coordinates, just change the values of `x` and `y`
	matrix[0][0] = 1
	matrix[1][2] = 1

	result := manhattan.Distance(matrix)
	fmt.Println("Manhattan distance:", result)
}
