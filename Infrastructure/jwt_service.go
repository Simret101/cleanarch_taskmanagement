package Infrastructure

import (
	"errors"
	"task/Domain"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateJWT(username string) (string, error)

	ValidateToken(tokenString string) (*Domain.Claims, error)
}

type jwtService struct {
	SecretKey       string
	TokenExpiration time.Duration
}

func NewJWTService(secretKey string, expiration time.Duration) JWTService {
	return &jwtService{
		SecretKey:       secretKey,
		TokenExpiration: expiration,
	}
}

func (j *jwtService) GenerateJWT(username string) (string, error) {

	expirationTime := time.Now().Add(j.TokenExpiration)

	claims := &Domain.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.SecretKey))
}

func (j *jwtService) ValidateToken(tokenString string) (*Domain.Claims, error) {

	claims := &Domain.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
