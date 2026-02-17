package bruteforceoptimized

import (
	bruteforce "Tucil1/packages/bruteforce"
	"fmt"
	"time"
)

var (
	iteration int
	OnStep    func([]int, int)
	StopFlag  *bool
)

func isPositionValid(qRow, qCol int, queensPlacement []int, numQueens int, originalGrid [][]byte, col int) bool {
	for i := 0; i < numQueens; i++ {
		pos := queensPlacement[i]
		r := pos / col
		c := pos % col

		// same row
		if r == qRow {
			return false
		}

		// same column
		if c == qCol {
			return false
		}

		// adjacent (king's move)
		rowDiff := qRow - r
		if rowDiff < 0 {
			rowDiff = -rowDiff
		}
		colDiff := qCol - c
		if colDiff < 0 {
			colDiff = -colDiff
		}
		if rowDiff <= 1 && colDiff <= 1 {
			return false
		}

		// same region
		if originalGrid[r][c] == originalGrid[qRow][qCol] {
			return false
		}
	}
	return true
}

// Cari region yang belum solved dengan jumlah available cell paling sedikit (MRV heuristic)
// Return 0 kalo ada region unsolved yang ga punya available cell (dead end / pruning)
func FindSmallestUnsolvedRegion(originalGrid [][]byte, row, col int, queensPlacement []int, numQueens int, solvedRegions [26]bool) byte {
	var availableCount [26]int
	var totalCount [26]int

	for i := 0; i < 26; i++ {
		availableCount[i] = 0
		totalCount[i] = 0
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			ch := originalGrid[i][j]
			if ch >= 'A' && ch <= 'Z' && !solvedRegions[ch-'A'] {
				totalCount[ch-'A']++
				if isPositionValid(i, j, queensPlacement, numQueens, originalGrid, col) {
					availableCount[ch-'A']++
				}
			}
		}
	}

	// Dead end check: ada region yang masih butuh queen tapi ga ada cell yang valid
	for i := 0; i < 26; i++ {
		if totalCount[i] > 0 && availableCount[i] == 0 {
			return 0
		}
	}

	// Cari region dengan available cell paling sedikit
	min := row*col + 1
	var minRegion byte = 0
	for i := 0; i < 26; i++ {
		if availableCount[i] > 0 && availableCount[i] < min {
			min = availableCount[i]
			minRegion = byte('A' + i)
		}
	}

	return minRegion
}

func countTotalRegions(grid [][]byte, row, col int) int {
	var color [26]bool
	for i := 0; i < 26; i++ {
		color[i] = false
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			ch := grid[i][j]
			if ch >= 'A' && ch <= 'Z' {
				color[ch-'A'] = true
			}
		}
	}

	regionNum := 0
	for i := 0; i < 26; i++ {
		if color[i] {
			regionNum++
		}
	}
	return regionNum
}

func SolveSmallestRegion(originalGrid [][]byte, row, col int, queensPlacement []int, numQueens int, solvedRegions [26]bool, totalRegions int) ([]int, bool) {

	if StopFlag != nil && *StopFlag {
		return nil, false
	}

	iteration++
	if OnStep != nil {
		// optimized punya sedikit iterasi, callback tiap kali ok
		OnStep(queensPlacement, iteration)
	} else if iteration%5 == 0 {
		fmt.Printf("%dth Iterations\n", iteration)
		bruteforce.PrintGrid(originalGrid, queensPlacement, row, col)
	}

	if numQueens == totalRegions {
		return queensPlacement, true
	}

	targetRegion := FindSmallestUnsolvedRegion(originalGrid, row, col, queensPlacement, numQueens, solvedRegions)
	if targetRegion == 0 {
		return nil, false
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if originalGrid[i][j] != targetRegion {
				continue
			}

			if !isPositionValid(i, j, queensPlacement, numQueens, originalGrid, col) {
				continue
			}

			newPlacement := make([]int, numQueens+1)
			for k := 0; k < numQueens; k++ {
				newPlacement[k] = queensPlacement[k]
			}
			newPlacement[numQueens] = i*col + j
			solvedRegions[targetRegion-'A'] = true

			// rekursi
			if result, found := SolveSmallestRegion(originalGrid, row, col, newPlacement, numQueens+1, solvedRegions, totalRegions); found {
				return result, true
			}

			// backtrack: un-solve region
			solvedRegions[targetRegion-'A'] = false
		}
	}

	return nil, false
}

func Bruteforce_optimized_solve(grid [][]byte, row, col int) ([]int, bool) {
	iteration = 0
	queensPlacement := make([]int, 0)

	// copy grid supaya ga modify yang asli
	originalGrid := make([][]byte, row)
	for i := 0; i < row; i++ {
		originalGrid[i] = make([]byte, col)
		copy(originalGrid[i], grid[i])
	}

	totalRegions := countTotalRegions(originalGrid, row, col)
	var solvedRegions [26]bool
	for i := 0; i < 26; i++ {
		solvedRegions[i] = false
	}

	fmt.Println("\n========================================")
	fmt.Println("Starting Optimized BruteForce Solver")
	fmt.Println("========================================")
	fmt.Printf("Grid Size: %d x %d\n", row, col)
	fmt.Printf("Numbers of Regions: %d\n", totalRegions)
	fmt.Println("========================================")

	startTime := time.Now()

	finalPlacement, found := SolveSmallestRegion(originalGrid, row, col, queensPlacement, 0, solvedRegions, totalRegions)

	duration := time.Since(startTime)
	milliseconds := duration.Milliseconds()
	fmt.Printf("\n========================================\n")
	if found {
		fmt.Printf("Success!, Solution found\n")
		fmt.Printf("Iterations (Recursive Steps): %d\n", iteration)
		fmt.Printf("Time: %d ms\n", milliseconds)
		bruteforce.PrintGrid(originalGrid, finalPlacement, row, col)
		fmt.Println("========================================")
		return finalPlacement, true
	} else {
		fmt.Printf("No Solution Found :<\n")
		fmt.Printf("Iterations: %d\n", iteration)
		fmt.Printf("Time: %d ms\n", milliseconds)
	}
	fmt.Println("========================================")
	return nil, false
}
