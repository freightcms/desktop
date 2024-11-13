package organizations

import (
	"desktop/logging"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewTable() fyne.CanvasObject {
	tbl := widget.NewTableWithHeaders(func() (rows, cols int) {
		return 1, 4
	}, func() fyne.CanvasObject {
		return widget.NewLabel("Hello World")
	}, func(i widget.TableCellID, o fyne.CanvasObject) {
		o.(*widget.Label).SetText("Hello form the other side")
	})

	return tbl
}

func NewSearchForm() fyne.CanvasObject {
	searchText := binding.NewString()
	searchEntry := widget.NewEntry()
	searchEntry.Bind(searchText)

	searchSection := widget.NewForm(
		widget.NewFormItem("Search", searchEntry),
	)
	searchSection.SubmitText = "Search"
	searchSection.OnSubmit = func() {
		text, _ := searchText.Get()
		logging.Logger.Debug("Searching %s", text)
	}

	searchSection.OnCancel = func() {
		// TODO: cancel the search
		logging.Logger.Debug("Canceling search")
	}

	return searchSection
}

// ListView creates a new view where all the organizations appear in a list and can be searched
func ListView() *fyne.Container {
	form := NewSearchForm()
	table := NewTable()
	innerLayout := container.NewBorder(form, nil, nil, nil, table)
	l := container.New(
		layout.NewPaddedLayout(),
		innerLayout,
	)
	return l
}
