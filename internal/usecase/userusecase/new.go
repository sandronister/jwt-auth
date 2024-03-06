package userusecase

import (
	"github.com/sandronister/jwt-auth/internal/domain"
)

type User struct {
	repository domain.UserRepository
}

func NewUserUsecase(repository domain.UserRepository) *User {
	return &User{repository: repository}
}
