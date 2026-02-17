package gui

import (
	output "Tucil1/packages/output"
	"fmt"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// showSaveDialog — dialog buat save solution ke file .txt
func (qa *QueensApp) showSaveDialog() {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("solution")

	items := []*widget.FormItem{
		widget.NewFormItem("Filename", entry),
	}

	dialog.ShowForm("Save Solution", "Save", "Cancel", items, func(ok bool) {
		if !ok {
			return
		}

		filename := entry.Text
		if filename == "" {
			filename = "solution"
		}

		err := output.SaveToTxt(filename+".txt", qa.originalGrid, qa.lastSolution, qa.row, qa.col)
		if err != nil {
			dialog.ShowError(err, qa.window)
		} else {
			dialog.ShowInformation("Saved",
				fmt.Sprintf("Solution saved to %s.txt", filename),
				qa.window)
		}
	}, qa.window)
}
