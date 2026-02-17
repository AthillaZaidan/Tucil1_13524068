package bruteforce

import (
	"fmt"
	"time"
)

var iteration int

func GenerateCombinations(grid [][]byte, row, col int, numQueens int, maxQueens int, queensPlacement []int, pos int) ([]int, bool) {

	if numQueens == maxQueens {
		iteration++
		if iteration%1000 == 0 {
			fmt.Printf("%dth Iteration\n", iteration)
			PrintGrid(grid, queensPlacement, row, col)
		}

		if isValid(grid, queensPlacement, row, col) {
			fmt.Println("\n========================================")
			fmt.Println("Solution found")
			fmt.Println("========================================")
			fmt.Printf("Combination: %v\n", queensPlacement)
			PrintGrid(grid, queensPlacement, row, col)
			fmt.Println("========================================")
			return queensPlacement, true
		}
		return nil, false
	}

	if pos >= row*col {
		return nil, false
	}

	newPlacement := append(queensPlacement, pos)
	if solution, found := GenerateCombinations(grid, row, col, numQueens+1, maxQueens, newPlacement, pos+1); found {
		return solution, true
	}

	if solution, found := GenerateCombinations(grid, row, col, numQueens, maxQueens, queensPlacement, pos+1); found {
		return solution, true
	}

	return nil, false
}

func Bruteforce_solve(grid [][]byte, row, col int) ([]int, bool) {
	iteration = 0
	maxQueens := countRegion(grid, row, col)
	queensPlacement := make([]int, 0, maxQueens)

	fmt.Println("\n========================================")
	fmt.Println("Starting BruteForce Solver")
	fmt.Println("========================================")
	fmt.Printf("Grid Size: %d x %d\n", row, col)
	fmt.Printf("Numbers of Regions: %d\n", maxQueens)
	fmt.Printf("Numbers of Queens: %d\n", maxQueens)
	fmt.Println("========================================")

	startTime := time.Now()

	solution, found := GenerateCombinations(grid, row, col, 0, maxQueens, queensPlacement, 0)

	duration := time.Since(startTime)
	milliseconds := duration.Milliseconds()
	fmt.Printf("\n========================================\n")
	if found {
		fmt.Printf("Success!, Solution found\n")
		fmt.Printf("Iterations: %d\n", iteration)
		fmt.Printf("Time: %d ms\n", milliseconds)
		fmt.Println("========================================")
		return solution, true
	} else {
		fmt.Printf("No Solution Found :<\n")
		fmt.Printf("Iterations: %d\n", iteration)
		fmt.Printf("Time: %d ms\n", milliseconds)
	}
	fmt.Println("========================================")
	return nil, false
}

func PrintGrid(grid [][]byte, queensPlacement []int, row, col int) {
	gridDup := make([][]byte, row)
	for i := 0; i < row; i++ {
		gridDup[i] = make([]byte, col)
		copy(gridDup[i], grid[i])
	}

	for i := 0; i < len(queensPlacement); i++ {
		pos := queensPlacement[i]

		rowPos := pos / col
		colPos := pos % col

		gridDup[rowPos][colPos] = '#'
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%c ", gridDup[i][j])
		}
		fmt.Println()
	}
}
