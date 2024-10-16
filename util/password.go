package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashedPassword return bcrypt of the hash password
func HashedPassword(password string) (string, error) {
	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(bcryptedPassword), nil
}

// CheckPassword checks the given password is correct or not
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}