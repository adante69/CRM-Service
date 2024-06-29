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

func GetKey() (key string) {
	conf, err := config.LoadConfiguration()
	if err != nil {
		panic(err)
	}
	secretKey := conf.JWT.Secret
	return secretKey
}
func GenerateJWT(email string) (string, error) {
	secretKey := []byte(GetKey())
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func ValidateJWT(token string) (*Claims, error) {
	secretKey := []byte(GetKey())
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) { return secretKey, nil })
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
