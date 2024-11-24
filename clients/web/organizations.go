package web 

import (
	"desktop/clients"
)

type (
	organizationService struct {

	}
)

func NewOrganizationService() clients.OrganizationService {
	return &organizationService{}	
}

func (s *organizationService) Get(criteria interface{}, fields []string) []clients.Organization {
	return []clients.Organization{}
}

// Delete removes an organization from the list of items
func (s *organizationService) Delete(id string) error {
	return nil
}

func (s *organizationService) Update(id string, org clients.UpsertOrganization) error {
	return nil
}

// Get returns the 
func (s *organizationService) Create(org clients.UpsertOrganization) (id string) {
	return ""
}
