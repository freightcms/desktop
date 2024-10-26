package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	_ "github.com/freightcms/desktop/logging"
)

func NewHomePage() *fyne.Container {
	homePageLayout := layout.NewGridLayoutWithColumns(4)
	organizationCard := widget.NewCard("Organizations", "Manage", widget.NewIcon(theme.SettingsIcon()))

	return container.New(
		homePageLayout,
		container.NewCenter(
			container.New(homePageLayout, organizationCard)),
	)
}
