package services

import (
	"CRM-Service/internal/models"
	"CRM-Service/internal/repositories"
)

type PartnerService struct {
	partnerRep *repositories.PartnerRepository
}

func NewPartnerService(partnerRep *repositories.PartnerRepository) *PartnerService {
	return &PartnerService{partnerRep: partnerRep}
}

func (partnerService *PartnerService) GetAllPartners() ([]*models.Partner, error) {
	return partnerService.partnerRep.GetAll()
}

func (partnerService *PartnerService) GetPartner(id int) (*models.Partner, error) {
	return partnerService.partnerRep.GetById(id)
}

func (partnerService *PartnerService) CreatePartner(partner *models.Partner) error {
	return partnerService.partnerRep.Create(partner)
}

func (partnerService *PartnerService) UpdatePartner(id int, partner *models.Partner) error {
	existPartner, err := partnerService.partnerRep.GetById(partner.ID)
	if err != nil {
		return err
	}
	existPartner.Description = partner.Description
	existPartner.Name = partner.Name
	existPartner.Contacts = partner.Contacts

	return partnerService.partnerRep.Update(existPartner)
}

func (partnerService *PartnerService) DeletePartner(id int) error {
	partner, err := partnerService.partnerRep.GetById(id)
	if err != nil {
		return err
	}
	return partnerService.partnerRep.Delete(partner)
}
