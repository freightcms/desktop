package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	_ "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type AppNavigationOptions int

const (
	HomeNavigationOption          AppNavigationOptions = 1 << iota
	LoginNavigationOption                              = 2
	LogoutNavigationOption                             = 4
	SettingsNavigationOption                           = 8
	OrganizationsNavigationOption                      = 16
	AccountSettingsOption                              = 32
)

func NewHomePage(onNavigation func(selected AppNavigationOptions)) *fyne.Container {
	return container.NewPadded(
		container.NewCenter(
			container.NewGridWithColumns(4,
				container.NewPadded(
					widget.NewButtonWithIcon("Organizations", theme.HomeIcon(), func() {
						onNavigation(OrganizationsNavigationOption)
					}),
					widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), func() {
						onNavigation(SettingsNavigationOption)
					}),
				),
			),
		),
	)
}
