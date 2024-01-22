package common

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (hashed string, err error) {
	hashedPasssword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPasssword), err
}

func ComparePassword(password, verifypassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(verifypassword))
}
