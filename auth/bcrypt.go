package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash Password
func Hash(pass string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Compare a hashed password with plain text
func Compare(hashedPass, plainPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plainPass))
	if err != nil {
		return false
	}

	return true
}
