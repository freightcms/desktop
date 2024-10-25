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

func NewTab(text string, action func(string)) *container.TabItem {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			action("home")
		}),
	)
	// boostrap the "+" icon to make it so users to create a new tab
	tabLayout := container.New(layout.NewVBoxLayout(), toolbar)
	item := container.NewTabItem(text, tabLayout)
	return item
}

func closeTab(tabContainer *container.AppTabs, index int) {
	logging.Logger.Debug("Closing Tab")
	tabContainer.RemoveIndex(index)
	logging.Logger.Debug("Selected index", tabContainer.SelectedIndex())
}

// appendTab handles the scenario for when user is clicking on the "new" tab item to add a new tab to their application
//  1. create a new tab that needs to be added to the list of tabs
//  2. take all the existing tab items and grab them all execpt the last one that is assumed to be the special tab
//     which allows users to create new tabs
//  3. Add the special "new tab" special case
//  4. Select the newly created tab for the user
func appendTab(tabContainer *container.AppTabs) {
	logging.Logger.Debug("Appending Tab")
	// TODO; we want to insert the new tab rather than just select it
	tabToAdd := container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Getting Started"))
	tabContainer.Append(tabToAdd)
	tabContainer.SelectIndex(len(tabContainer.Items) - 1)
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
			// TODO: do something
		}),
	)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.LoginIcon(), func() {
			logging.Logger.Debug("User Clicked Login. Not Implemented.")
		}),
	)
	main := container.New(layout.NewVBoxLayout(), toolbar, widget.NewSeparator(), tabContainer)

	//
	// setup short cuts
	// these need to be after the fact because we reference the unboxed variable tabContaine
	w.Canvas().AddShortcut(newTabShortcut, func(shortcut fyne.Shortcut) {
		logging.Logger.Debug("New Tab Shortcut")
		appendTab(tabContainer)
	})
	w.Canvas().AddShortcut(closeTabShortcut, func(shortcut fyne.Shortcut) {
		logging.Logger.Debug("Close Tab Shortcut")
		closeTab(tabContainer, tabContainer.SelectedIndex())
		items := tabContainer.Items
		if len(items) == 0 {
			logging.Logger.Debug("Last tab closed. Closing Application")
			w.Close()
		}
	})

	logging.Logger.Debug("Setting Main Content")
	w.SetContent(main)

	logging.Logger.Debug("Resizing Application to 1024 x 768")
	w.Resize(fyne.NewSize(1024, 768))

	logging.Logger.Debug("Centering Application")
	w.CenterOnScreen()

	logging.Logger.Debug("Setting Application to Visibile")
	w.Show()

	a.Run()
}
