package utils

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

// HashPassword returns the bcrypt hash of the input password
// It uses bcrypt to generate a password hash
func HashPassword(password string) (string, error) {
 
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password")
	}
	return string(hashedPassword), nil
}

// CheckPassword checks if the provided password is correct or not
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}