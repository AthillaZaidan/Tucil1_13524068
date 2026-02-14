package utils

func MarkQueens(grid [][]byte, rowPos, colPos int, row, col int) {
	grid[rowPos][colPos] = '#' // ini sebagai queen

	for i := 0; i < row; i++ {
		if i != rowPos {
			grid[i][colPos] = '.' // ini biar gaada queen nabrak
		} else {
			continue
		}
	}

	for i := 0; i < col; i++ {
		if i != colPos {
			grid[rowPos][i] = '.' // ini sama dengan atas
		}
	}

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			newRow := rowPos + i
			newCol := colPos + j

			// Cek apakah posisi valid dan bukan posisi queen
			if newRow >= 0 && newRow < row && newCol >= 0 && newCol < col {
				if !(i == 0 && j == 0) { // jangan overwrite queen
					grid[newRow][newCol] = '.'
				}
			}
		}
	}
}

func isSafe(grid [][]byte, rowPos, colPos int) bool {
	if grid[rowPos][colPos] == '.'{
		return false
	} else {
		return true
	}
}
