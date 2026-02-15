package bruteforceoptimized

import (
	bruteforce "Tucil1/packages/bruteforce"
	"fmt"
	"time"
)

var (
	iteration int
)

func HasNonQueenRegion(grid [][]byte, row, col int) bool {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] >= 'A' && grid[i][j] <= 'Z' {
				return true
			}
		}
	}
	return false
}

func restoreGrid(grid [][]byte, backup [][]byte, row, col int) {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			grid[i][j] = backup[i][j]
		}

	}
}

func MarkQueenArea(grid [][]byte, row, col, qRow, qCol int) {
	for i := 0; i < col; i++ {
		if grid[qRow][i] != '#' {
			grid[qRow][i] = '.'
		}
	}

	for i := 0; i < row; i++ {
		if grid[i][qCol] != '#' {
			grid[i][qCol] = '.'
		}
	}

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			r := qRow + i
			c := qCol + j

			if r >= 0 && r < row && c >= 0 && c < col {
				if grid[r][c] != '#' {
					grid[r][c] = '.'
				}
			}
		}

	}

	grid[qRow][qCol] = '#'
}

func FindSmallestRegion(grid [][]byte, row, col int) byte {
	var color = make([]int, 26)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] >= 'A' && grid[i][j] <= 'Z' {
				color[grid[i][j]-'A']++
			}
		}
	}

	min := row*col + 1
	minPos := byte(0)

	for i := 0; i < 26; i++ {
		if color[i] < min && color[i] > 0 {
			minPos = byte('A' + i)
			min = color[i]
		}
	}

	return minPos
}

func backupGrid(grid [][]byte, row, col int) [][]byte {
	backup := make([][]byte, row)
	for i := 0; i < row; i++ {
		backup[i] = make([]byte, col)
		for j := 0; j < col; j++ {
			backup[i][j] = grid[i][j]
		}
	}
	return backup
}

func SolveSmallestRegion(grid [][]byte, originalGrid [][]byte, row, col int, queensPlacement []int) ([]int, bool) {

	iteration++
	if iteration%5 == 0 {
		fmt.Printf("%dth Iterationt\n", iteration)
		bruteforce.PrintGrid(originalGrid, queensPlacement, row, col)
	}

	// IZIN KING
	// cari region terkecil
	// taro queen disitu
	// cek is valid
	// kalo valid gas lagi

	if !HasNonQueenRegion(grid, row, col) {
		return queensPlacement, true
	}

	// cari dulu targetRegion
	targetRegion := FindSmallestRegion(grid, row, col)
	if targetRegion == 0 {
		return nil, false
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] != targetRegion {
				continue
			}

			backup := backupGrid(grid, row, col)

			// nah ini proses utamanya queen taro"
			MarkQueenArea(grid, row, col, i, j)

			for r := 0; r < row; r++ {
				for c := 0; c < col; c++ {
					if grid[r][c] == targetRegion {
						grid[r][c] = '.'
					}
				}
			}

			newPlacement := append(queensPlacement, i*col+j)

			if resultPlacement, found := SolveSmallestRegion(grid, originalGrid, row, col, newPlacement); found {
				return resultPlacement, true
			}

			restoreGrid(grid, backup, row, col)
		}
	}
	return nil, false
}

func Bruteforce_optimized_solve(grid [][]byte, row, col int) {
	iteration = 0
	queensPlacement := make([]int, 0)
	originalGrid := make([][]byte, row)
	for i := 0; i < row; i++ {
		originalGrid[i] = make([]byte, col)
		copy(originalGrid[i], grid[i])
	}

	fmt.Println("\n========================================")
	fmt.Println("Starting Optimized BruteForce Solver")
	fmt.Println("========================================")
	fmt.Printf("Grid Size: %d x %d\n", row, col)
	fmt.Println("========================================")

	startTime := time.Now()

	finalPlacement, found := SolveSmallestRegion(grid, originalGrid, row, col, queensPlacement)

	duration := time.Since(startTime)
	milliseconds := duration.Milliseconds()
	fmt.Printf("\n========================================\n")
	if found {
		fmt.Printf("Success!, Solution found\n")
		fmt.Printf("Iterations (Recursive Steps): %d\n", iteration)
		fmt.Printf("Time: %d ms\n", milliseconds)
		bruteforce.PrintGrid(originalGrid, finalPlacement, row, col)
	} else {
		fmt.Printf("No Solution Found :<\n")
		fmt.Printf("Iterations: %d\n", iteration)
		fmt.Printf("Time: %d ms\n", milliseconds)
	}
	fmt.Println("========================================")
}
