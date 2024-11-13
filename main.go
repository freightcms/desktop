package main

import (
	"fmt"
	"runtime"

	"desktop/logging"
	"desktop/theme"
	"desktop/views"
	"desktop/views/organizations"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func handleNavigation(page views.AppNavigationOptions) {
	if page == views.OrganizationsNavigationOption {
		item := tabContainer.Selected()
		item.Text = "Organizations"
		item.Icon = theme.OrganizationIcon()
		item.Content = organizations.ListView()
	} else {
		logging.Logger.Debug("Navigation to %d", page)
	}
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
	tabToAdd := container.NewTabItemWithIcon("Home", theme.HomeIcon(), views.WelcomeView())
	tabContainer.Append(tabToAdd)
	tabContainer.SelectIndex(len(tabContainer.Items) - 1)
}

var (
	desktopApp     fyne.App
	tabContainer   *container.AppTabs
	newTabShortcut = &desktop.CustomShortcut{
		KeyName:  fyne.KeyN,
		Modifier: fyne.KeyModifierControl,
	}
	closeTabShortcut = &desktop.CustomShortcut{
		KeyName:  fyne.KeyW,
		Modifier: fyne.KeyModifierControl,
	}
	toolbar = widget.NewToolbar(
		widget.NewToolbarAction(theme.LoginAsUserIcon(), func() {
			logging.Logger.Debug("User Clicked Account. Not Implemented.")
		}),
	)
	navbar     = views.Navbar(handleNavigation)
	mainWindow fyne.Window
)

func main() {
	desktopApp = app.New()
	mainWindow := desktopApp.NewWindow("Freight CMS")

	tabContainer = container.NewAppTabs()
	appendTab(tabContainer)

	appbar := container.NewBorder(nil, nil, nil, toolbar, nil)
	detailsContainer := container.New(layout.NewCenterLayout(),
		widget.NewLabel(fmt.Sprintf("Operating System: %s", runtime.GOOS)))
	body := container.NewBorder(appbar, detailsContainer, navbar, nil, tabContainer)

	//
	// setup short cuts
	// these need to be after the fact because we reference the unboxed variable tabContaine
	mainWindow.Canvas().AddShortcut(newTabShortcut, func(shortcut fyne.Shortcut) {
		logging.Logger.Debug("New Tab Shortcut")
		appendTab(tabContainer)
	})
	mainWindow.Canvas().AddShortcut(closeTabShortcut, func(shortcut fyne.Shortcut) {
		logging.Logger.Debug("Close Tab Shortcut")
		closeTab(tabContainer, tabContainer.SelectedIndex())
		items := tabContainer.Items
		if len(items) == 0 {
			logging.Logger.Debug("Last tab closed. Closing Application")
			mainWindow.Close()
		}
	})

	logging.Logger.Debug("Setting Main Content")
	mainWindow.SetContent(body)

	logging.Logger.Debug("Resizing Application to 1024 x 768")
	mainWindow.Resize(fyne.NewSize(1024, 768))

	logging.Logger.Debug("Centering Application")
	mainWindow.CenterOnScreen()

	logging.Logger.Debug("Setting Application to Visibile")
	mainWindow.Show()

	desktopApp.Run()
}
