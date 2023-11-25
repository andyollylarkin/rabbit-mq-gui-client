package pkg

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/streadway/amqp"
)

func CreateProducerForm(app fyne.App) fyne.CanvasObject {
	exchangeNameInput := &widget.Entry{}
	uriInput := &widget.Entry{Text: "amqp://guest:guest@0.0.0.0:5672"}
	rkInput := &widget.Entry{}
	var componentsList = []string{}
	componentsTree := widget.NewList(
		func() int {
			return len(componentsList)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(componentsList[i]) // i need to update this when componentsList was updated
		},
	)
	componentsTree.Resize(fyne.NewSize(1920, 1080))

	var amqpChan *amqp.Channel

	inputWidget := CreateSendMessageInput(func(textMessage string) {
		msg := amqp.Publishing{
			ContentType: "plain/text",
			Body:        []byte(textMessage),
		}
		err := amqpChan.Publish(exchangeNameInput.Text, rkInput.Text, false, false, msg)
		if err != nil {
			errw := ErrorWindow(err.Error(), app)
			errw.Show()
		} else {
			t := time.Now().Format(time.DateTime)
			componentsList = append(componentsList, fmt.Sprintf("%s\t%s", t, textMessage))
			componentsTree.Refresh()
			log.Println("Successfully published")
		}
		return
	})
	inputWidget.Hide()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Exchange name", Widget: exchangeNameInput},
			{Text: "URI", Widget: uriInput},
			{Text: "Routing key", Widget: rkInput},
			{Text: "Input", Widget: inputWidget},
			{Text: "Messages", Widget: componentsTree},
		},
		SubmitText: "Connect",
		OnSubmit: func() { // optional, handle form submission
			ch, err := CreateProducer(uriInput.Text)
			if err != nil {
				errW := ErrorWindow(err.Error(), app)
				errW.Show()
				return
			}
			amqpChan = ch
			inputWidget.Show()
		},
	}

	return form
}
