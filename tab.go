package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewTab(text string, action func(string)) *container.TabItem {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			action("home")
		}),
	)
	// boostrap the "+" icon to make it so users to create a new tab
	tabLayout := container.New(layout.NewVBoxLayout(), toolbar)
	item := container.NewTabItem(text, tabLayout)
	return item
}
