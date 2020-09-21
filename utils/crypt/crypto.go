package crypt

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func Validate(hp string, pt string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hp), []byte(pt)); err != nil {
		return err
	}
	return nil
}
