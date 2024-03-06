package userusecase

import (
	"github.com/sandronister/jwt-auth/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

func (u *User) Login(username, password string) (entity.User, error) {
	user, err := u.repository.GetUserByUsername(username)
	if err != nil {
		return entity.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
