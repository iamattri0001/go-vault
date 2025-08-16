package jwt

import (
	"fmt"
	customerrors "go-vault/custom_errors"
	"go-vault/database/models"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var jwtSecret = []byte("your_secret_key")

type Claims struct {
	Username string    `json:"username"`
	ID       uuid.UUID `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(user *models.User, expirationInHours int) (string, error) {
	claims := Claims{
		Username: user.Username,
		ID:       user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expirationInHours) * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Panicf("Failed to generate token: %v", err)
		return "", customerrors.ErrSomethingWentWrong
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
