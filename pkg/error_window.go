package pkg

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func ErrorWindow(errorText string, app fyne.App) fyne.Window {
	w := app.NewWindow("Error")
	w.SetContent(widget.NewLabel(errorText))

	return w
}
