package manhattan

// Distance calculates the Manhattan distance between two points in a matrix
func Distance(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return -1
	}

	var coordinates [][2]int
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == 1 {
				coordinates = append(coordinates, [2]int{row, col})
			}
		}
	}

	// Checks if there are exactly two points
	if len(coordinates) != 2 {
		return -1
	}

	x1 := coordinates[0][0]
	x2 := coordinates[1][0]

	y1 := coordinates[0][1]
	y2 := coordinates[1][1]

	// Calculates the Manhattan distance: |x1-x2| + |y1-y2|
	return abs(x1-x2) + abs(y1-y2)
}

// abs returns the absolute value of an integer
func abs(x int) int {
	// Details of implementation:
	//
	// I would use `math.Abs`, but I would lose in code readability,
	// I would have to make many conversions from `int` to `float64` and vice-versa,
	// as the challenge asks me to return an integer value, I ended up implementing `abs`
	if x < 0 {
		return -x
	}
	return x
}
