package helpers

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func MakeHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", errors.New("Error make Hash : " + err.Error())
	}

	return string(hash), nil
}

func CheckHashPassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
