package repositories

import (
	"CRM-Service/internal/models"
	"gorm.io/gorm"
)

type ContactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) GetAll() ([]models.Contact, error) {
	var contacts []models.Contact
	if err := r.db.Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}

func (r *ContactRepository) GetById(id int) (*models.Contact, error) {
	var contact models.Contact
	if err := r.db.First(&contact, id).Error; err != nil {
		return nil, err
	}
	return &contact, nil
}

func (r *ContactRepository) Create(contact *models.Contact) error {
	return r.db.Create(contact).Error
}

func (r *ContactRepository) Update(contact *models.Contact) error {
	return r.db.Save(contact).Error
}

func (r *ContactRepository) Delete(id int) error {
	return r.db.Delete(&models.Contact{}, id).Error
}
