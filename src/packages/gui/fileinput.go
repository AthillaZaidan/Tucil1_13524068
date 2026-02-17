package gui

import (
	imageprocessor "Tucil1/src/packages/imageprocessor"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// showFileInput — layar input nama file (mode 1 = txt, mode 2 = image)
func (qa *QueensApp) showFileInput(mode int) {
	var labelText string
	var placeholder string

	if mode == 1 {
		labelText = "Insert Text File Name:"
		placeholder = "Filename"
	} else {
		labelText = "Insert Image File Name:"
		placeholder = "Filename"
	}

	entry := widget.NewEntry()
	entry.SetPlaceHolder(placeholder)
	entryBox := container.NewGridWrap(fyne.NewSize(400, 36), entry)

	statusLabel := widget.NewLabel("")
	statusLabel.Wrapping = fyne.TextWrapWord

	loadBtn := widget.NewButton("Load File", func() {
		filename := entry.Text
		if filename == "" {
			statusLabel.SetText("Error: filename cannot be empty!")
			return
		}

		var err error
		if mode == 1 {
			err = qa.loadTextFile(filename)
		} else {
			err = qa.loadImageFile(filename)
		}

		if err != nil {
			statusLabel.SetText(fmt.Sprintf("Error: %v", err))
			return
		}

		qa.showGridView()
	})

	backBtn := widget.NewButton("<- Back", func() {
		qa.showMainMenu()
	})

	content := container.NewVBox(
		container.NewHBox(backBtn),
		layout.NewSpacer(),
		container.NewCenter(widget.NewLabel(labelText)),
		container.NewCenter(entryBox),
		container.NewCenter(loadBtn),
		container.NewCenter(statusLabel),
		layout.NewSpacer(),
	)

	qa.window.SetContent(content)
}

// loadTextFile — baca file .txt dan parse jadi grid byte
func (qa *QueensApp) loadTextFile(filename string) error {
	data, err := os.ReadFile("../test/" + filename)
	if err != nil {
		return err
	}

	// cari banyak kolom (sampai newline pertama)
	col := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' || data[i] == '\r' {
			break
		}
		col++
	}

	// cari banyak baris
	row := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			row++
		}
	}
	if len(data) > 0 && data[len(data)-1] != '\n' {
		row++
	}

	// parse grid byte per byte
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

	// simpan ke struct
	qa.grid = grid
	qa.row = row
	qa.col = col
	qa.lastSolution = nil

	// validasi grid harus NxN
	if row != col {
		return fmt.Errorf("Grid must be square (NxN). Current grid is %dx%d", row, col)
	}

	// backup original grid
	qa.originalGrid = make([][]byte, row)
	for i := 0; i < row; i++ {
		qa.originalGrid[i] = make([]byte, col)
		for j := 0; j < col; j++ {
			qa.originalGrid[i][j] = grid[i][j]
		}
	}

	return nil
}

// loadImageFile — baca file image dan convert ke grid byte via imageprocessor
func (qa *QueensApp) loadImageFile(filename string) error {
	img, err := imageprocessor.LoadImage("../test/" + filename)
	if err != nil {
		return err
	}

	cellSize := imageprocessor.DetectCellSize(img)
	grid, rows, cols := imageprocessor.ImageToGrid(img, cellSize)

	// simpan ke struct
	qa.grid = grid
	qa.row = rows
	qa.col = cols
	qa.lastSolution = nil

	// validasi grid harus NxN
	if rows != cols {
		return fmt.Errorf("Grid must be square (NxN). Current grid is %dx%d", rows, cols)
	}

	// backup original grid
	qa.originalGrid = make([][]byte, rows)
	for i := 0; i < rows; i++ {
		qa.originalGrid[i] = make([]byte, cols)
		for j := 0; j < cols; j++ {
			qa.originalGrid[i][j] = grid[i][j]
		}
	}

	return nil
}
