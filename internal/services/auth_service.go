package services

import (
	"CRM-Service/internal/auth"
	"CRM-Service/internal/models"
	"CRM-Service/internal/repositories"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	accountRepository *repositories.AccountRepository
}

func NewAuthService(accountRepository *repositories.AccountRepository) *AuthService {
	return &AuthService{accountRepository: accountRepository}
}

func (s *AuthService) Register(email, password string) error {
	_, err := s.accountRepository.FindByEmail(email)
	if err == nil {
		return errors.New("user already exists")
	}

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	account := models.Account{
		Email:    email,
		Password: string(hashedpassword),
	}
	return s.accountRepository.Create(account)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.accountRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}
	token, err := auth.GenerateJWT(email)
	if err != nil {
		return "", err
	}
	return token, nil
}
