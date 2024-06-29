package repositories

import (
	"CRM-Service/internal/models"
	"gorm.io/gorm"
)

type PartnerRepository struct {
	db *gorm.DB
}

func NewPartnerRepository(db *gorm.DB) *PartnerRepository {
	return &PartnerRepository{db: db}
}

func (r *PartnerRepository) GetAll() ([]*models.Partner, error) {
	var partners []*models.Partner
	if err := r.db.Find(&partners).Error; err != nil {
		return nil, err
	}
	return partners, nil
}

func (r *PartnerRepository) GetById(id int) (*models.Partner, error) {
	var partner models.Partner
	if err := r.db.First(&partner, id).Error; err != nil {
		return nil, err
	}
	return &partner, nil
}

func (r *PartnerRepository) Create(partner *models.Partner) error {
	return r.db.Create(partner).Error
}

func (r *PartnerRepository) Update(partner *models.Partner) error {
	return r.db.Save(partner).Error
}

func (r *PartnerRepository) Delete(partner *models.Partner) error {
	return r.db.Delete(partner).Error
}
