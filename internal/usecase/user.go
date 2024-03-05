package usecase

import (
	"github.com/sandronister/jwt-auth/internal/domain"
	"github.com/sandronister/jwt-auth/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	repository domain.UserRepository
}

func NewUserUsecase(repository domain.UserRepository) *User {
	return &User{repository: repository}
}

func (u *User) CreateUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return u.repository.CreateUser(username, string(hashedPassword))
}

func (u *User) GetUserByUsername(username string) (entity.User, error) {
	return u.repository.GetUserByUsername(username)
}
