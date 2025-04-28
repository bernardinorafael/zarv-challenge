package manhattan

// Distance calcula a distância de Manhattan entre dois pontos em uma matriz
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

	// Verificamos se há exatamente dois pontos
	if len(coordinates) != 2 {
		return -1
	}

	x1 := coordinates[0][0]
	x2 := coordinates[1][0]

	y1 := coordinates[0][1]
	y2 := coordinates[1][1]

	// calcula a distância de Manhattan: |x1-x2| + |y1-y2|
	return abs(x1-x2) + abs(y1-y2)
}

// abs retorna o valor absoluto de um número inteiro
func abs(x int) int {
	// Detalhes de implementação:
	//
	// Eu usaria `math.Abs`, mas eu iria perder em legibilidade de código,
	// teria que fazer muitas conversões de `int` para `float64` e vice-versa,
	// como o desafio pede que eu retorne um valor inteiro, acabei implementando `abs`
	if x < 0 {
		return -x
	}
	return x
}
