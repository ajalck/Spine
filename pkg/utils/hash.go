package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHashFromPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
}

func VerifyPassword(hashedPassword, password string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil, err
}
