package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	salt := 8

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(hashed)

}

func ComparePassword(password string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))

	return err == nil
}
