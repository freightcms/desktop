package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Freight CMS")
	links := []string{"Home", "About", "Contact"}
	navigationList := widget.NewList(func() int {
		return len(links)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("A template")
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		object.(*widget.Label).SetText(links[id])
	})

	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", container.NewStack(
			widget.NewForm(
				widget.NewFormItem("Name", widget.NewEntry()),
				widget.NewFormItem("DBA", widget.NewEntry()),
			), // end NewForm
		)),
	)
	main := container.New(layout.NewGridLayout(2), navigationList, tabs)

	w.SetContent(main)
	w.Resize(fyne.NewSize(1024, 768))
	w.Show()

	a.Run()
}
