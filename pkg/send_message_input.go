package pkg

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateSendMessageInput(onSubmit func(messageText string)) fyne.CanvasObject {
	input := &widget.Entry{}
	submitButton := &widget.Button{Text: "Send message"}
	submitButton.OnTapped = func() {
		text := input.Text
		onSubmit(text)
	}
	w := container.New(layout.NewGridLayout(2), input, submitButton)

	return w
}
