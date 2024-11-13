package organizations

import (
	"desktop/logging"
	"desktop/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Table() fyne.CanvasObject {
	tbl := widget.NewTableWithHeaders(func() (rows, cols int) {
		return 0, 0
	}, func() fyne.CanvasObject {
		return widget.NewLabel("Hello World")
	}, func(i widget.TableCellID, o fyne.CanvasObject) {
		o.(*widget.Label).SetText("Hello form the other side")
	})

	return tbl
}

// ListView creates a new view where all the organizations appear in a list and can be searched
func ListView() *fyne.Container {
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
