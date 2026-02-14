package main

import (
	"Tucil1/packages/utils"
  "core/main/packages"
	"fmt"
	"log"
	"os"
)

func main() {
	var namaFile string
	fmt.Print("Masukkan nama file: ")
	fmt.Scan(&namaFile)

	data, err := os.ReadFile(namaFile)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Berhasil membaca file\n%s", namaFile)

	// cari banyak kolom
	col := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' || data[i] == '\r' {
			break
		}
		col++
	}
	fmt.Printf("Panjang col: %d\n", col)

	row := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			row++
		}
	}
	if len(data) > 0 && data[len(data)-1] != '\n' {
		row++
	}
	fmt.Printf("Panjang Row: %d\n", row)

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
	
	fmt.Printf("Ukuran grid: %d x %d\n", len(grid), col)
	
	gridDup := grid
}
