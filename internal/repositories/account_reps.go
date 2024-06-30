package repositories

import (
	"CRM-Service/internal/models"
	"errors"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Create(user models.Account) error {
	return r.db.Create(&user).Error
}

func (r *AccountRepository) FindByEmail(email string) (*models.Account, error) {
	var account models.Account
	if err := r.db.Where("email = ?", email).First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("account not found")
		}
		return nil, err
	}
	return &account, nil
}
