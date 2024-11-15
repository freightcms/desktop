package organizations

import (
	"desktop/logging"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type OrganizationItemViewModel struct {
}

type OrganizationTableViewModel struct {
	page     int
	pageSize int
	sortBy   string
	items    []*OrganizationItemViewModel
}

func NewOrganizationTableViewModel() *OrganizationTableViewModel {
	return &OrganizationTableViewModel{
		page:     0,
		pageSize: 10,
		sortBy:   "id",
		items:    []*OrganizationItemViewModel{},
	}
}

// Length gets the number of records available
func (v *OrganizationTableViewModel) Length() int {
	return len(v.items)
}

// At gets the organization view item at the index specified. If the index is out of bounds
// of the array then the function returns `nil`, and `false`. If the index somehow fetches
// a `nil` value from the items then it will return `nil` and `false` as the item was
// present in the array but stores as `nil`
func (v *OrganizationTableViewModel) At(i int) (*OrganizationItemViewModel, bool) {
	if i < 0 || i > len(v.items) {
		return nil, false
	}
	item := v.items[i]
	return item, item != nil
}

func NewTable(vm *OrganizationTableViewModel) fyne.CanvasObject {
	headers := []string{"ID", "Name", "DBA", "Rollup ID"}
	tbl := widget.NewTableWithHeaders(func() (rows, cols int) {
		return vm.Length(), 4
	}, func() fyne.CanvasObject {
		return widget.NewLabel("Hello World")
	}, func(i widget.TableCellID, o fyne.CanvasObject) {
		if i.Row == 0 {
			o.(*widget.Label).SetText(headers[i.Col])
			return
		}
		data, ok := vm.At(i.Row)
		if ok {
			o.(*widget.Label).SetText(fmt.Sprintf("Row %d, Col %d"))
		}
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
