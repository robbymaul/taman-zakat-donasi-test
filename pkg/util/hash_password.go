package util

import (
	"strings"

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

func ComparePassword(hashPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func ComparePasswordWP(hashPassword string, password string) bool {
	wp := strings.Split(hashPassword, "$wp")
	hashPW := wp[1]
	err := bcrypt.CompareHashAndPassword([]byte(hashPW), []byte(password))
	if err != nil {
		return false
	}

	return true
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

func PHPPassCheckPassword(hashPassword, password string) bool {
	pg := phpass.New(nil)

	return pg.Check([]byte(password), []byte(hashPassword))
}

func CheckPassword(hashPassword, password string) bool {
	var valid bool

	valid = PHPPassCheckPassword(hashPassword, password)
	if valid {
		return true
	}

	valid = ComparePassword(hashPassword, password)
	if valid {
		return true
	}

	valid = ComparePasswordWP(hashPassword, password)
	if valid {
		return true
	}

	return false
}
