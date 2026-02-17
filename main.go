package main

import (
	bruteforce "Tucil1/packages/bruteforce"
	bruteforceoptimized "Tucil1/packages/bruteforce-optimized"
	output "Tucil1/packages/output"
	utils "Tucil1/packages/utils"
	"fmt"
	"log"
	"os"
)

func main() {
	var namaFile string
	fmt.Print("Insert File Name: ")
	fmt.Scan(&namaFile)

	data, err := os.ReadFile(namaFile)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Success reading file: %s", namaFile)

	// cari banyak kolom
	col := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' || data[i] == '\r' {
			break
		}
		col++
	}
	fmt.Printf("Numbers of Columns: %d\n", col)

	row := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			row++
		}
	}
	if len(data) > 0 && data[len(data)-1] != '\n' {
		row++
	}
	fmt.Printf("Numbers of Rows: %d\n", row)

	grid := make([][]byte, 0, row)
	currentRow := make([]byte, 0, col)

	for i := 0; i < len(data); i++ {
		if data[i] == '\r' {
			continue
		}

		if data[i] == '\n' {
			grid = append(grid, currentRow)
			currentRow = make([]byte, 0, col)
			continue
		}

		currentRow = append(currentRow, data[i])
	}

	if len(data) > 0 && data[len(data)-1] != '\n' {
		grid = append(grid, currentRow)
	}

	fmt.Printf("Grid Size: %d x %d\n", len(grid), col)

	originalGrid := make([][]byte, row)
	for i := 0; i < row; i++ {
		originalGrid[i] = make([]byte, col)
		copy(originalGrid[i], grid[i])
	}

	var pilihan int
	pilihan = 0

	for pilihan != 3 {
		utils.PrintMenu()
		fmt.Scan(&pilihan)

		if pilihan == 3 {
			fmt.Println("Exiting Program...")
			break
		}

		if pilihan != 1 && pilihan != 2 {
			fmt.Println("Invalid choice! Please choose 1, 2, or 3.")
			continue
		}

		var solution []int
		var found bool

		if pilihan == 1 {
			solution, found = bruteforce.Bruteforce_solve(grid, row, col)
		} else {
			solution, found = bruteforceoptimized.Bruteforce_optimized_solve(grid, row, col)
		}

		if found {
			var saveAnswer string
			fmt.Print("\nDo you want to save the solution to file? (y/n): ")
			fmt.Scan(&saveAnswer)

			if saveAnswer == "y" || saveAnswer == "Y" {
				var filename string
				fmt.Print("Enter filename (without extension): ")
				fmt.Scan(&filename)

				if filename == "" {
					filename = "solution"
				}

				err := output.SaveToTxt(filename+".txt", originalGrid, solution, row, col)
				if err != nil {
					fmt.Printf("Error saving file: %v\n", err)
				} else {
					fmt.Printf("Solution saved to %s.txt\n", filename)
				}
			}
		}

		fmt.Println()
	}
}
