package modularHTTP

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateSecurePassword(password string) (string, error) {
	// Hash the password using bcrypt with a cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func IsPasswordIdentical(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err != nil 

}
