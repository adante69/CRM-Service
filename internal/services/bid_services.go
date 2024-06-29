package services

import (
	"CRM-Service/internal/models"
	"CRM-Service/internal/repositories"
)

type BidService struct {
	bidRep *repositories.BidRepository
}

func NewBidService(rep *repositories.BidRepository) *BidService {
	return &BidService{bidRep: rep}
}

func (s *BidService) GetAllBids() ([]*models.Bid, error) {
	return s.bidRep.GetAll()
}

func (s *BidService) GetBid(id int) (*models.Bid, error) {
	return s.bidRep.GetById(id)
}

func (s *BidService) CreateBid(b *models.Bid) error {
	return s.bidRep.Create(b)
}

func (s *BidService) UpdateBid(id int, b *models.Bid) error {
	existingBid, err := s.bidRep.GetById(id)
	if err != nil {
		return err
	}

	existingBid.Amount = b.Amount
	existingBid.CreatedAt = b.CreatedAt
	existingBid.Description = b.Description

	return s.bidRep.Update(existingBid)
}

func (s *BidService) DeleteBid(id int) error {
	existingBid, err := s.bidRep.GetById(id)
	if err != nil {
		return err
	}
	return s.bidRep.Delete(existingBid)
}
