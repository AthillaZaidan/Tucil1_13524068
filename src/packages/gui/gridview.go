package gui

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// showGridView — layar utama: tampilkan grid + tombol solver
func (qa *QueensApp) showGridView() {
	// hitung cell size biar muat di window
	maxGridPx := float32(520)
	cellSize := maxGridPx / float32(qa.col)
	if maxGridPx/float32(qa.row) < cellSize {
		cellSize = maxGridPx / float32(qa.row)
	}
	if cellSize < 30 {
		cellSize = 30
	}
	if cellSize > 70 {
		cellSize = 70
	}

	// bikin grid visual
	qa.cellRects = make([][]*canvas.Rectangle, qa.row)
	qa.cellQueens = make([][]*canvas.Text, qa.row)
	qa.gridBox = container.NewGridWithColumns(qa.col)

	for i := 0; i < qa.row; i++ {
		qa.cellRects[i] = make([]*canvas.Rectangle, qa.col)
		qa.cellQueens[i] = make([]*canvas.Text, qa.col)

		for j := 0; j < qa.col; j++ {
			ch := qa.originalGrid[i][j]

			// ambil warna region
			var bgColor color.NRGBA
			if ch >= 'A' && ch <= 'Z' {
				bgColor = regionColors[ch-'A']
			} else {
				bgColor = color.NRGBA{0x44, 0x44, 0x44, 0xFF}
			}

			rect := canvas.NewRectangle(bgColor)
			rect.SetMinSize(fyne.NewSize(cellSize, cellSize))
			rect.StrokeColor = color.NRGBA{0x22, 0x22, 0x22, 0xFF}
			rect.StrokeWidth = 2

			queenText := canvas.NewText("", color.NRGBA{0x00, 0x00, 0x00, 0xFF})
			queenText.TextSize = cellSize * 0.55
			queenText.Alignment = fyne.TextAlignCenter

			qa.cellRects[i][j] = rect
			qa.cellQueens[i][j] = queenText

			cell := container.NewStack(rect, container.NewCenter(queenText))
			qa.gridBox.Add(cell)
		}
	}

	// info
	gridInfo := widget.NewLabel(fmt.Sprintf("Grid: %d x %d", qa.row, qa.col))

	// status labels
	qa.statusLabel = widget.NewLabel("Status: Ready")
	qa.iterLabel = widget.NewLabel("Iterations: 0")
	qa.timeLabel = widget.NewLabel("Time: 0 ms")

	// speed slider
	speedLabel := widget.NewLabel(fmt.Sprintf("Delay: %dms", int(qa.stepDelay.Milliseconds())))
	speedSlider := widget.NewSlider(0, 500)
	speedSlider.Value = float64(qa.stepDelay.Milliseconds())
	speedSlider.Step = 10
	speedSlider.OnChanged = func(val float64) {
		qa.stepDelay = time.Duration(val) * time.Millisecond
		speedLabel.SetText(fmt.Sprintf("Delay: %dms", int(val)))
	}

	// tombol solver
	pureBruteBtn := widget.NewButton("Pure Bruteforce", func() {
		qa.mu.Lock()
		if qa.solving {
			qa.mu.Unlock()
			return
		}
		qa.mu.Unlock()
		go qa.runSolver(1)
	})

	optimizedBtn := widget.NewButton("Optimized Bruteforce", func() {
		qa.mu.Lock()
		if qa.solving {
			qa.mu.Unlock()
			return
		}
		qa.mu.Unlock()
		go qa.runSolver(2)
	})

	stopBtn := widget.NewButton("Stop", func() {
		qa.mu.Lock()
		qa.stopFlag = true
		qa.mu.Unlock()
	})

	resetBtn := widget.NewButton("Reset", func() {
		qa.mu.Lock()
		if qa.solving {
			qa.mu.Unlock()
			return
		}
		qa.mu.Unlock()
		qa.clearQueens()
		qa.lastSolution = nil
		qa.statusLabel.SetText("Status: Ready")
		qa.iterLabel.SetText("Iterations: 0")
		qa.timeLabel.SetText("Time: 0 ms")
	})

	saveBtn := widget.NewButton("Save Solution", func() {
		qa.mu.Lock()
		if qa.solving || qa.lastSolution == nil {
			qa.mu.Unlock()
			return
		}
		qa.mu.Unlock()
		qa.showSaveDialog()
	})

	backBtn := widget.NewButton("<- Back to Menu", func() {
		qa.mu.Lock()
		if qa.solving {
			qa.mu.Unlock()
			return
		}
		qa.mu.Unlock()
		qa.showMainMenu()
	})

	// layout
	solverBar := container.NewHBox(
		pureBruteBtn,
		optimizedBtn,
		widget.NewSeparator(),
		stopBtn,
		resetBtn,
		saveBtn,
	)

	speedBar := container.NewBorder(nil, nil, speedLabel, nil, speedSlider)

	infoBar := container.NewHBox(
		qa.statusLabel,
		layout.NewSpacer(),
		qa.iterLabel,
		widget.NewSeparator(),
		qa.timeLabel,
	)

	content := container.NewVBox(
		container.NewHBox(backBtn, layout.NewSpacer(), gridInfo),
		widget.NewSeparator(),
		container.NewCenter(qa.gridBox),
		widget.NewSeparator(),
		solverBar,
		speedBar,
		widget.NewSeparator(),
		infoBar,
	)

	scroll := container.NewScroll(content)
	qa.window.SetContent(scroll)
}

// clearQueens — hapus semua queen dari display grid
func (qa *QueensApp) clearQueens() {
	for i := 0; i < qa.row; i++ {
		for j := 0; j < qa.col; j++ {
			qa.cellQueens[i][j].Text = ""
			qa.cellQueens[i][j].Refresh()
		}
	}
}

// updateGridQueens — update tampilan queen di grid berdasarkan placement
func (qa *QueensApp) updateGridQueens(queensPlacement []int, numQueens int) {
	// clear semua queen dulu
	for i := 0; i < qa.row; i++ {
		for j := 0; j < qa.col; j++ {
			if qa.cellQueens[i][j].Text != "" {
				qa.cellQueens[i][j].Text = ""
				qa.cellQueens[i][j].Refresh()
			}
		}
	}

	// set queen di posisi yang sesuai
	for k := 0; k < numQueens; k++ {
		pos := queensPlacement[k]
		r := pos / qa.col
		c := pos % qa.col
		if r >= 0 && r < qa.row && c >= 0 && c < qa.col {
			qa.cellQueens[r][c].Text = "\U0001F451"
			qa.cellQueens[r][c].Refresh()
		}
	}
}
