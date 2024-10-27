package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	_ "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/freightcms/desktop/logging"
)

func NewHomePage() *fyne.Container {
	return container.NewPadded(
		container.NewCenter(
			container.NewGridWithColumns(4,
				container.NewPadded(
					widget.NewButtonWithIcon("Organizations", theme.HomeIcon(), func() {
						logging.Logger.Debug("Click on settings")
					}),
					widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), func() {
						logging.Logger.Debug("Click on settings")
					}),
				),
			),
		),
	)
}
