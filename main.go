package main

import (
	bruteforce "Tucil1/packages/bruteforce"
	bruteforceoptimized "Tucil1/packages/bruteforce-optimized"
	imageprocessor "Tucil1/packages/imageprocessor"
	output "Tucil1/packages/output"
	utils "Tucil1/packages/utils"
	"fmt"
	"log"
	"os"
)

func main() {
	var inputMode int
	inputMode = 0

	for inputMode != 3 {
		fmt.Println("\n===================================")
		fmt.Println("Choose Input Mode")
		fmt.Println("1. Text File (.txt)")
		fmt.Println("2. Image File (.jpg/.png)")
		fmt.Println("3. Exit Program")
		fmt.Println("===================================")
		fmt.Scan(&inputMode)

		if inputMode == 3 {
			fmt.Println("Exiting Program...")
			break
		}

		if inputMode != 1 && inputMode != 2 {
			fmt.Println("Invalid choice! Please choose 1, 2, or 3.")
			continue
		}

		var grid [][]byte
		var row, col int
		var originalGrid [][]byte

		if inputMode == 2 {
			var namaFile string
			fmt.Print("Insert Image File Name: ")
			fmt.Scan(&namaFile)

			img, err := imageprocessor.LoadImage(namaFile)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Success reading image: %s\n", namaFile)
			cellSize := imageprocessor.DetectCellSize(img)
			fmt.Printf("Auto-detected cell size: %d pixels\n", cellSize)

			grid, row, col = imageprocessor.ImageToGrid(img, cellSize)

			fmt.Printf("Detected Grid Size: %d x %d\n", row, col)
			fmt.Println("\nDetected Grid:")
			for i := 0; i < row; i++ {
				for j := 0; j < col; j++ {
					fmt.Printf("%c ", grid[i][j])
				}
				fmt.Println()
			}

			// backup original grid
			originalGrid = make([][]byte, row)
			for i := 0; i < row; i++ {
				originalGrid[i] = make([]byte, col)
				for j := 0; j < col; j++ {
					originalGrid[i][j] = grid[i][j]
				}
			}

		} else {
			var namaFile string
			fmt.Print("Insert File Name: ")
			fmt.Scan(&namaFile)

			data, err := os.ReadFile(namaFile)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Success reading file: %s\n", namaFile)

			// cari banyak kolom
			col = 0
			for i := 0; i < len(data); i++ {
				if data[i] == '\n' || data[i] == '\r' {
					break
				}
				col++
			}
			fmt.Printf("Numbers of Columns: %d\n", col)

			row = 0
			for i := 0; i < len(data); i++ {
				if data[i] == '\n' {
					row++
				}
			}
			if len(data) > 0 && data[len(data)-1] != '\n' {
				row++
			}
			fmt.Printf("Numbers of Rows: %d\n", row)

			grid = make([][]byte, 0, row)
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

			// backup original grid
			originalGrid = make([][]byte, row)
			for i := 0; i < row; i++ {
				originalGrid[i] = make([]byte, col)
				for j := 0; j < col; j++ {
					originalGrid[i][j] = grid[i][j]
				}
			}
		}
		var pilihan int
		pilihan = 0

		for pilihan != 3 {
			utils.PrintMenu()
			fmt.Scan(&pilihan)

			if pilihan == 3 {
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
}
