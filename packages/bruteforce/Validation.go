package bruteforce

func isValid(grid [][]byte, queensPlacement []int, row, col int) bool {
	for i := 0; i < len(queensPlacement); i++ {
		for j := i + 1; j < len(queensPlacement); j++ {
			pos1 := queensPlacement[i]
			pos2 := queensPlacement[j]

			row1 := pos1 / col
			col1 := pos1 % col

			row2 := pos2 / col
			col2 := pos2 % col

			// Check: sama row
			if row1 == row2 {
				return false
			}

			// Check: sama column
			if col1 == col2 {
				return false
			}

			// Check: ADJACENT (sebelahan) - ga boleh dalam jarak 1 cell (8 arah)
			rowDiff := row1 - row2
			if rowDiff < 0 {
				rowDiff = -rowDiff
			}
			colDiff := col1 - col2
			if colDiff < 0 {
				colDiff = -colDiff
			}
			if rowDiff <= 1 && colDiff <= 1 {
				return false
			}

			// Check: sama region
			if grid[row1][col1] == grid[row2][col2] {
				return false
			}
		}
	}
	return true
}
