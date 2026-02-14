package utils

func GenerateCombinations(grid [][]byte, row, col int, numQueens int, maxQueens int, queensPlacement []int, pos int){

	if numQueens == maxQueens {
		return
	} 

	if pos >= row * col {
		return
	}

	// rekursif
	queensPlacement = append(queensPlacement, pos)



}

func Bruteforce_solve(grid [][]byte, row, col int){
	maxQueens := countRegion(grid, row, col)

	queensPlacement := make([]int, 0, maxQueens)


}




func PrintSolution ()