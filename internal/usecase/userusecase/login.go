package userusecase

import (
	"errors"

	"github.com/sandronister/jwt-auth/internal/tools/tokenservice"
	"golang.org/x/crypto/bcrypt"
)

func (u *User) Login(username, password string) (string, error) {
	user, err := u.repository.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("not authorized")
	}

	token, err := tokenservice.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
