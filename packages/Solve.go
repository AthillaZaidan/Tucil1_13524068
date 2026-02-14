package utils

import (
	"fmt"
)

var iteration int

func GenerateCombinations(grid [][]byte, row, col int, numQueens int, maxQueens int, queensPlacement []int, pos int) {

	if numQueens == maxQueens {

		iteration++

		if iteration % 100 == 0 {
			fmt.Printf("Iterasi ke-%d: Checking kombinasi %v\n", iteration, queensPlacement)
			PrintGrid(grid, queensPlacement, row, col)
		}
		
		if isValid(grid, queensPlacement, row, col) {
			fmt.Println("SOLUSI DIDAPATKAN")
			PrintGrid(grid, queensPlacement, row, col)
		}
		return
	}

	if pos >= row*col {
		return
	}

	// rekursif
	queensPlacement = append(queensPlacement, pos)
	GenerateCombinations(grid, row, col, numQueens+1, maxQueens, queensPlacement, pos+1)
	queensPlacement = queensPlacement[:len(queensPlacement)-1]

  GenerateCombinations(grid, row, col, numQueens, maxQueens, queensPlacement, pos+1)

}

func Bruteforce_solve(grid [][]byte, row, col int) {
	maxQueens := countRegion(grid, row, col)

	queensPlacement := make([]int, 0, maxQueens)
	GenerateCombinations(grid, row, col, 0, maxQueens, queensPlacement, 0)

	fmt.Printf("\nTotal iterasi: %d\n", iteration)
}

func PrintGrid(grid	[][]byte, queensPlacement[]int, row, col int){
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