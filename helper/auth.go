package helper

import "golang.org/x/crypto/bcrypt"

func GenerateHashedPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err.Error()
	}

	return string(hashedPassword)
}
