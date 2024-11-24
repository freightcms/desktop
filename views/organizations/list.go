package organizations

import (
	"desktop/clients/web"
	"desktop/logging"
	"desktop/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewTable() fyne.CanvasObject {
	orgClient := web.NewOrganizationService()
	data := orgClient.Get(nil, []string{}) 
	length := func() (rows, cols int) {
		return len(data), 4
	}
	createCellTemplate := func() fyne.CanvasObject {
		return container.New(
			layout.NewCenterLayout(),
			nil,
		)
	}
	updateCell := func(i widget.TableCellID, cell fyne.CanvasObject) {
		record := data[i.Row]
		var label *widget.Label
		switch i.Col {
			case 0:
				label = widget.NewLabel(record.ID)
				cell.(*fyne.Container).Add(label)
				break
			case 1:
				label = widget.NewLabel(record.Name)
				cell.(*fyne.Container).Add(label)
				break 
			case 2:
				if record.DBA == nil { 
					label = widget.NewLabel("")
					cell.(*fyne.Container).Add(label)
				} else {
					label = widget.NewLabel(*record.DBA)
					cell.(*fyne.Container).Add(label)
				}
				break
			case 3:
				if record.RollupID == nil { 
					label = widget.NewLabel("")
					cell.(*fyne.Container).Add(label)
				} else {
					label = widget.NewLabel(*record.RollupID)
					cell.(*fyne.Container).Add(label)
				}
				cell.(*fyne.Container).Add(label)
				break
			case 4:
				actionButtons := container.NewHBox(
					widget.NewButtonWithIcon("Edit", theme.EditIcon(), func() {
						logging.Logger.Debug("Clicked edit button in row")
					}), 
					widget.NewButtonWithIcon("Delete", theme.RemoveIcon(), func() {
						logging.Logger.Debug("Delete icon clicked")
					}),
				)
				cell.(*fyne.Container).Add(actionButtons)
				break
			default:
				break

		}
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
			widget.NewButtonWithIcon("Previous", theme.RetateLeftIcon(), func() {
				logging.Logger.Debug("Clicking Previous")
			}),
			widget.NewButtonWithIcon("Next", theme.RetateRightIcon(), func() {
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
