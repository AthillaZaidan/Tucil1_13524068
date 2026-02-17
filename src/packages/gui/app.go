package gui

import (
	"image/color"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// QueensApp — struct utama yang hold semua state GUI
type QueensApp struct {
	fyneApp fyne.App
	window  fyne.Window

	// data grid
	grid         [][]byte
	originalGrid [][]byte
	row          int
	col          int

	// elemen visual grid
	cellRects  [][]*canvas.Rectangle
	cellQueens [][]*canvas.Text
	gridBox    *fyne.Container

	// label status
	statusLabel *widget.Label
	iterLabel   *widget.Label
	timeLabel   *widget.Label

	// kontrol solver
	solving      bool
	stopFlag     bool
	mu           sync.Mutex
	stepDelay    time.Duration
	lastSolution []int

	// shared state buat solver callback (atomic-like via mutex)
	latestPlacement []int
	latestIter      int
}

// Run — entry point utama GUI, dipanggil dari main.go
func Run() {
	qa := &QueensApp{
		stepDelay: 50 * time.Millisecond,
	}

	qa.fyneApp = app.New()
	qa.window = qa.fyneApp.NewWindow("Queens Puzzle Solver")
	qa.window.Resize(fyne.NewSize(900, 750))
	qa.window.CenterOnScreen()

	qa.showMainMenu()
	qa.window.ShowAndRun()
}

// regionColors — 26 warna buat region A-Z (mirip LinkedIn Queens game)
var regionColors = [26]color.NRGBA{
	{0xE8, 0x6B, 0x6B, 0xFF}, // A - Red
	{0xCD, 0xDC, 0x39, 0xFF}, // B - Lime
	{0x64, 0xB5, 0xF6, 0xFF}, // C - Blue
	{0x81, 0xC7, 0x84, 0xFF}, // D - Green
	{0xCE, 0x93, 0xD8, 0xFF}, // E - Purple
	{0xFF, 0xB7, 0x4D, 0xFF}, // F - Orange
	{0xF4, 0x8F, 0xB1, 0xFF}, // G - Pink
	{0xB0, 0xBE, 0xC5, 0xFF}, // H - Gray
	{0xFF, 0xD5, 0x4F, 0xFF}, // I - Gold
	{0x4D, 0xD0, 0xE1, 0xFF}, // J - Cyan
	{0xA1, 0x88, 0x7F, 0xFF}, // K - Brown
	{0x90, 0xCA, 0xF9, 0xFF}, // L - Light Blue
	{0xBA, 0x68, 0xC8, 0xFF}, // M - Violet
	{0xFF, 0xCC, 0x80, 0xFF}, // N - Light Orange
	{0x80, 0xDE, 0xEA, 0xFF}, // O - Light Cyan
	{0xEF, 0x9A, 0x9A, 0xFF}, // P - Light Red
	{0xE6, 0xEE, 0x9C, 0xFF}, // Q - Light Lime
	{0xA5, 0xD6, 0xA7, 0xFF}, // R - Light Green
	{0xBC, 0xAA, 0xA4, 0xFF}, // S - Tan
	{0xFF, 0xAB, 0x91, 0xFF}, // T - Deep Orange
	{0xB3, 0x9D, 0xDB, 0xFF}, // U - Lavender
	{0x9F, 0xA8, 0xDA, 0xFF}, // V - Indigo
	{0xFF, 0x80, 0xAB, 0xFF}, // W - Hot Pink
	{0xEA, 0x80, 0xFC, 0xFF}, // X - Magenta
	{0x84, 0xFF, 0xFF, 0xFF}, // Y - Aqua
	{0xCC, 0xFF, 0x90, 0xFF}, // Z - Lime Green
}
