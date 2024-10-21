package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewHomePage() *fyne.Container {
	homePageLayout := layout.NewGridLayoutWithColumns(4)
	orgWidget := widget.NewButtonWithIcon("Organizations", theme.HomeIcon(), func() {
		debugLogger.Printf("Clicked on the Organization Tile\n")
	})
	return container.New(
		homePageLayout,
		orgWidget,
	)
}
