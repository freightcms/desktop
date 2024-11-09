package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/freightcms/desktop/logging"
	"github.com/freightcms/desktop/theme"
)

func OrganizationListView() *fyne.Container {
	searchSection := container.NewHBox(
		widget.NewEntry(),
		widget.NewButtonWithIcon("Search", theme.InspectionIcon(), func() {

		}),
	)
	table := widget.NewTableWithHeaders(func() (int, int) {
		return 0, 3
	}, func() fyne.CanvasObject {
		return widget.NewLabel("hello world")
	}, func(i widget.TableCellID, o fyne.CanvasObject) {
		logging.Logger.Debug("Clicked on table cell row[%d], col[%d]", i.Row, i.Col)
	})
	innerLayout := container.NewBorder(searchSection, nil, nil, nil, table)
	l := container.New(
		layout.NewPaddedLayout(),
		innerLayout,
	)
	return l
}
