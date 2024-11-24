package organizations

import (
	"desktop/logging"
	"desktop/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewTable() fyne.CanvasObject {
	var service services.OrganizationService
	data := service.Get(nil, []string{}) 
	length := func() (rows, cols int) {
		return len(data), 4
	}
	fieldNameMap := map[int]string {
		0: "id",
		1: "name",
		2: "dba",
		3: "rollupID",
	}
	createCellTemplate := func() fyne.CanvasObject {
		return widget.NewLabel("a row")
	}
	updateCell := func(i widget.TableCellID, cell fyne.CanvasObject) {
		record := data[i.Row]
		cellData := record[fieldNameMap[i.Col]]
		cell.(*widget.Label).SetText(cellData)
	}
	tbl := widget.NewTableWithHeaders(length, createCellTemplate, updateCell)
	// 	tbl.OnSelected = func(i widget.TableCellID) {
	// 		logging.Logger.Debug("Selected row %d, column %d", i.Row, i.Col)
	// 	}
	// 	tbl.OnUnselected = func(i widget.TableCellID) {
	// 		logging.Logger.Debug("UnSelected row %d, column %d", i.Row, i.Col)
	// 	}

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

	navigationButtons := container.NewBorder(
		nil,
		nil,
		nil,
		container.NewHBox(
			widget.NewButtonWithIcon("Previous", theme.NavigateBackIcon(), func() {
				logging.Logger.Debug("Clicking Previous")
			}),
			widget.NewButtonWithIcon("Next", theme.NavigateNextIcon(), func() {
				logging.Logger.Debug("Clicking Next")
			}),
		),
		nil,
	)
	l := container.NewBorder(
		container.NewVBox(
			NewSearchForm(),
			navigationButtons,
		),
		nil,
		nil,
		nil,
		NewTable(),
	)
	return l
}
