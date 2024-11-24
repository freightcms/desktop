package services

type (
	Organization struct {
		ID       string 
		Name     string
		DBA      string
		RollupID string
	}
	UpsertOrganization struct {
		Name     string
		DBA      *string
		RollupID *string
	}

	OrganizationService interface {
		// Get fetches all the organizations as transformed information that can
		// be displayed in the lists or tables
		Get(criteria interface{}, fields []string) []map[string]string
		// Delete removes an organization from the list of items
		Delete(id string) error
		Update(id string, org Organization) error
		// Get returns the 
		Create(org UpsertOrganization) (id string)
	}
)
