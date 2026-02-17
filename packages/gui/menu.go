package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// showMainMenu — tampilkan menu utama (pilih input mode)
func (qa *QueensApp) showMainMenu() {
	title := canvas.NewText("Queens Puzzle Solver", color.White)
	title.TextSize = 32
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	subtitle := canvas.NewText("Brute Force Algorithm", color.NRGBA{0xBB, 0xBB, 0xBB, 0xFF})
	subtitle.TextSize = 16
	subtitle.Alignment = fyne.TextAlignCenter

	txtBtn := widget.NewButton("Text File (.txt)", func() {
		qa.showFileInput(1)
	})
	imgBtn := widget.NewButton("Image File (.jpg/.png)", func() {
		qa.showFileInput(2)
	})
	exitBtn := widget.NewButton("Exit Program", func() {
		qa.fyneApp.Quit()
	})

	content := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(title),
		container.NewCenter(subtitle),
		layout.NewSpacer(),
		container.NewCenter(container.NewVBox(
			txtBtn,
			imgBtn,
			widget.NewSeparator(),
			exitBtn,
		)),
		layout.NewSpacer(),
	)

	qa.window.SetContent(content)
}
