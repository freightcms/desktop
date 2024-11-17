package services

type (
	Organization struct {
		ID       interface{}
		Name     string
		DBA      string
		RollupID string
	}
	CreateOrganization struct {
		Name     string
		DBA      *string
		RollupID *string
	}

	OrganizationService interface {
		Get(criteria interface{}, fields []string)
		Delete(id interface{})
		Update(id interface{}, org Organization) error
		Create(org CreateOrganization) (id interface{})
	}
)
