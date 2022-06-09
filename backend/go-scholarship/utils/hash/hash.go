package hash

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) (string, error) {
	if p == "" {
		return "", fmt.Errorf("password is empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hashStr := string(hash)

	return hashStr, nil
}

func CheckPassword(hash, p string) error {
	if p == "" {
		return fmt.Errorf("password is empty")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	if err != nil {
		return err
	}

	return nil
}
