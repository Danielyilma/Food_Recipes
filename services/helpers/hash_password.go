package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string) string {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		// Handle the error
		panic(err)
	}

	return string(hashedPass)
}

func CheckPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	var check = true

	if err != nil {
		check = false
	}
	return check
}
