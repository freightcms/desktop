package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/freightcms/desktop/logging"
)

func closeTab(tabContainer *container.AppTabs, index int) {
	items := tabContainer.Items
	selectedIndex := tabContainer.SelectedIndex()
	if index >= len(items) {
		return
	}
	if index == len(items)-1 {
		items = items[:index]
	} else if index == 0 {
		items = items[1:]
	} else {
		items = append(items, items[:index]...)
		items = append(items, items[index+1:]...)
	}

	tabContainer.SetItems(items)
	// we need to figure out where the user was in the list of tabs. It's possible that there are no tabs that can be selected by the user at this point
	if selectedIndex == 0 && len(items) == 0 {
		if len(items) == 0 {
			// user has closed all tabs we should now close the window
			tabContainer.SelectIndex(0)
		} else {
			// there are more items but we can set the selected index still to 0
			tabContainer.SelectIndex(0)
		}
	} else if selectedIndex >= len(items) {
		// the user was selected on the last index so lets go to the left
		tabContainer.SelectIndex(selectedIndex - 1)
	} else {
		// we can just select whatever tab was to the left the user
		tabContainer.SelectIndex(selectedIndex + 1)
	}
}

// appendTab handles the scenario for when user is clicking on the "new" tab item to add a new tab to their application
//  1. create a new tab that needs to be added to the list of tabs
//  2. take all the existing tab items and grab them all execpt the last one that is assumed to be the special tab
//     which allows users to create new tabs
//  3. Add the special "new tab" special case
//  4. Select the newly created tab for the user
func appendTab(tabContainer *container.AppTabs) {
	// TODO; we want to insert the new tab rather than just select it
	tabToAdd := container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Getting Started"))
	items := tabContainer.Items
	items = append(items, tabToAdd)
	tabContainer.SetItems(items)
	tabContainer.SelectIndex(len(items) - 1) // select the new tab for the user
}

var (
	newTabShortcut = &desktop.CustomShortcut{
		KeyName:  fyne.KeyN,
		Modifier: fyne.KeyModifierControl,
	}
	closeTabShortcut = &desktop.CustomShortcut{
		KeyName:  fyne.KeyW,
		Modifier: fyne.KeyModifierControl,
	}
)

func main() {
	a := app.New()
	w := a.NewWindow("Freight CMS")

	tabContainer := container.NewAppTabs(
		NewTab("Home", func(tabName string) {
			logging.debugLogger.Printf("Tab name = %s\n", tabName)
		}),
	)

	tabContainer.OnSelected = func(item *container.TabItem) {
		if item.Text == "New" {
			appendTab(tabContainer)
		}
	}

	main := container.New(layout.NewStackLayout(), tabContainer)

	w.Canvas().AddShortcut(newTabShortcut, func(shortcut fyne.Shortcut) {
		appendTab(tabContainer)
	})
	w.Canvas().AddShortcut(closeTabShortcut, func(shortcut fyne.Shortcut) {
		closeTab(tabContainer, tabContainer.SelectedIndex())
		items := tabContainer.Items
		if len(items) == 0 {
			w.Close()
		}
	})

	w.SetContent(main)
	w.Resize(fyne.NewSize(1024, 768))
	w.CenterOnScreen()
	w.Show()

	a.Run()
}
