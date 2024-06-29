package repositories

import (
	"CRM-Service/internal/models"
	"gorm.io/gorm"
)

type BidRepository struct {
	db *gorm.DB
}

func NewBidRepository(db *gorm.DB) *BidRepository {
	return &BidRepository{db: db}
}

func (r *BidRepository) GetAll() ([]*models.Bid, error) {
	var bids []*models.Bid
	err := r.db.Find(&bids).Error
	if err != nil {
		return nil, err
	}
	return bids, nil
}

func (r *BidRepository) GetById(id int) (*models.Bid, error) {
	var bid *models.Bid
	err := r.db.First(&bid, id).Error
	if err != nil {
		return bid, err
	}
	return bid, nil
}

func (r *BidRepository) Create(bid *models.Bid) error {
	return r.db.Create(&bid).Error
}

func (r *BidRepository) Update(bid *models.Bid) error {
	return r.db.Save(&bid).Error
}

func (r *BidRepository) Delete(bid *models.Bid) error {
	return r.db.Delete(&bid).Error
}
