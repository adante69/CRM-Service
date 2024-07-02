package auth

import (
	"CRM-Service/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
type JWT struct {
	conf *config.Configuration
}

func GenerateJWT(cf *config.Configuration, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(cf.JWT.Secret)
}

func ValidateJWT(cf *config.Configuration, token string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) { return cf.JWT.Secret, nil })
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}
	if !tkn.Valid {
		return nil, err
	}
	return claims, nil
}
