package utils

func isValid(grid [][]byte, queensPlacement[]int, row, col int) bool {
	for i := 0; i < len(queensPlacement); i++ {
		for j := i+1 ; j < len(queensPlacement); j++ {
			pos1 := queensPlacement[i]
			pos2 := queensPlacement[j]

			row1 := pos1 / col
			col1 := pos1 % col

			row2 := pos2 / col
			col2 := pos2 % col

			if row1 == row2 {
				return false
			}

			if col1 == col2 {
				return false
			}

			if row1/3 == row2/3 && col1/3 == col2/3 {
				return false
			}

			if grid[row1][col1] == grid[row2][col2] {
				return  false
			}
		}
	}
	return  true
}
