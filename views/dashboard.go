package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/freightcms/desktop/logging"
)

func NewHomePage() *fyne.Container {
	homePageLayout := layout.NewGridLayoutWithColumns(4)
	orgWidget := widget.NewButtonWithIcon("Organizations", theme.HomeIcon(), func() {
		logging.Logger.Debug("Clicked Organization Button Tile")
	})
	return container.New(
		homePageLayout,
		orgWidget,
	)
}
