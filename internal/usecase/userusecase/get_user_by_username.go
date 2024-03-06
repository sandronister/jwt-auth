package userusecase

import "github.com/sandronister/jwt-auth/internal/entity"

func (u *User) GetUserByUsername(username string) (entity.User, error) {
	return u.repository.GetUserByUsername(username)
}
