package mongodb

import (
	"desktop/services"
)

type (
	organizationService struct {

	}
)

func NewOrganizationService() services.OrganizationService {
	return &organizationService{}	
}

func (s *organizationService) Get(criteria interface{}, fields []string) []map[string]string {
	return []map[string]string{}
}

// Delete removes an organization from the list of items
func (s *organizationService) Delete(id string) error {
	return nil
}

func (s *organizationService) Update(id string, org services.Organization) error {
	return nil
}

// Get returns the 
func (s *organizationService) Create(org services.UpsertOrganization) (id string) {
	return ""
}
