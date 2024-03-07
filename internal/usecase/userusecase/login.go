package userusecase

import (
	"errors"

	"github.com/sandronister/jwt-auth/internal/tools/middleware"
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

	token, err := middleware.GenerateToken(user.ID, struct {
		Name string
		Role string
	}{
		Name: user.Username,
		Role: "Admin",
	})

	if err != nil {
		return "", err
	}

	return token, nil
}
