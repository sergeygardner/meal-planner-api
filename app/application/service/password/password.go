package password

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const UserPasswordCost = 14

var (
	passwordSalt      string
	errorCostOfHashed = errors.New("password hash is not properly hashed")
)

func CastPassword(password string) ([]byte, error) {
	bytes, errorGenerateFromPassword := bcrypt.GenerateFromPassword([]byte(passwordSalt+password), UserPasswordCost)

	if errorGenerateFromPassword != nil {
		return nil, errors.Wrapf(errorGenerateFromPassword, "an error occurred while generating a hashed password for a user by provided data %p", &password)
	}

	return bytes, errorGenerateFromPassword
}

func CheckPassword(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordSalt+password)) != nil
}

func CheckHashedPassword(password *string) (bool, error) {
	if cost, errorCost := bcrypt.Cost([]byte(*password)); errorCost != nil {
		return false, errors.Wrapf(errorCost, "an error occurred while getting a cost for a hashed password by provided data %p", password)
	} else if cost != UserPasswordCost {
		return false, errorCostOfHashed
	}

	return true, nil
}

func SetPasswordSalt(salt string) {
	passwordSalt = salt
}
