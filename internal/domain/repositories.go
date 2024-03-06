package domain

import "github.com/sandronister/jwt-auth/internal/entity"

type UserRepository interface {
	CreateUser(username, password string) error
	GetUserByUsername(username string) (entity.User, error)
	GetUserById(id int) (entity.User, error)
}
