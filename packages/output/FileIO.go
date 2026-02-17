package output

import (
	"fmt"
	"os"
)

func SaveToTxt(filename string, originalGrid [][]byte, queensPlacement []int, row, col int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// kalo gak ada solusi ya langsung return aja
	if queensPlacement == nil {
		file.WriteString("No solution found.\n")
		return nil
	}

	// copy grid original biar gak keubah
	displayGrid := make([][]byte, row)
	for i := 0; i < row; i++ {
		displayGrid[i] = make([]byte, col)
		for j := 0; j < col; j++ {
			displayGrid[i][j] = originalGrid[i][j]
		}
	}

	// taro queen di grid
	for i := 0; i < len(queensPlacement); i++ {
		pos := queensPlacement[i]
		rowPos := pos / col
		colPos := pos % col

		if rowPos < row && colPos < col {
			displayGrid[rowPos][colPos] = '#'
		}
	}

	// tulis ke file
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			file.WriteString(fmt.Sprintf("%c ", displayGrid[i][j]))
		}
		file.WriteString("\n")
	}

	return nil
}
