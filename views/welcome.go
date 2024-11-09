package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func welcomeLabel() fyne.CanvasObject {
	style := fyne.TextStyle{
		Bold: true,
	}
	txt := &canvas.Text{
		Alignment: fyne.TextAlignCenter,
		Text:      "Welcome!",
		TextSize:  36,
		TextStyle: style,
	}
	return txt
}

func WelcomeView() *fyne.Container {
	return container.New(
		layout.NewCenterLayout(),
		welcomeLabel(),
	)
}
