package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(p string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to generate password hash: %v", err)
		return ""
	}
	return string(hash)
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
