package forms

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateOrganizationView() *fyne.Container {
	return container.New(layout.NewFormLayout(),
		widget.NewLabel("*Name"),
		widget.NewEntry(),
		widget.NewLabel("DBA"),
		widget.NewEntry(),
	)
}
