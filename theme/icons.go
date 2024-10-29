package theme

import (
	"fyne.io/fyne/v2"
	ft "fyne.io/fyne/v2/theme"
)

type CustomIconName string

const (
	SettingsIconName         CustomIconName = "./assets/icons/icons8-settings-48.png"
	UserGroupIconName        CustomIconName = "./assets/icon/icons8-add-user-group-woman-man-48.png"
	EditIconName             CustomIconName = "./assets/icon/icons-8-edit-48.png"
	IndexIconName            CustomIconName = "./assets/icon/icons-8-index-48.png"
	InspectionIconName       CustomIconName = "./assets/icon/icons-8-inspection-48.png"
	LoginAsUserIconName      CustomIconName = "./assets/icon/icons-8-login-as-user-48.png"
	MaleUserIconName         CustomIconName = "./assets/icon/icons-8-login-male-user-48.png"
	NewCompanyIconName       CustomIconName = "./assets/icon/icons-8-new-company-48.png"
	NewContactIconName       CustomIconName = "./assets/icon/icons-8-new-contact-48.png"
	OrganizationIconName     CustomIconName = "./assets/icon/icons-8-organization-48.png"
	PeopleIconName           CustomIconName = "./assets/icon/icons-8-people-48.png"
	RelatedCompaniesIconName CustomIconName = "./assets/icon/icons-8-related-companies-48.png"
	RemoveIconName           CustomIconName = "./assets/icon/icons-8-remove-48.png"
	RotateLeftIconName       CustomIconName = "./assets/icon/icons-8-rotate-left-48.png"
	RotateRightIconName      CustomIconName = "./assets/icon/icons-8-rotate-right-48.png"
	RulerIconName            CustomIconName = "./assets/icon/icons-8-ruler-48.png"
	SortingIconName          CustomIconName = "./assets/icon/icons-8-sorting-48.png"
	ThumbnailsIconName       CustomIconName = "./assets/icon/icons-8-thumbnails-48.png"
)

func resourceFromPathOrFallback(iconName CustomIconName, fallback fyne.Resource) fyne.Resource {
	r, err := fyne.LoadResourceFromPath(string(iconName))
	if err == nil {
		return r
	}
	return fallback
}

func SettingsIcon() fyne.Resource {
	return resourceFromPathOrFallback(SettingsIconName, ft.SettingsIcon())
}

func UserGroupIcon() fyne.Resource {
	return resourceFromPathOrFallback(UserGroupIconName, ft.AccountIcon())
}

func EditIcon() fyne.Resource {
	return resourceFromPathOrFallback(UserGroupIconName, ft.DocumentCreateIcon())
}

func IndexIcon() fyne.Resource {
	return resourceFromPathOrFallback(UserGroupIconName, ft.AccountIcon())
}

func InspectionIcon() fyne.Resource {
	return resourceFromPathOrFallback(InspectionIconName, ft.SearchIcon())
}

func LoginAsUserIcon() fyne.Resource {
	return resourceFromPathOrFallback(LoginAsUserIconName, ft.LoginIcon())
}

func MaleUserIcon() fyne.Resource {
	return resourceFromPathOrFallback(MaleUserIconName, ft.AccountIcon())
}

func NewCompanyIcon() fyne.Resource {
	return resourceFromPathOrFallback(NewCompanyIconName, ft.AccountIcon())
}

func NewContactIcon() fyne.Resource {
	return resourceFromPathOrFallback(NewContactIconName, ft.QuestionIcon())
}

func OrganizationIcon() fyne.Resource {
	return resourceFromPathOrFallback(OrganizationIconName, ft.QuestionIcon())
}

func PeopleIcon() fyne.Resource {
	return resourceFromPathOrFallback(PeopleIconName, ft.QuestionIcon())
}

func RelatedCompaniesIcon() fyne.Resource {
	return resourceFromPathOrFallback(RelatedCompaniesIconName, ft.QuestionIcon())
}

func RemoveIcon() fyne.Resource {
	return resourceFromPathOrFallback(RemoveIconName, ft.DeleteIcon())
}

func RetateLeftIcon() fyne.Resource {
	return resourceFromPathOrFallback(RotateLeftIconName, ft.ViewRefreshIcon())
}

func RetateRightIcon() fyne.Resource {
	return resourceFromPathOrFallback(RotateRightIconName, ft.ViewRefreshIcon())
}

func RulerIcon() fyne.Resource {
	return resourceFromPathOrFallback(RulerIconName, ft.QuestionIcon())
}

func SortingIcon() fyne.Resource {
	return resourceFromPathOrFallback(SortingIconName, ft.QuestionIcon())
}

func ThumbnailsIcon() fyne.Resource {
	return resourceFromPathOrFallback(ThumbnailsIconName, ft.QuestionIcon())
}
