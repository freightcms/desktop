package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewHomePage(onNavigation func(selected AppNavigationOptions)) *fyne.Container {
	return container.NewPadded(
		container.NewCenter(
			container.NewGridWithColumns(4,
				container.NewPadded(
					widget.NewButtonWithIcon("Account", theme.AccountIcon(), func() {
						onNavigation(SettingsNavigationOption)
					}),
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
