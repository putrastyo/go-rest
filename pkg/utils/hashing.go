package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Hash(text string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash: " + err.Error())
	}
	return string(hashed), nil
}

func CheckHash(hashed, text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(text))
	return err == nil
}
