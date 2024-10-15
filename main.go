package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Freight CMS")

	content := widget.NewLabel("A label")
	w.SetContent(container.NewVBox(
		content,
		widget.NewButton("Hi!", func() {
			content.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
