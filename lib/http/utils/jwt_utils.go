package modularHTTP

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(key []byte) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload) for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate the token string using HMAC
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string, key []byte) (bool, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		// Return the key used to sign the token
		return key, nil
	})
	if err != nil {
		return false, err
	}

	// Verify the token is valid
	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return false, errors.New("invalid token")
	}

	return true, nil
}
