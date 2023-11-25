package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/andyollylarkin/rabbit-mq-gui-client/pkg"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("RabbitMQ client")

	tabs := container.NewAppTabs(
		container.NewTabItem("Producer", pkg.CreateProducerForm(myApp)),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
