package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type Password string

func Encrypt(password string) (Password, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return Password(hash), err
}

func Verify(password string, hash Password) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
