package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HassPass(p string) string {
	salt := 8
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		return "salah"
	}
	return string(hash)
}

func ComparePass(h, p []byte) bool {
	err := bcrypt.CompareHashAndPassword(h, p)

	return err == nil
}
