package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/freightcms/desktop/theme"
)

func newButtonWithIcon(text string, icon fyne.Resource, tapped func(), alignment widget.ButtonAlign) *widget.Button {
	btn := widget.NewButtonWithIcon(text, icon, tapped)

	btn.Alignment = alignment

	return btn
}

func Navbar(onNavigation func(selected AppNavigationOptions)) *fyne.Container {
	buttons := []fyne.CanvasObject{
		newButtonWithIcon("Account", theme.LoginAsUserIcon(), func() {
			onNavigation(AccountSettingsOption)
		}, widget.ButtonAlignLeading),
		widget.NewAccordion(
			widget.NewAccordionItem("Organizations",
				newButtonWithIcon("Organizations", theme.OrganizationIcon(), func() {
					onNavigation(OrganizationsNavigationOption)
				}, widget.ButtonAlignLeading),
			),
		),
		newButtonWithIcon("Settings", theme.SettingsIcon(), func() {
			onNavigation(SettingsNavigationOption)
		}, widget.ButtonAlignLeading),
	}

	return container.New(
		layout.NewPaddedLayout(),
		container.NewVBox(buttons...),
	)
}
