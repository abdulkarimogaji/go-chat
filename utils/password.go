package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates and returns a hashed form of inp
func HashPassword(inp string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inp), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("Failed to hash password")
	}
	return string(hashedPassword), nil
}

// CheckPassword returns error is check doesn't match with hashed
func CheckPassword(check, compare string) error {
	return bcrypt.CompareHashAndPassword([]byte(compare), []byte(check))
}
