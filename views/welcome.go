package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func WelcomeView() *fyne.Container {
	return container.New(
		layout.NewCenterLayout(),
		widget.NewLabel("Welcome!"),
	)
}
