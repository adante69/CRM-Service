package services

import (
	"CRM-Service/internal/models"
	"CRM-Service/internal/repositories"
)

type ContactService struct {
	contactRep *repositories.ContactRepository
}

func NewContactService(contactRep *repositories.ContactRepository) *ContactService {
	return &ContactService{contactRep: contactRep}
}

func (c *ContactService) GetAllContacts() ([]models.Contact, error) {
	return c.contactRep.GetAll()
}

func (c *ContactService) GetContactById(id int) (*models.Contact, error) {
	return c.contactRep.GetById(id)
}

func (c *ContactService) CreateContact(contact *models.Contact) error {
	return c.contactRep.Create(contact)
}

func (c *ContactService) UpdateContact(id int, contact *models.Contact) error {
	existContact, err := c.GetContactById(id)
	if err != nil {
		return err
	}
	existContact.Name = contact.Name
	existContact.Phone = contact.Phone
	existContact.Description = contact.Description
	return c.contactRep.Update(existContact)
}

func (c *ContactService) DeleteContact(id int) error {
	return c.contactRep.Delete(id)
}
