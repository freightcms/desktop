package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func AboutPage() *fyne.Container {
	return container.New(
		layout.NewStackLayout(),
		widget.NewLabel("TODO"),
	)
}
