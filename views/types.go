package views

type AppNavigationOptions int

const (
	HomeNavigationOption          AppNavigationOptions = 1 << iota
	LoginNavigationOption                              = 2
	LogoutNavigationOption                             = 4
	SettingsNavigationOption                           = 8
	OrganizationsNavigationOption                      = 16
	AccountSettingsOption                              = 32
)
