package util

import (
	"github.com/resonatecoop/phpass"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

func PHPPassHashPassword(password string) (string, error) {
	var err error

	ph := phpass.New(nil)

	hash, err := ph.Hash([]byte(password))
	if err != nil {
		return "", err
	}

	return string(hash), err
}

func PHPPassCheckPassword(password string, hashPassword string) bool {
	pg := phpass.New(nil)

	return pg.Check([]byte(password), []byte(hashPassword))
}
