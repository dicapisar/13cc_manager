package auth

import (
	"github.com/dicapisar/13cc_manager/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string, config *config.AuthConfig) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), int(config.Cost))
	return string(bytes), err
}

func CheckPasswordHash(password *string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(*password))
	return err == nil
}
