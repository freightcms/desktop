package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/freightcms/desktop/theme"
)

func NewHomePage(onNavigation func(selected AppNavigationOptions)) *fyne.Container {
	return container.NewPadded(
		container.NewCenter(
			container.NewGridWithColumns(4,
				widget.NewButtonWithIcon("Account", theme.LoginAsUserIcon(), func() {
					onNavigation(SettingsNavigationOption)
				}),
				widget.NewButtonWithIcon("Organizations", theme.OrganizationIcon(), func() {
					onNavigation(OrganizationsNavigationOption)
				}),
				widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), func() {
					onNavigation(SettingsNavigationOption)
				}),
			),
		),
	)
}
