package hash

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error generating hash: %v", err)
		return "", err
	}
	return string(bytes), nil
}

func CheckHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
